package compare

import (
	"math/rand"
	"time"

	"golang.org/x/crypto/bcrypt"
)

func Crypto() int {
	rand.Seed(time.Now().Unix())

	return 4 + rand.Intn(30-4)

}

// Password Yang Ingin Di has
// Masukkan Password
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), Crypto())
	return string(bytes), err
}
