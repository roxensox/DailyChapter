package auth

import (
	"crypto/rsa"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"time"
)

func MakeJWT(userID uuid.UUID, privateKey *rsa.PrivateKey, expiresIn time.Duration) (string, error) {
	// Creates a new JWT and returns it as a signed string

	// Initializes jwt.NumericDate objects
	issuedAt := jwt.NumericDate{
		Time: time.Now().UTC(),
	}
	expiresAt := jwt.NumericDate{
		Time: time.Now().UTC().Add(expiresIn),
	}

	claims := jwt.RegisteredClaims{
		Issuer:    "DailyChapter",
		IssuedAt:  &issuedAt,
		ExpiresAt: &expiresAt,
		Subject:   userID.String(),
	}

	// Creates the JWT
	newJWT := jwt.NewWithClaims(
		// Specifies signing method uses RS256
		jwt.SigningMethodRS256,
		claims,
	)

	// Generates a signed string with the input tokenSecret
	JWT, err := newJWT.SignedString(privateKey)
	if err != nil {
		return "", err
	}

	// Returns the signed string
	return JWT, nil
}

func ValidateJWT(tokenString string, privateKey *rsa.PrivateKey) (uuid.UUID, error) {
	keyFunc := func(token *jwt.Token) (interface{}, error) {
		return privateKey, nil
	}
	out, err := jwt.ParseWithClaims(tokenString, &jwt.RegisteredClaims{}, keyFunc)

	if err != nil {
		return uuid.UUID{}, err
	}
	if out.Valid {
		if clms, ok := out.Claims.(*jwt.RegisteredClaims); ok {
			uid, err := uuid.Parse(clms.Subject)
			if err != nil {
				return uuid.UUID{}, err
			}
			return uid, nil
		}
	}
	return uuid.UUID{}, nil
}
