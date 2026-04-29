package dto

import "time"

type ResponseUserDTO struct {
	UserID           string  `json:"user_id"`
	Name             string  `json:"name"`
	Email            string  `json:"email"`
	Avatar           *string `json:"avatar"`
	Role             string  `json:"role"`
	SubscriptionPlan string  `json:"subscriptionPlan"`
}

type LoginResponse struct {
	SessionId    string    `json:"session_id"`
	AccessToken  string    `json:"accessToken"`
	RefreshToken string    `json:"refresh_token"`
	IsBlocked    bool      `json:"is_blocked"`
	ExpiresAt    time.Time `json:"expires_at"`
}

type RegisterResponse struct {
	User         *ResponseUserDTO `json:"user"`
	AccessToken  string           `json:"access_token"`
	RefreshToken string           `json:"refresh_token"`
}
