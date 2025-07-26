package main

import (
	"fmt"
	"log"
	"net/url"
	"time"

	"github.com/mawngo/surl"
)

func main() {
	signer := surl.New([]byte("secret_key"), surl.WithSkipScheme())

	// Create a signed URL that expires in one hour.
	signed, _ := signer.Sign("https://example.com/a/b/c?page=1", time.Now().Add(time.Hour))

	// Change signed URL's scheme from https to http
	u, _ := url.Parse(signed)
	u.Scheme = "http"

	// Verification should still succeed despite scheme having changed.
	err := signer.Verify(u.String())
	if err != nil {
		log.Fatal("verification failed: ", err.Error())
	}
	fmt.Println("verification succeeded")
}
