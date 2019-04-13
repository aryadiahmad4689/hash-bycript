package crypto

import (
	"math/rand"
	"time"
)

func Crypto() int {
	rand.Seed(time.Now().UTC().UnixNano())

	return 4 + rand.Intn(30-4)

}
