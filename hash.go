package compare

import (
	"golang.org/x/crypto/bcrypt"
)

// Password Yang Ingin Di has
// Masukkan Password
func HashPassword(password string, crypto int) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), crypto)
	return string(bytes), err
}
