package fox

import (
	"sync"
	"gopkg.in/gcfg.v1"
	"log"
)

// Data structure for configuration
type Config struct {
	Storage struct{
		Filepath string
			 }
}

// Sanitize the configuration 
func sanitize(c *Config){
	// Make sure the db path ends with a forwardslash
	s := c.Storage.Filepath
	if len(s) > 0{
		if string(s[len(s) - 1]) != "/" {
			c.Storage.Filepath = s + "/"
		}
	}
}

func LoadConfig(){
	var isFatal bool = (config == nil)
	var fName string = "config.gcfg"
	var tmp *Config
	
	tmp = new(Config)
	
	if err := gcfg.ReadFileInto(tmp, fName); err != nil{
		// No config to start up on
		if isFatal {
			panic(err)	
		} else {
			log.Println("Failed to load configuration from " + fName)
			return
		}	
	}
	
	sanitize(tmp)
	cLock.Lock()
	config = tmp
	cLock.Unlock()
	log.Println("Success loading configuration from " + fName)
}

func getConfig() *Config{
	cLock.RLock()
	defer cLock.RUnlock()	
	return config
}


// Global to hold the conf and a lock
var (
	config *Config
	cLock = new(sync.RWMutex)
)