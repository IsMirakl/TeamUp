package dto

import "time"


type CreateSessionDTO struct {
	ID string `json:"id" validate:"required"`
	UserID string `json:"user_id" validate:"required"`
	RefreshToken string `json:"refresh_token" validate:"required"`
	UserAgent string `json:"user_agent" validate:"required"`
	ClientIp string `json:"client_ip" validate:"required"`
	IsBlocked bool `json:"is_blocked" validate:"required"`
	ExpiresAt time.Time `json:"expires_at" validate:"required"`
}