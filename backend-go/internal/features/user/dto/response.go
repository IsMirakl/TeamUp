package dto

type ResponseUserDTO struct {
	UserID string `json:"user_id"`
	Name string `json:"name"`
	Email string `json:"email"`
	Avatar *string `json:"avatar"`
	Role string `json:"role"`
	SubscriptionPlan string `json:"subscriptionPlan"`	
}

type LoginResponse struct {
	AccessToken string `json:"accessToken"`
}
