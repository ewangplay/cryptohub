# cryptolib

A common cryptographical library in Golang.

## Feature

* SHA2-256 / SHA2-384 / SHA2-512 hashing algorithm.
* Generating Ed25519 / ECDSA key pair.
* Signing & verifying signature using Ed25519 / ECDSA algorithm.

## Usage

Get package:
```
go get github.com/ewangplay/cryptolib
```

### Hashing sample

```
package main

import (
	"fmt"
	"os"
	"encoding/hex"

	cl "github.com/ewangplay/cryptolib"
)

func main() {
	cfg := &cl.Config{
		ProviderName: "SW",
	}
	csp, err := cl.GetCSP(cfg)
	if err != nil {
		fmt.Printf("GetCSP failed: %v\n", err)
		os.Exit(1)
	}

	msg := []byte("hello,world")
	digest, err := csp.Hash(msg, &SHA256Opts{})
	if err != nil {
		fmt.Printf("Hash failed %v\n", err)
		os.Exit(1)
	}
	fmt.Println(hex.EncodeToString(digest))
```

### Generating Ed25519 key pair, signing and verifying signature sample

```
package main

import (
	"fmt"
	"os"

	cl "github.com/ewangplay/cryptolib"
)

func main() {
	cfg := &cl.Config{
		ProviderName: "SW",
	}
	csp, err := cl.GetCSP(cfg)
	if err != nil {
		fmt.Printf("GetCSP failed: %v\n", err)
		os.Exit(1)
	}

	k, err := csp.KeyGen(&cl.ED25519KeyGenOpts{})
	if err != nil {
		fmt.Printf("KeyGen failed: %v\n", err)
		os.Exit(1)
	}

	digest := []byte("hello,world")
	signature, err := csp.Sign(k, digest)
	if err != nil {
		fmt.Printf("Sign failed: %v\n", err)
		os.Exit(1)
	}

	pubKey, err := k.PublicKey()
	if err != nil {
		fmt.Printf("Get public key failed: %v\n", err)
		os.Exit(1)
	}

	valid, err := csp.Verify(pubKey, digest, signature)
	if err != nil {
		fmt.Printf("Verify failed: %v\n", err)
		os.Exit(1)
	}
	if !valid {
		fmt.Println("The signature should be validated.")
		os.Exit(1)
	}

	fmt.Println("The signature is validated successfully.")
}
```
