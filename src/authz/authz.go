package authz
import(
	"github.com/op/go-logging"
)
	
var log = logging.MustGetLogger("AuthZ")


// Provider is the generic interface all authorization providers 
// must implement
type Provider interface{
	IsAuthorized(string, string, string) bool	
}

// GetProvider returns the current authz provider
// Ideally is configurable and all but currently just a 
// trivial one is implemented 
func GetProvider() Provider{
	return new(SimpleProvider)
}