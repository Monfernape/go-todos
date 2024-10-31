package utils

import (
	"math/rand"
)

func GenerateRandomIdNumeric() int {
	return rand.Intn(1000)
}
