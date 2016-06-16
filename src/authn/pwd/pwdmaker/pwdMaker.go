package main

import (
	pwd "authn/pwd"
	"flag"
	"fmt"
	"os"
	"util"

	logging "github.com/op/go-logging"
)

// Generate a token based on the key file and username
func main() {
	var password string
	var username string
	const required = "REQUIRED"

	be := logging.NewLogBackend(os.Stderr, "", 0)

	logging.SetBackend(be)
	logging.SetLevel(logging.CRITICAL, "")

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
