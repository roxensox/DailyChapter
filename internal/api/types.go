package api

import (
	"crypto/rsa"

	"github.com/roxensox/dailychapter/internal/database"
)

type ApiConfig struct {
	DBConn     *database.Queries
	PrivateKey *rsa.PrivateKey
	PublicKey  *rsa.PublicKey
	Secret     string
	APIKey     string
}
