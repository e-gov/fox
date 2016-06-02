package authz
import(
	"github.com/op/go-logging"
)
	
var log = logging.MustGetLogger("AuthZ")
var provider Provider

// Provider is the generic interface all authorization providers 
// must implement
type Provider interface{
	// IsAuthorized returns if the user is authorized to access the method  
	// on the URL. In case no information is available, the method MUST
	// fall back to false
	IsAuthorized(string, string, string) bool
	// AddRestriction adds information about which roles are necessary to 
	// access a given URL/method 
	AddRestriction(string, string, string)	
	
	// GetRoles returns a list of authorized roles for a given token
	GetRoles(string)[]string
}

// GetProvider returns the current authz provider
// Ideally is configurable and all but currently just a 
// trivial one is implemented 
func GetProvider() Provider{
	if (provider == nil){
		provider = new(SimpleProvider)
	}
	return provider
}