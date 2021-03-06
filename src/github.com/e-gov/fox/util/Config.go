package util

import (
	"os"
	"os/user"
	"path/filepath"
	"strings"
	"sync"

	log "github.com/Sirupsen/logrus"
	"github.com/spf13/viper"
)

//var log = logging.MustGetLogger("Config")

// Config is the data structure for passing configuration info
type Config struct {
	Version int
	Storage struct {
		Filepath string
	}
	Authn struct {
		MintKeyName      string
		ValidateKeyName  string
		TokenTTL         int64
		PwdProvider      struct {
			PwdFileName string
			Salt        string
		}
	}
	Authz struct {
		Provider     string
		LDAPProvider struct {
			User     string
			Password string
		}
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
	LoadConfigByPathWOExtension("config")
}

// LoadConfigByName loads a config from a specific file
// Used for separating test from operational configuration
func LoadConfigByPathWOExtension(name string) {
	var isFatal bool
	var tmp *Config

	tmp = new(Config)

	cLock.RLock()
	isFatal = (config == nil)
	cLock.RUnlock()
	viper.SetConfigName(name)
	viper.SetConfigType("json")

	userConfigFolder := getUserConfigFolderPath()
	configFolder := getConfigFolderPath()

        if (userConfigFolder != "") {
                viper.AddConfigPath(userConfigFolder) // user's own personal config file
        }
        if (configFolder != "") {
                viper.AddConfigPath(configFolder)     // General fallback config file
        }

	if err := viper.ReadInConfig(); err != nil {
		// No config to start up on
		if isFatal {
			log.Debugf("Failed to find config in: %s (user) or %s (fallback)",
				userConfigFolder, configFolder)
			panic(err)
		} else {
			log.Errorf("Failed to load configuration from %s\n", name)
			return
		}
	}

	viper.Unmarshal(tmp)
	sanitize(tmp)

	cLock.Lock()
	if config == nil {
		tmp.Version = 1
	} else {
		tmp.Version = config.Version + 1
	}

	config = tmp
	cLock.Unlock()

	log.WithFields(log.Fields{
		"path": viper.ConfigFileUsed(),
		"confVersion": config.Version,
	}).Info("Success loading configuration")

}

// TODO: make it so it loads the config if it is not there
func GetConfig() *Config {
	cLock.RLock()
	defer cLock.RUnlock()
	return config
}

// GetPaths returns absolute paths for input filenames.
// If file exists in user's config folder, returns path to it,
// otherwise returns path to file in 'config/' folder.
// If the file starts with a '/' and exists, leaves it alone - we
// are using absolute paths for some reason
func GetPaths(filenames []string) []string {
	cfgFolder := getConfigFolderPath()
	userCfgFolder := getUserConfigFolderPath()

	var paths []string

	for _, name := range filenames {
		if strings.HasPrefix(name, "/") {
			if _, err := os.Stat(name); err == nil {
				// File exists and starts with a '/', must be an absolute path
				paths = append(paths, name)
				continue
			}
		}

		path := cfgFolder + name
		userPath := userCfgFolder + name

		if _, err := os.Stat(userPath); err == nil {
			paths = append(paths, userPath)
		} else {
			paths = append(paths, path)
		}
	}

	return paths
}

// Generates path to user's config folder
func getUserConfigFolderPath() string {

	userName := getUserName()

        if userName == "" {
                return ""
        }

	cfgFolder := getConfigFolderPath()
	sep := string(filepath.Separator)

	path := cfgFolder + userName + sep

	return path
}

// Return currently logged in user's username
func getUserName() string {
	u, err := user.Current()
	if err != nil {
		log.Errorf("Cannot find current user -- %s", err.Error())
		return ""
	}
	return u.Username
}

// Generates path to general config folder (which contains all user's folders)
func getConfigFolderPath() string {
	sep := string(filepath.Separator)
	wd, _ := os.Getwd()

	wdPath := strings.Split(wd, sep)
	indexOfSrc := lastIndexOf(wdPath, "src")
	indexOfBin := lastIndexOf(wdPath, "bin")

	cfgPath := ""
	var pathEl []string
	if indexOfBin > -1 && indexOfBin > indexOfSrc {
		pathEl = wdPath[:indexOfBin] // take up to bin (exclusive)
	} else if indexOfSrc > -1 {
		pathEl = wdPath[:indexOfSrc] // take up to src (exclusive)
	}

	if len(pathEl) > 0 {
		cfgPath = strings.Join(pathEl, sep) + sep
		cfgPath += "config" + sep
	}

	return cfgPath
}

func lastIndexOf(h []string, n string) int {
	for i := len(h) - 1; i > 0; i-- {
		if h[i] == n {
			return i
		}
	}
	return -1
}

// Global to hold the conf and a lock
var (
	config *Config
	cLock  = new(sync.RWMutex)
)
