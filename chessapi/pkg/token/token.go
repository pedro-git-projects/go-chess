package token

import (
	"crypto/rand"
	"fmt"
)

const length = 32 // len in bytes

// Generate is a cryptographically secure token generator
func Generate() (token string, err error) {
	tokenBytes := make([]byte, length)
	_, err = rand.Read(tokenBytes)
	if err != nil {
		return token, err
	}

	token = fmt.Sprintf("%x", tokenBytes)
	return token, nil
}
