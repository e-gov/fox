package util

import (
	"gopkg.in/gcfg.v1"
	"os"
	"sync"
	"github.com/op/go-logging"
)

var log = logging.MustGetLogger("LoginService")

// Config is the data structure for passing configuration info
type Config struct {
	Version int
	Storage struct {
		Filepath string
	}
	Authn struct {
		MintKeyName string
		ValidateKeyNames []string
		TokenTTL int
	}
}

// Sanitize the configuration
func sanitize(c *Config) {
	s := c.Storage.Filepath
	if len(s) > 0 {
		// Make sure the db path ends with a forwardslash
		if string(s[len(s)-1]) != "/" {
			s = s + "/"
			log.Debugf("Added forwardslash to db path '%s' ", s)
		}
		// Handle relative paths
		if string(s[0]) != "/" {
			pwd, _ := os.Getwd()
			s = pwd + "/" + s
			log.Debugf("Added pwd to db path '%s' ", s)
		}
		c.Storage.Filepath = s
	}
}

// LoadConfig loads configuration using a hard-coded name
// This is what gets called during normal operation
func LoadConfig() {
	LoadConfigByName("config.gcfg")
}

// LoadConfigByName loads a config from a specific file
// Used for separating test from operational conifguration
func LoadConfigByName(name string) {
	var isFatal = (config == nil)
	var fName = name
	var tmp *Config

	tmp = new(Config)

	if err := gcfg.ReadFileInto(tmp, fName); err != nil {
		// No config to start up on
		if isFatal {
			panic(err)
		} else {
			log.Errorf("Failed to load configuration from %s", fName)
			return
		}
	}

	sanitize(tmp)
	cLock.Lock()
	if config == nil{
		tmp.Version = 1
	}else{
		tmp.Version = config.Version + 1
	}
	
	config = tmp
	cLock.Unlock()
	log.Infof("Success loading configuration ver %d from %s", config.Version, fName)
}

func GetConfig() *Config {
	cLock.RLock()
	defer cLock.RUnlock()
	return config
}


// Global to hold the conf and a lock
var (
	config *Config
	cLock  = new(sync.RWMutex)
)
