package registeruser

import (
	"backend/internal/features/user/dto"
	"backend/internal/features/user/model"

	"context"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type Service struct {
	db *gorm.DB
	repository Repository
	log *logrus.Logger
}

func NewUserService(db *gorm.DB, repository Repository, log *logrus.Logger) *Service {
	return &Service{
		db: db,
		repository: repository,
		log: log,
	}
}

func (s *Service) Create(ctx context.Context, dto *dto.CreateUserDTO) (*model.User, error) {
	
	tx := s.db.Begin()

	s.log.WithField("email", dto.Email).Info("Creating user")

	hash, err := model.HashPassword(dto.Password)
	if err != nil {
		tx.Rollback()
		
		s.log.WithError(err).Error("Failed to hash password")
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
			s.log.WithFields(logrus.Fields{
				"user_id": user.UserID,
				"email": user.Email,
			}).WithError(err).Error("Failed create user")

			return nil, err
		}

		s.log.WithFields(logrus.Fields{
			"user_id": user.UserID,
			"email": user.Email,
		}).Info("User successfully created")

		account := &model.Account{
			UserID: user.UserID,
			PasswordHash: hash,
			Provider: "local",
		}
		
		err = s.repository.CreateAccount(ctx, tx, account) 
		if err != nil {
			tx.Rollback()
			s.log.WithField("user_id", account.UserID).WithError(err).Error("Failed create account")

			return nil, err
		}

		if err := tx.Commit().Error; err != nil {
			return nil, err
		}

		s.log.WithField("user_id", account.UserID).Info("Account successfully created")
		
		return user, nil
}
