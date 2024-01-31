package helper

import (
	"math/rand"
	"time"
)

var characters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

var seed *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))

func GenerateRandomCode(length int) string {
	bytes := make([]byte, length)

	for i := range bytes {
		bytes[i] = characters[seed.Intn(len(characters))]
	}

	return string(bytes)
}
