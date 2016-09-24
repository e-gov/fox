package authn

import (
	"io/ioutil"
	"sync"

	"github.com/e-gov/fox/util"

	jwt "github.com/dgrijalva/jwt-go"
	log "github.com/Sirupsen/logrus"
	"time"
	"crypto/rsa"
)

var mint struct {
	Key *rsa.PrivateKey
	From string
	sync.RWMutex
}

func InitMint() {
	confVersion = util.GetConfig().Version
	loadMintKey()
}

func loadMintKey() {
	LoadMintKeyByName(util.GetConfig().Authn.MintKeyName)
	confVersion = util.GetConfig().Version
}

// loadMintKeyByName loads a key by filename and stores it in the struct
// The function is threadsafe and panics if the key file is invalid
func LoadMintKeyByName(filename string) {

	keyPath := util.GetPaths([]string{filename})[0]


	b, err := ioutil.ReadFile(keyPath)

	if err != nil {
		log.WithFields(log.Fields{
			"path": keyPath,
		}).Panic("Failed to load mint key: ", err)
	}

	k, err := jwt.ParseRSAPrivateKeyFromPEM(b)
	if err != nil {
		log.WithFields(log.Fields{
			"path": keyPath,
		}).Panic("Failed to parse mint key: ", err)
	}


	log.WithFields(log.Fields{
		"path": keyPath,
	}).Debugf("Successfully loaded mint key from %s", keyPath)
	// Store only after we are sure loading was good

	mint.Lock()
	defer mint.Unlock()
	mint.Key = k
	mint.From = keyPath
}

// GetToken wraps the incoming username into a TokenStruct, serializes the result to json
// and generates a Fernet token based on the resulting string
func GetToken(username string) string {
	// If the configuration has changed, re-load the keys
	if confVersion != util.GetConfig().Version {
		loadMintKey()
	}

	claims := jwt.StandardClaims{
			Issuer: "FoxAuthn",
			Subject: username,
			IssuedAt: time.Now().Unix(),
			ExpiresAt: time.Now().Add(time.Duration(util.GetConfig().Authn.TokenTTL) * time.Second).Unix(),
	}

	log.WithFields(log.Fields{
		"claims":claims,
	}).Debug("Going to sign with these claims")

	token := jwt.NewWithClaims(jwt.SigningMethodRS384, claims)
	ss, err := token.SignedString(GetKey())
	if err != nil{
		log.WithFields(log.Fields{
			"path": mint.From,
		}).Panic("Failed to create signed token: ", err)
	}
	return ss
}

// GetKey returns the current key used for session tokens
// If the key not initialized, nil is returned
func GetKey() *rsa.PrivateKey {
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

	log.Debug("Reissue. Username = " + decryptedTokensUsername)
	newToken = GetToken(decryptedTokensUsername)

	return newToken, nil
}
