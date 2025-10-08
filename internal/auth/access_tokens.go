package auth

import (
	"crypto/rsa"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"net/http"
	"strings"
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

func ValidateJWT(tokenString string, publicKey *rsa.PublicKey) (uuid.UUID, error) {
	// Checks the validity of the JWT token and returns its user ID

	// Establishes the parser callback
	keyFunc := func(token *jwt.Token) (any, error) {
		return publicKey, nil
	}

	// Parses the token
	tkn, err := jwt.ParseWithClaims(tokenString, &jwt.RegisteredClaims{}, keyFunc)
	if err != nil {
		return uuid.UUID{}, err
	}

	// Checks the validity of the token
	if !tkn.Valid {
		fmt.Println(tokenString)
		return uuid.UUID{}, fmt.Errorf("Invalid token")
	}

	// Gets the claims from the token
	clms, ok := tkn.Claims.(*jwt.RegisteredClaims)
	if !ok {
		return uuid.UUID{}, fmt.Errorf("Claims not found")
	}

	// Parses the ID to UUID
	uid, err := uuid.Parse(clms.Subject)
	if err != nil {
		return uuid.UUID{}, nil
	}

	return uid, nil
}

func GetJWT(header http.Header) (string, error) {
	// Pulls token from request header and returns it as a string

	// Gets authorization from header
	authorization := header.Get("Authorization")
	if authorization == "" {
		return "", fmt.Errorf("No authorization found")
	}

	// Splits authorization into parts
	authParts := strings.Fields(authorization)

	// Checks if the authorization label is token
	if authParts[0] != "Token" {
		return "", fmt.Errorf("Incorrect authorization: %s should be Token", authParts[0])
	}

	// Checks if a token value was provided, returns it if so
	if len(authParts) <= 1 {
		return "", fmt.Errorf("No Token provided")
	}
	return authParts[1], nil
}
