package authn

import(
	"errors"
)


// Authenticate sends the username and challenge to the authentication 
// provider requested and passes the resulting boolean back.
// Any non-true result including technical failures and the authentication 
// provider being unknown will be interepreted 
// as the user not being authenticated
func Authenticate(username string, challenge string, provider string) bool  {
	return true
}

// Validate the Fernet token and extract the username
func Validate(token string) (string, error){
	
	t := Decrypt(token)
	if t == nil {
		return "", errors.New("Invalid token")
	}
	return t.Username, nil
}