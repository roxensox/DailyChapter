package auth

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
)

func MakeRefreshToken() (string, error) {
	// Generates a random 32-byte refresh token string

	nums := make([]byte, 32)
	_, err := rand.Read(nums)
	if err != nil {
		fmt.Println("Failed to generate token")
		return "", err
	}
	outStr := hex.EncodeToString(nums)
	return outStr, nil
}
