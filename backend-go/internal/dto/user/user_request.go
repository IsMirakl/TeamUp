package user

type CreateUserDTO struct {
	Email string `json:"email" validate:"required,email"`
	Name string `json:"name" validate:"required,max=25"`
	Avatar *string `json:"avatar"`
	Password string `json:"password" validate:"required"`
}

type LoginUserDTO struct {
	UserID string `json:"user_id" validate:"required"`
	Password string `json:"password" validate:"required"`
}