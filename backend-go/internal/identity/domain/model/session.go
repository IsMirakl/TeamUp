package model

import (
	"time"
)

type Session struct {
	ID string
	UserID string

	RefreshToken string
	UserAgent string
	ClientIp string
	IsBlocked bool

	ExpiresAt time.Time
	CreatedAt time.Time
	RevokedAt time.Time
}