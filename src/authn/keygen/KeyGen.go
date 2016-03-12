package main

import(
	fernet "github.com/fernet/fernet-go"
	"fmt"
)
	
// Generate a fernet key
func main() {
	var k *fernet.Key
	
	k = new(fernet.Key)
	_ = k.Generate()
	fmt.Print(k.Encode())
	
}