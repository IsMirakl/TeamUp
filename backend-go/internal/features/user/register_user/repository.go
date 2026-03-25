package registeruser

import (
	"backend/internal/features/user/model"
	"context"

	"gorm.io/gorm"
)

type Repository interface {
	Create(ctx context.Context, tx *gorm.DB, user *model.User) error
	CreateAccount(ctx context.Context, tx *gorm.DB, account *model.Account) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) Repository {
	return &userRepository{db: db}
}


func (r *userRepository) Create(ctx context.Context, tx *gorm.DB, user *model.User) (error) {
	
	if tx == nil {
		tx = r.db
	}

	return tx.WithContext(ctx).Create(user).Error
}

func (r *userRepository) CreateAccount(ctx context.Context, tx *gorm.DB, account *model.Account) (error) {
	
	if tx  == nil {
		tx = r.db
	}

	return tx.WithContext(ctx).Create(account).Error
}