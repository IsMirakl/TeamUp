package registeruser

import (
	"backend/internal/features/user/model"
	"backend/internal/features/user/dto"

	"context"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Service struct {
	db *gorm.DB
	repository Repository
}

func NewUserService(db *gorm.DB, repository Repository) *Service {
	return &Service{
		db: db,
		repository: repository,
	}
}

func (s *Service) Create(ctx context.Context, dto *dto.CreateUserDTO) (*model.User, error) {
	
	tx := s.db.Begin()


	hash, err := model.HashPassword(dto.Password)
	if err != nil {
		tx.Rollback()
		return nil, err
	}


	user := &model.User{
		UserID: uuid.NewString(),
		Email: dto.Email,
		Name: dto.Name,
		Avatar: dto.Avatar,
	}

	err = s.repository.Create(ctx, tx, user) 
		if err != nil {
			tx.Rollback()
			return nil, err
		}

		account := &model.Account{
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
