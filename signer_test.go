package surl

import (
	"crypto/rand"
	"errors"
	"net/url"
	"path"
	"sync"
	"testing"
	"time"

	leg100 "github.com/leg100/surl/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var (
	compatModes = []struct {
		name string
		mode Option
	}{
		{
			name: "query",
		},
		{
			name: "compat",
			mode: WithCompatMode(VerifyCompatible),
		},
		{
			name: "compat++",
			mode: WithCompatMode(SignVerifyCompatible),
		},
	}

	bases = []struct {
		name string
		base Option
	}{
		{
			name: "decimal",
		},
		{
			name: "base32",
			base: WithBase32Expiry(),
		},
	}

	opts = []struct {
		name      string
		options   []Option
		compat    []leg100.Option
		skipBench bool
	}{
		{
			name: "no opts",
		},
		{
			name:    "prefix",
			options: []Option{WithPrefixPath("/signed")},
			compat:  []leg100.Option{leg100.PrefixPath("/signed")},
		},
		{
			name:    "skip query",
			options: []Option{WithSkipQuery()},
			compat:  []leg100.Option{leg100.SkipQuery()},
		},
		{
			name:    "skip scheme",
			options: []Option{WithSkipScheme()},
			compat:  []leg100.Option{leg100.SkipScheme()},
		},
		{
			name:    "prefix and skip query",
			options: []Option{WithSkipQuery(), WithPrefixPath("/signed")},
			compat:  []leg100.Option{leg100.SkipQuery(), leg100.PrefixPath("/signed")},
		},
		{
			name:      "prefix and skip scheme",
			options:   []Option{WithSkipScheme(), WithPrefixPath("/signed")},
			compat:    []leg100.Option{leg100.SkipScheme(), leg100.PrefixPath("/signed")},
			skipBench: true,
		},
		{
			name:      "prefix and skip query and skip scheme",
			options:   []Option{WithSkipQuery(), WithSkipScheme(), WithPrefixPath("/signed")},
			compat:    []leg100.Option{leg100.SkipQuery(), leg100.SkipScheme(), leg100.PrefixPath("/signed")},
			skipBench: true,
		},
		{
			name:      "custom query param",
			options:   []Option{WithSignatureParam("my_sig"), WithExpiryParam("my_exp")},
			skipBench: true,
		},
	}
)

func TestSigner(t *testing.T) {
	inputs := []struct {
		name     string
		unsigned string
	}{
		{
			name:     "with query",
			unsigned: "https://example.com/a/b/c?foo=bar",
		},
		{
			name:     "with query",
			unsigned: "https://example.com/a/b/c?foo=bar",
		},
		{
			name:     "without query",
			unsigned: "https://example.com/a/b/c",
		},
		{
			name:     "with only question mark",
			unsigned: "https://example.com/a/b/c?",
		},
		{
			name:     "absolute path",
			unsigned: "/a/b/c",
		},
	}
	// invoke test for each combination of unsigned url, mode, base, and set of
	// options
	for _, tt := range inputs {
		for _, f := range compatModes {
			for _, enc := range bases {
				for _, opt := range opts {
					options := append([]Option{f.mode, enc.base}, opt.options...)
					signer := New([]byte("abc123"), options...)

					t.Run(path.Join(tt.name, f.name, enc.name, opt.name), func(t *testing.T) {
						signed, err := signer.Sign(tt.unsigned, time.Now().Add(time.Second*10))
						require.NoError(t, err)

						// check valid URL
						_, err = url.Parse(signed)
						require.NoError(t, err)

						// check valid signature
						err = signer.Verify(signed)
						require.NoError(t, err)
					})
				}
			}
		}
	}
}

func TestSigner_SkipQuery(t *testing.T) {
	signer := New([]byte("abc123"))

	// Demonstrate the WithSkipQuery option by changing the
	// query string on the signed URL and showing it still verifies.
	t.Run("skip query", func(t *testing.T) {
		sign := New([]byte("abc123"), WithSkipQuery())

		u := "https://example.com/a/b/c?foo=bar"
		signed, err := sign.Sign(u, time.Now().Add(time.Minute))
		require.NoError(t, err)

		signed += "&page_num=3&page_size=20"

		err = sign.Verify(signed)
		require.NoError(t, err)

		t.Run("check original query parameters are intact", func(t *testing.T) {
			u, err := url.Parse(signed)
			require.NoError(t, err)
			assert.Equal(t, "bar", u.Query().Get("foo"))
		})
	})

	// Demonstrate how changing the query string invalidates the signed URL
	t.Run("do not skip query", func(t *testing.T) {
		u := "https://example.com/a/b/c?foo=bar"
		signed, err := signer.Sign(u, time.Now().Add(time.Minute))
		require.NoError(t, err)

		signed += "&page_num=3&page_size=20"

		err = signer.Verify(signed)
		assert.True(t, errors.Is(err, ErrInvalidSignature))
	})
}

func TestSigner_SkipScheme(t *testing.T) {
	// Demonstrate the WithSkipScheme option by changing the scheme on the signed
	// URL and showing it still verifies.
	t.Run("skip scheme", func(t *testing.T) {
		signer := New([]byte("abc123"), WithSkipScheme())

		unsigned := "https://example.com/a/b/c?foo=bar"
		signed, err := signer.Sign(unsigned, time.Now().Add(time.Minute))
		require.NoError(t, err)

		u, err := url.Parse(signed)
		require.NoError(t, err)
		u.Scheme = "http"

		err = signer.Verify(u.String())
		require.NoError(t, err)
	})

	// Demonstrate how changing the scheme invalidates the signed URL
	t.Run("do not skip scheme", func(t *testing.T) {
		signer := New([]byte("abc123"))

		unsigned := "https://example.com/a/b/c?foo=bar"
		signed, err := signer.Sign(unsigned, time.Now().Add(time.Minute))
		require.NoError(t, err)

		u, err := url.Parse(signed)
		require.NoError(t, err)
		u.Scheme = "http"

		err = signer.Verify(u.String())
		assert.True(t, errors.Is(err, ErrInvalidSignature))
	})
}

func TestSigner_Prefix(t *testing.T) {
	signer := New([]byte("abc123"), WithPrefixPath("/signed"))

	t.Run("invalid prefix", func(t *testing.T) {
		err := signer.Verify("http://abc.com/wrongprefix/foo/bar?expiry=123&signature=fJLFKJ3903")
		assert.True(t, errors.Is(err, ErrInvalidFormat))
	})
}

func TestSigner_Errors(t *testing.T) {
	t.Run("expired", func(t *testing.T) {
		signer := New([]byte("abc123"))

		u := "https://example.com/a/b/c?baz=cow&foo=bar"
		signed, err := signer.Sign(u, time.Now())
		require.NoError(t, err)

		err = signer.Verify(signed)
		assert.Equal(t, ErrExpired, err)
	})

	t.Run("relative path", func(t *testing.T) {
		signer := New([]byte("abc123"))
		_, err := signer.Sign("foo/bar", time.Now().Add(time.Minute))
		assert.Error(t, err)
	})

	t.Run("empty url", func(t *testing.T) {
		signer := New([]byte("abc123"))
		_, err := signer.Sign("", time.Now().Add(10*time.Second))
		assert.Error(t, err)
	})

	t.Run("not a url", func(t *testing.T) {
		signer := New([]byte("abc123"))
		_, err := signer.Sign("cod", time.Now().Add(10*time.Second))
		assert.Error(t, err)
	})

	t.Run("scheme has changed", func(t *testing.T) {
		signer := New([]byte("abc123"))
		signed, err := signer.Sign("https://example.com/a/b/c?baz=cow&foo=bar", time.Now().Add(10*time.Second))
		require.NoError(t, err)

		hacked, err := url.Parse(signed)
		require.NoError(t, err)
		hacked.Scheme = "http"

		err = signer.Verify(hacked.String())
		assert.Error(t, err)
	})

	t.Run("hostname has changed", func(t *testing.T) {
		signer := New([]byte("abc123"))
		signed, err := signer.Sign("https://example.com/a/b/c?baz=cow&foo=bar", time.Now().Add(10*time.Second))
		require.NoError(t, err)

		hacked, err := url.Parse(signed)
		require.NoError(t, err)
		hacked.Host = "hacked.com:1337"

		err = signer.Verify(hacked.String())
		assert.Error(t, err)
	})
}

func TestSignerCompat(t *testing.T) {
	inputs := []struct {
		name     string
		unsigned string
	}{
		{
			name:     "with query",
			unsigned: "https://example.com/a/b/c?foo=bar",
		},
		{
			name:     "with long query",
			unsigned: "https://example.com/a/b/c?foo=bar&foo2=bazz&abc=xyz",
		},
		{
			name:     "without query",
			unsigned: "https://example.com/a/b/c",
		},
		{
			name:     "with only question mark",
			unsigned: "https://example.com/a/b/c?",
		},
		{
			name:     "absolute path",
			unsigned: "/a/b/c",
		},
	}
	for _, tt := range inputs {
		// Sign/verify compatible.
		for _, enc := range bases {
			for _, opt := range opts {
				if len(opt.compat) == 0 {
					continue
				}

				options := append([]Option{WithCompatMode(SignVerifyCompatible), enc.base}, opt.options...)
				signer := New([]byte("abc123"), options...)
				leg100Signer := leg100.New([]byte("abc123"), opt.compat...)

				t.Run(path.Join(tt.name, "compat++ sign+leg100verify", enc.name, opt.name), func(t *testing.T) {
					signed, err := signer.Sign(tt.unsigned, time.Now().Add(time.Second*10))
					require.NoError(t, err)

					// check valid URL
					_, err = url.Parse(signed)
					require.NoError(t, err)

					// check valid signature
					err = leg100Signer.Verify(signed)
					require.NoError(t, err)
				})

				t.Run(path.Join(tt.name, "compat++ leg100sign+verify", enc.name, opt.name), func(t *testing.T) {
					signed, err := leg100Signer.Sign(tt.unsigned, time.Now().Add(time.Second*10))
					require.NoError(t, err)

					// check valid URL
					_, err = url.Parse(signed)
					require.NoError(t, err)

					// check valid signature
					err = signer.Verify(signed)
					require.NoError(t, err)
				})
			}
		}

		for _, enc := range bases {
			for _, opt := range opts {
				options := append([]Option{WithCompatMode(VerifyCompatible), enc.base}, opt.options...)
				signer := New([]byte("abc123"), options...)
				leg100Signer := leg100.New([]byte("abc123"), opt.compat...)

				t.Run(path.Join(tt.name, "compat", enc.name, opt.name), func(t *testing.T) {
					signed, err := leg100Signer.Sign(tt.unsigned, time.Now().Add(time.Second*10))
					require.NoError(t, err)

					// check valid URL
					_, err = url.Parse(signed)
					require.NoError(t, err)

					// check valid signature
					err = signer.Verify(signed)
					require.NoError(t, err)
				})
			}
		}
	}
}

func TestSignerCompat_SkipQuery(t *testing.T) {
	// Demonstrate the WithSkipQuery option by changing the
	// query string on the signed URL and showing it still verifies.
	t.Run("skip query", func(t *testing.T) {
		sign := New([]byte("abc123"), WithSkipQuery(), WithCompatMode(SignVerifyCompatible))
		leg100Sign := leg100.New([]byte("abc123"), leg100.SkipQuery())

		u := "https://example.com/a/b/c?foo=bar"
		signed, err := leg100Sign.Sign(u, time.Now().Add(time.Minute))
		require.NoError(t, err)
		signed += "&page_num=3&page_size=20"
		err = sign.Verify(signed)
		require.NoError(t, err)

		signed, err = sign.Sign(u, time.Now().Add(time.Minute))
		require.NoError(t, err)
		signed += "&page_num=3&page_size=20"
		err = leg100Sign.Verify(signed)
		require.NoError(t, err)

		t.Run("check original query parameters are intact", func(t *testing.T) {
			u, err := url.Parse(signed)
			require.NoError(t, err)
			assert.Equal(t, "bar", u.Query().Get("foo"))
		})
	})
}

func TestSignerCompat_SkipScheme(t *testing.T) {
	// Demonstrate the WithSkipScheme option by changing the scheme on the signed
	// URL and showing it still verifies.
	t.Run("skip scheme", func(t *testing.T) {
		signer := New([]byte("abc123"), WithSkipScheme(), WithCompatMode(SignVerifyCompatible))
		leg100Sign := leg100.New([]byte("abc123"), leg100.SkipScheme())

		unsigned := "https://example.com/a/b/c?foo=bar"

		signed, err := leg100Sign.Sign(unsigned, time.Now().Add(time.Minute))
		require.NoError(t, err)
		u, err := url.Parse(signed)
		require.NoError(t, err)
		u.Scheme = "http"
		err = signer.Verify(u.String())
		require.NoError(t, err)

		signed, err = signer.Sign(unsigned, time.Now().Add(time.Minute))
		require.NoError(t, err)
		u, err = url.Parse(signed)
		require.NoError(t, err)
		u.Scheme = "https"
		err = leg100Sign.Verify(u.String())
		require.NoError(t, err)
	})
}

var (
	bu   string
	berr error
)

func Benchmark(b *testing.B) {
	secret := make([]byte, 64)
	_, err := rand.Read(secret)
	require.NoError(b, err)

	// invoke bench for each combination of mode, base, and set of
	// options
	for _, f := range compatModes {
		for _, enc := range bases {
			for _, opt := range opts {
				if opt.skipBench {
					continue
				}
				options := append([]Option{f.mode, enc.base}, opt.options...)

				b.Run(path.Join("sign", f.name, enc.name, opt.name), func(b *testing.B) {
					signer := New(secret, options...)

					var u string
					for n := 0; n < b.N; n++ {
						// store result to prevent compiler eliminating func call
						u, _ = signer.Sign("https://example.com/a/b/c?x=1&y=2&z=3", time.Now().Add(time.Hour))
					}
					// store result in pkg var to to prevent compiler eliminating benchmark
					bu = u
				})

				b.Run(path.Join("verify", f.name, enc.name, opt.name), func(b *testing.B) {
					signer := New(secret, options...)
					signed, _ := signer.Sign("https://example.com/a/b/c?x=1&y=2&z=3", time.Now().Add(time.Hour))

					for n := 0; n < b.N; n++ {
						// store result to prevent compiler eliminating func call
						err = signer.Verify(signed)
					}
					// store result in pkg var to to prevent compiler eliminating benchmark
					berr = err
				})
			}
		}
	}
}

func BenchmarkConcurrent(b *testing.B) {
	b.Run("sign", func(b *testing.B) {
		signer := New([]byte("abc123"))
		var wg sync.WaitGroup
		for b.Loop() {
			for i := 0; i < 500; i++ {
				wg.Add(1)
				go func() {
					defer wg.Done()
					u, _ := signer.Sign("https://example.com/a/b/c?x=1&y=2&z=3", time.Now().Add(time.Hour))
					bu = u
				}()
			}
			wg.Wait()
		}
	})
}
