package utils

import (
	"math/rand"
)

func GenerateRandomIdNumeric() int {
	return rand.Intn(1000)
}

func GetIdFromPath(path string) string {
	return path[1:]
}
