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
