package util

import (
	"os"
	"os/user"
	"path/filepath"
	"sync"
	"github.com/op/go-logging"
	"github.com/spf13/viper"
	"strings"
)

var log = logging.MustGetLogger("LoginService")

// Config is the data structure for passing configuration info
type Config struct {
	Version int
	Storage struct {
			Filepath string
		}
	Authn   struct {
			MintKeyName      string
			ValidateKeyNames []string
			TokenTTL         int
		}
}

// Sanitize the configuration
func sanitize(c *Config) {
	s := c.Storage.Filepath
	if len(s) > 0 {
		// Make sure the db path ends with a forwardslash
		if string(s[len(s) - 1]) != "/" {
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
	LoadConfigByName("config")
}

// LoadConfigByName loads a config from a specific file
// Used for separating test from operational configuration
func LoadConfigByName(name string) {
	var isFatal bool
	var tmp *Config

	tmp = new(Config)

	cLock.RLock()
	isFatal = (config == nil)
	cLock.RUnlock()

	userName := getUserName()

	viper.SetConfigName(name)

	configFolder := getConfigPath(userName)
	viper.AddConfigPath(configFolder)
	viper.AddConfigPath(".") // default path


	if err := viper.ReadInConfig(); err != nil {
		// No config to start up on
		if isFatal {
			log.Debugf("Looking for config in: %s", configFolder)
			panic(err)
		} else {
			log.Errorf("Failed to load configuration from %s\n", name)
			return
		}
	}

	log.Infof("Config file found: %s\n", viper.ConfigFileUsed())

	viper.Unmarshal(tmp)
	sanitize(tmp)

	// TODO viper can reload config too. Remove this?
	cLock.Lock()
	if config == nil {
		tmp.Version = 1
	} else {
		tmp.Version = config.Version + 1
	}

	config = tmp
	cLock.Unlock()

	log.Infof("Success loading configuration ver %d from %s", config.Version, viper.ConfigFileUsed())
}

func GetConfig() *Config {
	cLock.RLock()
	defer cLock.RUnlock()
	return config
}

// Return currently logged in user's username
func getUserName() string {
	u, err := user.Current()
	if err != nil {
		log.Errorf("Cannot find current user")
	}
	return u.Username
}

// Generate path to config folder
func getConfigPath(userName string) string {
	wd, _ := os.Getwd()

	pathEl := strings.Split(wd, string(filepath.Separator))

	cfgPath := ""
	for i := 0; i < len(pathEl); i++ {
		cfgPath += pathEl[i] + string(filepath.Separator)

		if pathEl[i] == "src" || pathEl[i] == "bin" {
			cfgPath += "config/" + userName + "/"
			break
		}
	}
	return cfgPath
}

// Global to hold the conf and a lock
var (
	config *Config
	cLock = new(sync.RWMutex)
)
