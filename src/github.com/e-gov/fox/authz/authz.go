package authz

import (
	"sync"

	"github.com/e-gov/fox/util"

	log "github.com/Sirupsen/logrus"
)

var provider Provider
var pLock = new(sync.RWMutex)

// Provider is the generic interface all authorization providers
// must implement
type Provider interface {
	// IsAuthorized returns if the user is authorized to access the method
	// on the URL. In case no information is available, the method MUST
	// fall back to false
	IsAuthorized(string, string, string) bool
	// AddRestriction adds information about which roles are necessary to
	// access a given URL/method
	AddRestriction(string, string, string)

	// GetRoles returns a list of authorized roles for a given token
	GetRoles(string) []string

	// Returns the name of the provider so we can tell which one is being used
	GetName() string
}

// GetProvider returns the current authz provider and loads a new one
// if configuration has changed
func GetProvider() Provider {
	pLock.Lock()
	defer pLock.Unlock()
	p := util.GetConfig().Authz.Provider
	if p == "" {
		log.Warning("No authorization provider configured, all access will be denied")
		return nil
	}

	if provider == nil {
		loadProvider(p)
	}

	if provider.GetName() != p {
		log.Debugf("Changing authz provider. Was %s is %s", provider.GetName(), p)
		loadProvider(p)
	}

	return provider
}

func loadProvider(name string) {
	if name == "simple" {
		provider = new(SimpleProvider)
		log.Debug("Loading simple authz provider")
	}
	if name == "ldap" {
		provider = new(LdapProvider)
		log.Debug("Loading LDAP authz provider")
	}
	if name == "" {
		provider = nil
	}
}
