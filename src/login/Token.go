package login

import(
	"io/ioutil"
	fernet "github.com/fernet/fernet-go"
	"encoding/json"
	"time"
)

// TokenStruct contains the payload structure minted into the auth tokens 
type TokenStruct struct {
	Username	string	`json:"username"`
	MintTime 	string  `json:"mintTime"`
}

// Token is the return type for the API
type Token struct {
	Token 		string `json:"token"`
}

// Global variable to keep the key in
var mintKey *fernet.Key
var keyName = "key.base64"

// LoadKey initializes a Fernet key based on the contents of the fixed filename.
// It panics, if the file cannot be found
func LoadKey(){
	LoadKeyByName(keyName)
}

func LoadKeyByName(filename string)  {
	b, err := ioutil.ReadFile(filename)
	
	if err != nil{
		panic(err)		
	}

	k, err := fernet.DecodeKey(string(b))
	
	if err != nil{
		panic(err)		
	}

	// Store only after we are sure loading was good
	mintKey = k	
}

// GetToken wraps the incoming username into a TokenStruct, serializes the result to json
// and generates a Fernet token based on the resulting string
func GetToken(username string) string{
	var m []byte
	
	t, _ := time.Now().MarshalText()
	m, _ = json.Marshal(TokenStruct{Username: username, MintTime: string(t)})
	
	token, err := fernet.EncryptAndSign(m, mintKey)
	if err != nil{
		panic(err)
	}		
	return string(token)
}