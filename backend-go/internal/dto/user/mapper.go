package user

import "backend/internal/models"

func ToUserResponse(user *models.User) *ResponseUserDTO {
	return &ResponseUserDTO{
		UserID: user.UserID,
		Email: user.Email,
		Name: user.Name,
		Avatar: user.Avatar,
	}
}