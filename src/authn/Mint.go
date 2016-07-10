package authn

import (
	"encoding/json"
	"io/ioutil"
	"sync"
	"time"
	"util"

	fernet "github.com/fernet/fernet-go"
)

var mint struct {
	*fernet.Key
	*sync.RWMutex
}

func InitMint() {
	confVersion = util.GetConfig().Version
	loadMintKey()
}

func loadMintKey() {
	LoadMintKeyByName(util.GetConfig().Authn.MintKeyName)
	confVersion = util.GetConfig().Version
}

// loadMintKeyByName loads a key by filename and strores it in the struct
// The function is threadsafe and panics if the key file is invalid
func LoadMintKeyByName(filename string) {

	keyPath := util.GetPaths([]string{filename})[0]

	log.Debugf("Attempting to load mint key from %s", keyPath)
	b, err := ioutil.ReadFile(keyPath)

	if err != nil {
		panic(err)
	}

	k, err := fernet.DecodeKey(string(b))

	if err != nil {
		panic(err)
	}

	log.Debugf("Successfully loaded mint key from %s", keyPath)
	// Store only after we are sure loading was good
	mint.Lock()
	defer mint.Unlock()
	mint.Key = k
}

// GetToken wraps the incoming username into a TokenStruct, serializes the result to json
// and generates a Fernet token based on the resulting string
func GetToken(username string) string {
	var t []byte

	// If the configuration has changed, re-load the keys
	if confVersion != util.GetConfig().Version {
		loadMintKey()
	}

	n, _ := time.Now().MarshalText()
	t, _ = json.Marshal(TokenStruct{Username: username, MintTime: string(n)})

	token, err := fernet.EncryptAndSign(t, GetKey())
	if err != nil {
		panic(err)
	}
	return string(token)
}

// GetKey returns the current key used for session tokens
// If the key not initialized, nil is returned
func GetKey() *fernet.Key {
	mint.RLock()
	defer mint.RUnlock()
	return mint.Key

}

// ReissueToken re-issues a token based on a previous valid token
func ReissueToken(token string) (string, error) {
	var newToken string
	var decryptedTokensUsername string
	var e error

	if decryptedTokensUsername, e = Validate(token); e != nil {
		log.Info("Reissue request for invalid token " + token)
		return "", e
	}

	newToken = GetToken(decryptedTokensUsername)

	return newToken, nil
}
