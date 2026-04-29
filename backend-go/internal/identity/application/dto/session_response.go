package dto

type SessionResponse struct {
	ID           string `json:"id"`
	UserID       string `json:"user_id"`
	RefreshToken string `json:"refresh_token"`
	UserAgent    string `json:"user_agent"`
	ClientIp     string `json:"client_ip"`
	IsBlocked    bool   `json:"is_blocked"`
	ExpiresAt    string `json:"expires_at"`
}
