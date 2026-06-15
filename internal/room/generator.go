package room

import "math/rand"

const chars = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func GenerateCode() string {

	code := make([]byte, 6)

	for i := range code {
		code[i] = chars[rand.Intn(len(chars))]
	}

	return string(code)
}