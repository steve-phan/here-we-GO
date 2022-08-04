package helpers

import (
	"math/rand"
	"time"
)

func RandomNumber(number int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(number)
}
