# surl

Create signed URLs in Go.

This is a trimmed down and optimized version of [leg100/surl](https://github.com/leg100/surl).

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
	// https://example.com/a/b/c?expiry=1667331055&foo=bar&signature=TGvxmRwpoAUt9YEIbeJ164lMYrzA2DBnYB9Lcy9m1T

	err := signer.Verify(signed)
	if err != nil {
		log.Fatal("verification failed: ", err.Error())
	}
	fmt.Println("verification succeeded")
}
```

## Notes

* Any change in the order of the query parameters in a signed URL renders it invalid, unless `SkipQuery` is specified.
