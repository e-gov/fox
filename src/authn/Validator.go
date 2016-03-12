package authn

import (
	"encoding/json"
	fernet "github.com/fernet/fernet-go"
	"sync"
	"time"
	"io/ioutil"
	"util"
//	"fmt"
)

var (
	keys []*fernet.Key
	validateLock *sync.RWMutex
)

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
	m := fernet.VerifyAndDecrypt(tok, time.Duration(util.GetConfig().Authn.TokenTTL)*time.Minute, GetValidateKeys())

	err := json.Unmarshal(m, &message)
	if err != nil {
		log.Debug("Failed to unmarshall the message " + string(m))
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
	
	for _, name := range filenames{
		log.Debugf("Attempting to load validation key from %s", name)
		b, err := ioutil.ReadFile(name)

		if err != nil {
			log.Errorf("Could open a key file %s", name)
		}else{
			k, err := fernet.DecodeKey(string(b))

			if err != nil {
				log.Errorf("Could not parse a key from %s", name)
			}else{
				log.Debugf("Successfully loaded mint key from %s", name)
				tempKeys = append(tempKeys, k)			
			}
		}

	}
	if len(tempKeys) == 0{
		panic("Could not read any validation keys")
	}
	// Store only after we are sure loading was good
	validateLock.Lock()
	defer validateLock.Unlock()
	keys = tempKeys
}

func GetValidateKeys() []*fernet.Key{
	validateLock.RLock()
	defer validateLock.RUnlock()
	return keys
}
