package authn

import (
	"encoding/json"
	fernet "github.com/fernet/fernet-go"
	"io/ioutil"
	"sync"
	"time"
)

// TokenStruct contains the payload structure minted into the auth tokens
type TokenStruct struct {
	Username string `json:"username"`
	MintTime string `json:"mintTime"`
}

// Token is the return type for the API
type Token struct {
	Token string `json:"token"`
}

// Global variable to keep the key in
var mintKey *fernet.Key
var keyName = "key.base64"
var tokenTTL = 15 // Time To Live of the tokens in minutes

var lock = new(sync.RWMutex)

// LoadKey initializes a Fernet key based on the contents of the fixed filename.
// It panics, if the file cannot be found
func LoadKey() {
	LoadKeyByName(keyName)
}

// LoadKeyByName loads a key by filename. Used for testing
// In normal operation, LoadKey should be called
func LoadKeyByName(filename string) {
	b, err := ioutil.ReadFile(filename)

	if err != nil {
		panic(err)
	}

	k, err := fernet.DecodeKey(string(b))

	if err != nil {
		panic(err)
	}

	// Store only after we are sure loading was good
	lock.Lock()
	defer lock.Unlock()
	mintKey = k
}

// GetToken wraps the incoming username into a TokenStruct, serializes the result to json
// and generates a Fernet token based on the resulting string
func GetToken(username string) string {
	var m []byte

	t, _ := time.Now().MarshalText()
	m, _ = json.Marshal(TokenStruct{Username: username, MintTime: string(t)})

	token, err := fernet.EncryptAndSign(m, mintKey)
	if err != nil {
		panic(err)
	}
	return string(token)
}

// Decrypt decrypts a string containing a token
// It returns nil if 
// - the token has been minted more than tokenTTL minutes ago
// - the token message is not a valid TokenStruct
// - the token cannot be decrypted using known keys
func Decrypt(token string) *TokenStruct{
	var message TokenStruct
	
	tok := []byte(token)
	k := make([]*fernet.Key, 1)
	m := fernet.VerifyAndDecrypt(tok, time.Duration(tokenTTL)*time.Minute, append(k, GetKey())) 

	err := json.Unmarshal(m, message)
	if err != nil {
		return nil
	}
	return &message
}

// GetKey returns the current key used for session tokens
// If the key not initialized, nil is returned
func GetKey() *fernet.Key {
	lock.RLock()
	defer lock.RUnlock()
	return mintKey

}
