package auth

import (
	"github.com/alexedwards/argon2id"
)

func HashPassword(ptPass string) (string, error) {
	// Takes in a plaintext password, hashes it, and returns it
	return argon2id.CreateHash(ptPass, argon2id.DefaultParams)
}

func CheckPasswordHash(password, hash string) (bool, error) {
	// Takes in a plaintext password and a hash and checks whether they match
	return argon2id.ComparePasswordAndHash(password, hash)
}
