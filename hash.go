package compare

import (
	"golang.org/x/crypto/bcrypt"
)

// Password Yang Ingin Di has
// Masukkan Password
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}
