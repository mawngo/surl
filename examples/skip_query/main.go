package main

import (
	"fmt"
	"log"
	"net/url"
	"time"

	"github.com/mawngo/surl"
)

func main() {
	signer := surl.New([]byte("secret_key"), surl.WithSkipQuery())

	// Create a signed URL that expires in one hour.
	signed, _ := signer.Sign("https://example.com/a/b/c?page=1", time.Now().Add(time.Hour))

	pageTwo, _ := url.Parse(signed)
	q := pageTwo.Query()
	q.Set("page", "2")
	pageTwo.RawQuery = q.Encode()

	// Verification should still succeed despite page query having changed.
	err := signer.Verify(pageTwo.String())
	if err != nil {
		log.Fatal("verification failed: ", err.Error())
	}
	fmt.Println("verification succeeded")
}
