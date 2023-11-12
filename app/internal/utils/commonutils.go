package utils

import (
	"math/rand"
	"time"
)

func GenerateSequence(count int) string {
	if count == 0 {
		return ""
	}

	//символы, из которых генерируется случайное значение
	chars := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789_"

	var result string

	source := rand.NewSource(time.Now().UnixNano())
	rng := rand.New(source)

	for i := 0; i < count; i++ {
		result += string(chars[rng.Intn(len(chars))])
	}

	return result
}
