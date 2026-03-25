package dto

import "backend/internal/features/user/model"


func ToUserResponse(user *model.User) *ResponseUserDTO {
	return &ResponseUserDTO {
		UserID: user.UserID,
		Email: user.Email,
		Name: user.Name,
		Avatar: user.Avatar,
	}
}