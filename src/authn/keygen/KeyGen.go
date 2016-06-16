package main

import (
	"fmt"

	fernet "github.com/fernet/fernet-go"
)

// Generate a fernet key
func main() {
	var k *fernet.Key

	k = new(fernet.Key)
	_ = k.Generate()
	fmt.Print(k.Encode())

}
