package compare

import (
	"github.com/aryadiahmad4689/hash-bycript/crypto"
	"golang.org/x/crypto/bcrypt"
)

// Password Yang Ingin Di has
// Masukkan Password
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), crypto.Crypto())
	return string(bytes), err
}
