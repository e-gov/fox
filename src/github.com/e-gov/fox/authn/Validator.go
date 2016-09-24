package authn

import (
	"io/ioutil"
	"sync"

	"github.com/e-gov/fox/util"

	log "github.com/Sirupsen/logrus"
	"github.com/dgrijalva/jwt-go"
	"fmt"
	"crypto/rsa"
)

var (
	key *rsa.PublicKey
	//key *[]byte
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
func Decrypt(tokenString string) *jwt.StandardClaims{

	// If the configuration has changed, re-load the keys
	if confVersion != util.GetConfig().Version {
		loadValidateKeys()
	}

	token, err := jwt.ParseWithClaims(tokenString,  &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return GetValidateKey(), nil
	})


	if err != nil{
		if ve, ok := err.(*jwt.ValidationError); ok{
			if ve.Errors&jwt.ValidationErrorMalformed != 0{
				log.WithFields(log.Fields{
					"token": token,
				}).Error("Malformed token received")

			}else{
				if ve.Errors&jwt.ValidationErrorExpired != 0{
					log.WithFields(log.Fields{
						"token": token,
					}).Error("Token expired")

				}else{
					log.WithFields(log.Fields{
						"token": token,
					}).Error("Could not handle the token ", err)
				}
			}
		}
		return nil
	}else{
		if token.Valid {
			return token.Claims.(*jwt.StandardClaims)
		}else{
			log.WithFields(log.Fields{
				"token": token,
			}).Error("Invalid token but no error given")
			return nil
		}
	}
}

func loadValidateKeys() {
	loadValidateKeyByName(util.GetConfig().Authn.ValidateKeyName)
	confVersion = util.GetConfig().Version
}

// loadValidateKeyByName loads a key by filename and strores it in the struct
// The function is threadsafe and panics if the key file is invalid
func loadValidateKeyByName(filename string) {
	keyPath := util.GetPaths([]string{filename})[0]

	b, err := ioutil.ReadFile(keyPath)

	if err != nil {
		log.WithFields(log.Fields{
			"path": keyPath,
		}).Panic("Failed to load validation key: ", err)
	}

	k, err := jwt.ParseRSAPublicKeyFromPEM(b)
	if err != nil{
		log.WithFields(log.Fields{
			"path": keyPath,
		}).Panic("Failed to parse validation key: ", err)
	}

	log.WithFields(log.Fields{
		"path": keyPath,
	}).Debugf("Successfully loaded validation key from %s", keyPath)

	validateLock.Lock()
	defer validateLock.Unlock()
	key = k
}

// GetValidateKeys returns the key reference in a thread-safe fashion
func GetValidateKey() *rsa.PublicKey {
	validateLock.RLock()
	defer validateLock.RUnlock()
	return key
}
