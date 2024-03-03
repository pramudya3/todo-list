package password

import (
	"golang.org/x/crypto/bcrypt"
)

func HashAndSalt(pwd string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.MinCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func IsMatch(pwdHashed string, pwd string) bool {
	if err := bcrypt.CompareHashAndPassword([]byte(pwdHashed), []byte(pwd)); err != nil {
		return false
	}
	return true
}
