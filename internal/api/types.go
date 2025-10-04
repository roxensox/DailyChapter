package api

import (
	"github.com/roxensox/dailychapter/internal/database"
)

type ApiConfig struct {
	DBConn *database.Queries
}
