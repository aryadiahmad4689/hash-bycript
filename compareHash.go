package compare

import (
	"golang.org/x/crypto/bcrypt"
)

// PASSWORD AND HASH FROM DB
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
