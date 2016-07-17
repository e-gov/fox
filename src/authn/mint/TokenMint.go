package main

import (
	"flag"
	"fmt"
	"github.com/e-gov/fox/src/authn"
)

// Generate a token based on the key file and username
func main() {
	var keyfile string
	var username string

	flag.StringVar(&keyfile, "key", "key.base64", "A file containing the mint key")
	flag.StringVar(&username, "user", "", "Username to mint the token for")

	flag.Parse()

	authn.LoadMintKeyByName(keyfile)
	fmt.Print(authn.GetToken(username))

}
