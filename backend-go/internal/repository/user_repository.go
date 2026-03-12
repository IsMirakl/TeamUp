package repository

import (
	"backend/internal/models"
	"context"

	"gorm.io/gorm"
)

type UserRepository interface {
	Create(ctx context.Context, tx *gorm.DB, user *models.User) error
	CreateAccount(ctx context.Context, tx *gorm.DB, account *models.Account) error
	GetUserById(ctx context.Context, UserID uint, tx *gorm.DB,) (*models.User, error)
	GetUserByEmail(ctx context.Context, Email string, tx *gorm.DB)(*models.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}


func (r *userRepository) Create(ctx context.Context, tx *gorm.DB, user *models.User) (error) {
	
	if tx == nil {
		tx = r.db
	}

	return tx.WithContext(ctx).Create(user).Error
}

func (r *userRepository) CreateAccount(ctx context.Context, tx *gorm.DB, account *models.Account) (error) {
	
	if tx  == nil {
		tx = r.db
	}

	return tx.WithContext(ctx).Create(account).Error
}

func (r *userRepository) GetUserById(ctx context.Context, UserID uint, tx *gorm.DB) (*models.User, error) {

	if tx == nil {
		tx = r.db
	}

	var user models.User
	err := tx.WithContext(ctx).First(&user, UserID).Error


	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *userRepository) GetUserByEmail(ctx context.Context, email string, tx *gorm.DB) (*models.User, error) {
	
	if tx == nil {
		tx = r.db
	}

	var user models.User
	err := tx.WithContext(ctx).Where("email = ?", email).Take(&user).Error

	if err != nil {
		return nil, err
	}

	return &user, nil
	
}