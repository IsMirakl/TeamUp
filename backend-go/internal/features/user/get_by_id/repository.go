package getbyid

import (
	"backend/internal/features/user/model"
	"context"

	"gorm.io/gorm"
)

type Repository interface {
	GetUserById(ctx context.Context, userID string) (*model.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &userRepository{db: db}
}

func (r *userRepository) GetUserById(ctx context.Context, userID string) (*model.User, error) {
	var user model.User

	err := r.db.WithContext(ctx).First(&user, userID).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}
