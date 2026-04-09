package dto

import (
	database "backend/internal/database/sqlc"
)


func ToUserResponse(user *database.User) *ResponseUserDTO {
	var avatar *string

	if user.Avatar.Valid {
		avatar = &user.Avatar.String
	}

	return &ResponseUserDTO{
		UserID: user.UserID.String(),
		Email:  user.Email,
		Name:   user.Name,
		Avatar: avatar,
	}
}