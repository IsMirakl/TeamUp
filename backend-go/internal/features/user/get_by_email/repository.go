package getbyemail

import (
	"backend/internal/features/user/model"
	"context"

	"gorm.io/gorm"
)

type Repository interface {
	GetUserByEmail(ctx context.Context, email string) (*model.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &userRepository{db: db}
}

func (r *userRepository) GetUserByEmail(ctx context.Context, email string) (*model.User, error) {
	var user model.User

	err := r.db.WithContext(ctx).Preload("Account").Where("email = ?", email).Take(&user).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}
