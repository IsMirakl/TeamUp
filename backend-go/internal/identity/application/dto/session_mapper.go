package dto

import (
	database "backend/internal/database/sqlc"
	"time"
)

func ToSessionMapper(session *database.Session) *SessionResponse {
	expiresAt := ""
	if session.ExpiresAt.Valid {
		expiresAt = session.ExpiresAt.Time.UTC().Format(time.RFC3339Nano)
	}

	return &SessionResponse{
		UserID:       session.UserID.String(),
		RefreshToken: session.RefreshToken,
		UserAgent:    session.UserAgent,
		ClientIp:     session.ClientIp,
		IsBlocked:    session.IsBlocked,
		ExpiresAt:    expiresAt,
	}
}
