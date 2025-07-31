package surl

// CompatLevel specify the compatibility mode for [Signer].
// By default, we use incompatible mode for maximum performance.
// Url that signed by this [Signer] can't be verified by leg100/surl.
// Url that signed by leg100/surl [Signer] can't be verified by this optimized [Signer].
type CompatLevel int

const (
	// VerifyCompatible can verify both leg100/surl (default configuration) and optimized url.
	VerifyCompatible = 1
	// SignVerifyCompatible produce url that is similar to leg100/surl (default configuration).
	// Auto enable [VerifyCompatible] when set.
	SignVerifyCompatible = 2
)

// Option permits customising the construction of a Signer.
type Option func(*Signer)

// WithSkipQuery instructs Signer to skip the query string when computing the
// signature. This is useful, say, if you have pagination query parameters,
// but you want to use the same signed URL regardless of their value.
func WithSkipQuery() Option {
	return func(s *Signer) {
		s.skipQuery = true
	}
}

// WithSkipScheme instructs Signer to skip the scheme when computing the signature.
// This is useful, say, if you generate signed URLs in production, where you use https,
// but you want to use these URLs in development too, where you use http.
func WithSkipScheme() Option {
	return func(s *Signer) {
		s.skipScheme = true
	}
}

// WithPrefixPath prefixes the signed URL's path with a string. This can make it easier for a server
// to differentiate between signed and non-signed URLs. Note: the prefix is not
// part of the signature computation.
func WithPrefixPath(prefix string) Option {
	return func(s *Signer) {
		s.prefix = prefix
	}
}

// WithHexExpiry instructs [Signer] to use hex to encode the expiry.
func WithHexExpiry() Option {
	return func(s *Signer) {
		s.expiryBase = 16
	}
}

// WithBase32Expiry instructs [Signer] to use base32 to encode the expiry.
func WithBase32Expiry() Option {
	return func(s *Signer) {
		s.expiryBase = 32
	}
}

// WithCompatMode configure compatibility behaviour between [Signer] and leg100/surl.
func WithCompatMode(mode CompatLevel) Option {
	return func(s *Signer) {
		s.compatLvl = mode
	}
}

// WithSignatureParam configure the query parameter name of signature.
// This option will be ignored when [WithCompatMode] set to [SignVerifyCompatible].
func WithSignatureParam(name string) Option {
	return func(signer *Signer) {
		signer.signatureParam = name
	}
}

// WithExpiryParam configure the query parameter name of expiry.
// This option will be ignored when [WithCompatMode] set to [SignVerifyCompatible].
func WithExpiryParam(name string) Option {
	return func(signer *Signer) {
		signer.expiryParam = name
	}
}

// WithDisableHashPool disable the hash pool used for generating signature,
// which is enabled by default.
//
// Using pool greatly improving performance in high concurrency scenario.
// Deprecated.
func WithDisableHashPool() Option {
	return func(signer *Signer) {
		signer.disabledPool = true
	}
}
