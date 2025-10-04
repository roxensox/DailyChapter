package auth

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"time"
)

func MakeJWT(userID uuid.UUID, tokenSecret string, expiresIn time.Duration) (string, error) {
	// Creates a new JWT and returns it as a signed string

	// Initializes jwt.NumericDate objects
	issuedAt := jwt.NumericDate{
		Time: time.Now().UTC(),
	}
	expiresAt := jwt.NumericDate{
		Time: time.Now().UTC().Add(expiresIn),
	}

	// Creates the JWT
	newJWT := jwt.NewWithClaims(
		// Specifies signing method uses HS256
		jwt.SigningMethodHS256,
		jwt.RegisteredClaims{
			Issuer:    "DailyChapter",
			IssuedAt:  &issuedAt,
			ExpiresAt: &expiresAt,
			Subject:   userID.String(),
		},
	)

	// Generates a signed string with the input tokenSecret
	JWT, err := newJWT.SignedString([]byte(tokenSecret))
	if err != nil {
		return "", err
	}

	// Returns the signed string
	return JWT, nil
}
