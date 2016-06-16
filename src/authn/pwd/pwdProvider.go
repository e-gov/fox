package authn

import (
	"bufio"
	"encoding/base64"
	"os"
	"strings"
	"util"

	"github.com/op/go-logging"
	bcrypt "golang.org/x/crypto/bcrypt"
)

var log = logging.MustGetLogger("pwdProvider")

// PwdProvider creates a placeholder for the Provider interface implementation
type PwdProvider struct{}

// Authenticate implements the Provider interface authenticating a user against a
// password file. None of this is intended for prodcution use.
func (p *PwdProvider) Authenticate(user string, challenge string) bool {
	salt := util.GetConfig().Authn.PwdProvider.Salt
	pwdFilePath := util.GetPaths([]string{util.GetConfig().Authn.PwdProvider.PwdFileName})[0]
	log.Debugf("Reading passwords from %s", pwdFilePath)
	file, err := os.Open(pwdFilePath)

	if err != nil {
		panic(err)
	}

	defer file.Close()

	s := bufio.NewScanner(file)
	for s.Scan() {
		u, p := getUnP(s.Text())
		log.Debugf("Validating %s against %s and %s", user, u, p)
		if u == user {
			pwd, err := base64.StdEncoding.DecodeString(p)
			if err != nil {
				log.Debugf("Base64 decode failed for user %s", user)
				return false
			}
			r := bcrypt.CompareHashAndPassword(pwd, []byte(salt+challenge))
			if r == nil {
				log.Debugf("Authenticated user %s", user)
				return true
			}
		}
	}
	log.Debugf("User %s not found", user)

	return false
}

func getUnP(s string) (string, string) {
	return strings.Split(s, "=")[0], strings.Split(s, "=")[1]
}

// HashPassword creates a valid base64 encoded hash for the given password
func HashPassword(password string) string {
	salt := util.GetConfig().Authn.PwdProvider.Salt

	b, _ := bcrypt.GenerateFromPassword([]byte(salt+password), 10)

	return base64.StdEncoding.EncodeToString(b)

}
