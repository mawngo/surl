package surl

import (
	"crypto/subtle"
	"encoding/base32"
	"encoding/base64"
	"encoding/binary"
	"errors"
	"golang.org/x/crypto/blake2b"
	"hash"
	"log/slog"
	"net/url"
	"path"
	"strconv"
	"strings"
	"sync"
	"time"
)

const (
	defaultExpiryParam    = "expiry"
	defaultSignatureParam = "signature"
	defaultExpiryBase     = 10
)

var (
	// ErrInvalidSignature is returned when the signature is invalid.
	ErrInvalidSignature = errors.New("invalid signature")
	// ErrInvalidFormat is returned when the format of the signed URL is
	// invalid.
	ErrInvalidFormat = errors.New("invalid format")
	// ErrExpired is returned when a signed URL has expired.
	ErrExpired = errors.New("URL has expired")
)

// Signer is capable of signing and verifying signed URLs with an expiry.
type Signer struct {
	pool sync.Pool

	prefix     string
	skipQuery  bool
	skipScheme bool

	compatLvl         CompatLevel
	expiryBase        int
	expiryParam       string
	expiryParamRaw    string
	signatureParam    string
	signatureParamRaw string
}

// New constructs a new signer, performing the one-off task of generating a
// secure hash from the key. The key must be between 0 and 64 bytes long;
// anything longer is truncated. Options alter the default format and behavior
// of signed URLs.
func New(key []byte, opts ...Option) *Signer {
	key = key[0:min(64, len(key))]
	s := &Signer{
		expiryBase:     defaultExpiryBase,
		expiryParam:    defaultExpiryParam,
		signatureParam: defaultSignatureParam,
	}

	// Leave caller options til last so that they override defaults.
	for _, opt := range opts {
		if opt == nil {
			continue
		}
		opt(s)
	}

	s.pool = sync.Pool{
		New: func() any {
			h, err := blake2b.New256(key)
			if err != nil {
				slog.Error("Error creating hash for signing")
				panic(err)
			}
			return h
		},
	}

	s.expiryParamRaw = "&" + s.expiryParam + "="
	s.signatureParamRaw = "&" + s.signatureParam + "="
	return s
}

// Sign generates a signed URL with the given lifespan.
func (s *Signer) Sign(unsigned string, expiry time.Time) (string, error) {
	u, err := url.ParseRequestURI(unsigned)
	if err != nil {
		return "", err
	}
	s.SignInto(u, expiry)
	return u.String(), nil
}

// SignInto add signature into the url with the given lifespan.
// Any modification to the url after signed will make it invalid,
// except for:
// - RawQuery when [WithSkipQuery] is enabled (removing signature and expiry still invalidates the url).
// - Scheme when [WithSkipScheme] is enabled.
func (s *Signer) SignInto(u *url.URL, expiry time.Time) {
	if s.compatLvl >= SignVerifyCompatible {
		// Create legacy sign.
		s.signCompat(u, strconv.FormatInt(expiry.Unix(), 10))
		return
	}

	encodedExpiry := s.encodeExpiry(expiry.Unix())
	// Use string builder to modify the raw query is more efficient, as we only append query param, not modifying
	// or reading any.
	// However, this introduces breaking change, as the old behaviour using url.Values.Encode(),
	// which actually re-orders the query param.
	var q strings.Builder
	q.Grow(len(u.RawQuery) + len(s.expiryParamRaw) + len(encodedExpiry))
	q.WriteString(u.RawQuery)

	// Add expiry param.
	if u.RawQuery != "" {
		q.WriteString(s.expiryParamRaw)
	} else {
		// Does not have '&' if it is the only param.
		q.WriteString(s.expiryParam)
		q.WriteRune('=')
	}
	q.WriteString(encodedExpiry)

	var rawQuery string
	if s.skipQuery {
		rawQuery = s.expiryParam + "=" + encodedExpiry
	} else {
		rawQuery = q.String()
	}
	// Add sign param.
	sign := s.computeSign(*u, rawQuery)
	encodedSign := base64.RawURLEncoding.EncodeToString(sign)
	q.Grow(len(s.signatureParamRaw) + len(encodedSign))
	q.WriteString(s.signatureParamRaw)
	q.WriteString(encodedSign)

	if s.prefix != "" {
		u.Path = path.Join(s.prefix, u.Path)
	}

	u.RawQuery = q.String()
}

// signCompat add signature into this url in a compatible way with leg100/surl implementation default config.
// - Use "expiry" and "signature" query param name.
// - Use base 10 for expiry param.
// - Sorted query params.
// - Support skip query and skip scheme option.
func (s *Signer) signCompat(u *url.URL, encodedExpiry string) {
	q := u.Query()
	// Add expiry param.
	q.Set("expiry", encodedExpiry)

	var rawQuery string
	if s.skipQuery {
		rawQuery = "expiry=" + encodedExpiry
	} else {
		rawQuery = q.Encode()
	}
	// Add sign param.
	sign := s.computeSign(*u, rawQuery)

	encodedSign := base64.RawURLEncoding.EncodeToString(sign)
	q.Set("signature", encodedSign)

	if s.prefix != "" {
		u.Path = path.Join(s.prefix, u.Path)
	}

	u.RawQuery = q.Encode()
}

// computeSign compute the signature for given url.
func (s *Signer) computeSign(u url.URL, rawQuery string) []byte {
	if s.skipScheme {
		u.Scheme = ""
	}
	u.RawQuery = rawQuery

	h := s.pool.Get().(hash.Hash)
	h.Write([]byte(u.String()))
	sig := h.Sum(nil)
	h.Reset()
	s.pool.Put(h)
	return sig
}

// Verify verifies a signed URL, validating its signature and ensuring it is unexpired.
func (s *Signer) Verify(signed string) error {
	u, err := url.ParseRequestURI(signed)
	if err != nil {
		return err
	}

	if !strings.HasPrefix(u.Path, s.prefix) {
		return ErrInvalidFormat
	}
	u.Path = u.Path[len(s.prefix):]

	err = s.verify(u)
	if err == nil {
		return nil
	}

	if s.compatLvl < VerifyCompatible {
		return err
	}

	// We can safely return when the error is expired, as when it happened,
	// the signature and expiry param must have been successfully extracted and decoded,
	// which mean this is actually a valid formatted link.
	//
	// Use equal for speed, as the ErrExpired is always returned directly.
	// nolint:errorlint
	if err == ErrExpired {
		return ErrExpired
	}

	// When compatible mode enabled, we re-verify the old way.
	// This doubles the work.
	return s.verifyCompat(u)
}

// verify verifies the url.
func (s *Signer) verify(u *url.URL) error {
	rawQuery, encodedSig, encodedExpiry, err := s.extractSignatureAndExpiry(u)
	if err != nil {
		return err
	}

	sig, err := base64.RawURLEncoding.DecodeString(encodedSig)
	if err != nil {
		return errors.Join(ErrInvalidSignature, err)
	}

	expiry, err := s.decodeExpiry(encodedExpiry)
	if err != nil {
		return errors.Join(ErrInvalidFormat, err)
	}

	if time.Now().After(time.Unix(expiry, 0)) {
		return ErrExpired
	}

	compare := s.computeSign(*u, rawQuery)
	// Verify the computeSign.
	if subtle.ConstantTimeCompare(sig, compare) != 1 {
		return ErrInvalidSignature
	}
	return nil
}

// extractSignatureAndExpiry extract the signature, and expiry from given url.
// return the url without signature, encoded signature and encoded expiry.
func (s *Signer) extractSignatureAndExpiry(u *url.URL) (string, string, string, error) {
	if s.skipQuery {
		// When skip query is enabled, we cannot use the trick bellow anymore,
		// as user can reorder, changing the query parameter of this link.
		q := u.Query()
		encodedSig := q.Get(s.signatureParam)
		if encodedSig == "" {
			return "", "", "", ErrInvalidFormat
		}
		encodedExpiry := q.Get(s.expiryParam)
		if encodedExpiry == "" {
			return "", "", "", ErrInvalidFormat
		}
		return s.expiryParam + "=" + encodedExpiry, encodedSig, encodedExpiry, nil
	}

	rawQuery := u.RawQuery
	// Min length check.
	// Actually, here we should check for len(defaultExpiryParam)+1 instead, because in the case of no query param,
	// then the raw query will start with "defaultExpiryParam=" (without &).
	// However, because those param always require some value, this validation is still correct.
	if len(rawQuery) < len(s.expiryParamRaw)+len(s.signatureParamRaw) {
		return "", "", "", ErrInvalidFormat
	}

	// Abuse the fact that signature always at the end of the url.
	sigIndex := strings.LastIndex(rawQuery, s.signatureParamRaw)
	if sigIndex < 0 {
		return "", "", "", ErrInvalidFormat
	}
	encodedSig := rawQuery[sigIndex+len(s.signatureParamRaw):]

	// Abuse the fact that expiry always come before signature.
	var expIndex, expParamLen int
	if rawQuery[len(s.expiryParam)] == '=' && strings.HasPrefix(rawQuery, s.expiryParam) {
		// The signed url does not have any param, so the first param is s.expiryParam.
		expIndex = 0
		expParamLen = len(s.expiryParam) + 1
	} else {
		expIndex = strings.LastIndex(rawQuery[0:sigIndex], s.expiryParamRaw)
		expParamLen = len(s.expiryParamRaw)
	}

	if expIndex < 0 {
		return "", "", "", ErrInvalidFormat
	}

	rawQuery = rawQuery[0:sigIndex]
	encodedExpiry := rawQuery[expIndex+expParamLen:]
	return rawQuery, encodedSig, encodedExpiry, nil
}

// verifyCompat verify the url in a compatible way with leg100/surl implementation default config.
// - Use "expiry" and "signature" query param name.
// - Use base 10 for expiry param.
// - Sorted query params.
// - Support skip query and skip scheme option.
func (s *Signer) verifyCompat(u *url.URL) error {
	q := u.Query()
	// Extract the signature from query.
	encodedSig := q.Get("signature")
	if encodedSig == "" {
		return ErrInvalidFormat
	}
	sig, err := base64.RawURLEncoding.DecodeString(encodedSig)
	if err != nil {
		return errors.Join(ErrInvalidSignature, err)
	}
	q.Del("signature")

	// Get the expiry from query.
	// No need to remove the expiry param, as it is included when generate signature.
	encodedExpiry := q.Get("expiry")
	if encodedExpiry == "" {
		return ErrInvalidFormat
	}
	expiry, err := strconv.ParseInt(encodedExpiry, 10, 64)
	if err != nil {
		return errors.Join(ErrInvalidFormat, err)
	}

	if time.Now().After(time.Unix(expiry, 0)) {
		return ErrExpired
	}

	var rawQuery string
	if s.skipQuery {
		rawQuery = "expiry=" + encodedExpiry
	} else {
		rawQuery = q.Encode()
	}
	// Verify the computeSign.
	compare := s.computeSign(*u, rawQuery)
	if subtle.ConstantTimeCompare(sig, compare) != 1 {
		return ErrInvalidSignature
	}

	// valid, unexpired, signature
	return nil
}

func (s *Signer) encodeExpiry(expiry int64) string {
	if s.expiryBase == 32 {
		b := make([]byte, 8)
		binary.BigEndian.PutUint64(b, uint64(expiry))
		return base32.StdEncoding.EncodeToString(b)
	}
	return strconv.FormatInt(expiry, s.expiryBase)
}

func (s *Signer) decodeExpiry(expiry string) (int64, error) {
	if s.expiryBase == 32 {
		bytes, err := base32.StdEncoding.DecodeString(expiry)
		if err != nil {
			return 0, err
		}
		return int64(binary.BigEndian.Uint64(bytes)), nil
	}
	return strconv.ParseInt(expiry, s.expiryBase, 64)
}
