package user

type ResponseUserDTO struct {
	ID uint `json:"id"`
	Name string `json:"name"`
	Email string `json:"email"`
	Avatar *string `json:"avatar"`
	Role string `json:"role"`
	SubscriptionPlan string `json:"subscriptionPlan"`	
}
