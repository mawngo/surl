# surl

Create signed URLs in Go.

This is a trimmed down and optimized version of [leg100/surl](https://github.com/leg100/surl).

What's changed:

- ! Remove support for base58 expiry.
- ! Remove support for path formatter, only support query formatter.
- ! All `Option` now have `With` prefix.
- ! Optimized query format ter logic, this results in incompatible link format between this fork and
  original `leg100/surl` implementation.
- Support for changing expiry and signature query parameter name.

This fork can be configured to sign/verify link in a compatible way with original `leg100/surl` library's default config
using `WithCompatMode`. However, it is only guaranteed to compatible with `v2.0.0`
or [this commit](https://github.com/leg100/surl/commit/b65ab8d97a14851f8a9f80eae89a48b59efbe5d9)

## Installation

`go get -u github.com/mawngo/surl`

## Usage

```golang
package main

import (
	"fmt"
	"log"
	"time"

	"github.com/mawngo/surl"
)

func main() {
	signer := surl.New([]byte("secret_key"))

	// Create a signed URL that expires in one hour.
	signed, _ := signer.Sign("https://example.com/a/b/c?foo=bar", time.Now().Add(time.Hour))
	fmt.Println(signed)
	// Outputs something like:
	// https://example.com/a/b/c?foo=bar&expiry=1753548646&signature=QgtNxB9MsXQagB6m6vDe2j2WbuOncCcJcI34ze4AJUQ

	err := signer.Verify(signed)
	if err != nil {
		log.Fatal("verification failed: ", err.Error())
	}
	fmt.Println("verification succeeded")
}
```

## Notes

* Any change in the order of the query parameters in a signed URL renders it invalid, unless `SkipQuery` is specified.
