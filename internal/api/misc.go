package api

import (
	"context"
	"database/sql"
	"time"

	"github.com/google/uuid"
	"github.com/roxensox/dailychapter/internal/auth"
	"github.com/roxensox/dailychapter/internal/database"
)

func (cfg *ApiConfig) InsertRefreshToken(user_id uuid.UUID, ctxt context.Context) (string, error) {
	tkn, err := auth.MakeRefreshToken()
	if err != nil {
		return "", err
	}
	params := database.CreateRefreshTokenParams{
		ID:         uuid.New(),
		CreatedAt:  time.Now().UTC(),
		UpdatedAt:  time.Now().UTC(),
		Token:      tkn,
		UserID:     user_id,
		ValidUntil: time.Now().Add(time.Hour * 24 * 60).UTC(),
		RevokedAt:  sql.NullTime{},
	}
	resp, err := cfg.DBConn.CreateRefreshToken(ctxt, params)
	if err != nil {
		return "", err
	}
	return resp, nil
}
