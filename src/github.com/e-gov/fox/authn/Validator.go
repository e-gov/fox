package authn

import (
	"encoding/json"
	"io/ioutil"
	"math"
	"sync"
	"time"

	"github.com/e-gov/fox/util"

	fernet "github.com/fernet/fernet-go"
)

var (
	keys         []*fernet.Key
	validateLock *sync.RWMutex
)

// InitValidator initializes the validator by storing current config version,
// creating a new lock and loading validation keys
func InitValidator() {
	confVersion = util.GetConfig().Version
	validateLock = new(sync.RWMutex)
	loadValidateKeys()
}

// Decrypt decrypts a string containing a token
// It returns nil if
// - the token has been minted more than tokenTTL minutes ago
// - the token message is not a valid TokenStruct
// - the token cannot be decrypted using known keys
func Decrypt(token string) *TokenStruct {
	var message TokenStruct

	// If the configuration has changed, re-load the keys
	if confVersion != util.GetConfig().Version {
		loadValidateKeys()
	}

	tok := []byte(token)
	// Do the math, TTL in minutes could be a fraction
	ttl := int64(math.Floor(util.GetConfig().Authn.TokenTTL * float64(time.Minute)))

	m := fernet.VerifyAndDecrypt(tok, time.Duration(ttl), GetValidateKeys())

	err := json.Unmarshal(m, &message)
	if err != nil {
		log.Info("Failed to unmarshall token contents " + string(m))
		return nil
	}
	return &message
}

func loadValidateKeys() {
	loadValidateKeyByName(util.GetConfig().Authn.ValidateKeyNames)
	confVersion = util.GetConfig().Version
}

// loadValidateKeyByName loads a key by filename and strores it in the struct
// The function is threadsafe and panics if the key file is invalid
func loadValidateKeyByName(filenames []string) {
	var tempKeys []*fernet.Key

	keyPaths := util.GetPaths(filenames)

	for _, path := range keyPaths {

		log.Debugf("Attempting to load validation key from %s", path)
		b, err := ioutil.ReadFile(path)

		if err != nil {
			log.Errorf("Could not open a key file %s", path)
		} else {
			k, err := fernet.DecodeKey(string(b))

			if err != nil {
				log.Errorf("Could not parse a key from %s", path)
			} else {
				log.Debugf("Successfully loaded validation key from %s", path)
				tempKeys = append(tempKeys, k)
			}
		}

	}
	if len(tempKeys) == 0 {
		panic("Could not read any validation keys")
	}
	// Store only after we are sure loading was good
	validateLock.Lock()
	defer validateLock.Unlock()
	keys = tempKeys
}

// GetValidateKeys returns the key reference in a thread-safe fashion
func GetValidateKeys() []*fernet.Key {
	validateLock.RLock()
	defer validateLock.RUnlock()
	return keys
}
