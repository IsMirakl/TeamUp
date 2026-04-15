package dto

type CreateUserDTO struct {
	Email    string  `json:"email" validate:"required,email"`
	Name     string  `json:"name" validate:"required,max=25"`
	Avatar   *string `json:"avatar"`
	Password string  `json:"password" validate:"required,min=8"`
	UserAgent string  `json:"-"`
	ClientIP  string  `json:"-"`
}

type LoginUserDTO struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
	UserAgent string `json:"-"`
	ClientIP  string `json:"-"`
}

type GetUserByEmailDTO struct {
	Email string `json:"email" validate:"required,email"`
}