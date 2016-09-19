package main

import (
	pwd "github.com/e-gov/fox/authn/pwd"
	"flag"
	"fmt"

	"github.com/e-gov/fox/util"
)

// Generate a token based on the key file and username
func main() {
	var password string
	var username string
	const required = "REQUIRED"


	flag.StringVar(&username, "user", required, "Username to create the token for")
	flag.StringVar(&password, "pwd", required, "Password to use")

	flag.Parse()

	if username == required || password == required {
		flag.PrintDefaults()
		return
	}

	util.LoadConfig()
	fmt.Printf("%s=%s\n", username, pwd.HashPassword(password))
}
