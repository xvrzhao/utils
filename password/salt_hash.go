package password

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/micro-stacks/utils/strings"
)

// SaltHashPwd generates a salt and the salt-hashed password.
// The salt-hashed generation algorithm is:
//   saltHashPassword = md5(md5(password) + salt)
func SaltHashPwd(password string, saltLen int) (saltHashPassword, salt string) {
	salt = strings.RandLetterNum(saltLen)
	h := md5.New()
	h.Write([]byte(password))
	saltHashPassword = hex.EncodeToString(h.Sum(nil)) + salt
	h.Reset()
	h.Write([]byte(saltHashPassword))
	saltHashPassword = hex.EncodeToString(h.Sum(nil))
	return
}

// VerifySaltHashPwd verifies whether saltHashPassword was generated from password and salt.
func VerifySaltHashPwd(password, salt, saltHashPassword string) bool {
	h := md5.New()
	h.Write([]byte(password))
	p := hex.EncodeToString(h.Sum(nil)) + salt
	h.Reset()
	h.Write([]byte(p))
	if p = hex.EncodeToString(h.Sum(nil)); p != saltHashPassword {
		return false
	}
	return true
}
