package service

import (
	"backend/internal/auth"
	"backend/internal/dto/user"
	"backend/internal/models"
	"backend/internal/repository"
	"context"
	"errors"

	"gorm.io/gorm"
)

type UserService struct {
	db *gorm.DB
	repository repository.UserRepository
}

func NewUserService(db *gorm.DB, repository repository.UserRepository) *UserService {
	return &UserService{
		db: db,
		repository: repository,
	}
}

func (s *UserService) Create(ctx context.Context, dto *user.CreateUserDTO) (*models.User, error) {
	
	tx := s.db.Begin()


	hash, err := models.HashPassword(dto.Password)
	if err != nil {
		tx.Rollback()
		return nil, err
	}


	user := &models.User{
		Email: dto.Email,
		Name: dto.Name,
		Avatar: dto.Avatar,
	}

	err = s.repository.Create(ctx, tx, user) 
		if err != nil {
			tx.Rollback()
			return nil, err
		}

		account := &models.Account{
			UserID: user.UserID,
			PasswordHash: hash,
			Provider: "local",
		}
		
		err = s.repository.CreateAccount(ctx, tx, account) 
		if err != nil {
			tx.Rollback()
			return nil, err
		}

		if err := tx.Commit().Error; err != nil {
			return nil, err
		}
		
		return user, nil
}



func (s *UserService) GetUserById(ctx context.Context, UserID uint) (*models.User, error) {
	return s.repository.GetUserById(ctx, UserID, nil)
}

// func (s *UserService) GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
// 	return s.repository.GetUserByEmail(ctx, email, nil)
// }

func (s *UserService) Login(ctx context.Context, dto *user.LoginUserDTO) (string, error) {
	
	tx := s.db.Begin()

	user, err := s.repository.GetUserByEmail(ctx, dto.Email, tx)

	if err != nil {
		tx.Rollback()
		return "", err
	}

	ok := models.VerifyPassword(user.Account.PasswordHash, dto.Password)
	if ok != nil{
		tx.Rollback()
		return "", errors.New("")
	}

	token, err := auth.CreateToken(uint64(user.UserID))
	if err != nil {
		tx.Rollback()
		return "", err
	}

	if err := tx.Commit().Error; err != nil {
		return "", err
	}
		

	return token, nil

}