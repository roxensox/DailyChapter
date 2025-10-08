package api

import (
	"crypto/rsa"

	"github.com/roxensox/dailychapter/internal/database"
)

type ApiConfig struct {
	DBConn     *database.Queries
	PrivateKey *rsa.PrivateKey
	Secret     string
	APIKey     string
}
