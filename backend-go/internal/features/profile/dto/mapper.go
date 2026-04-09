package dto

import database "backend/internal/database/sqlc"


func ToProfileResponse(user *database.User) *ProfileResponse {
	var avatar *string

	if user.Avatar.Valid {
		avatar = &user.Avatar.String
	}

	return &ProfileResponse{
		UserID: user.UserID.String(),
		Email:  user.Email,
		Name:   user.Name,
		Avatar: avatar,
		Role:   string(user.Role),
		SubscriptionPlan: string(user.SubscriptionPlan),

	}

}