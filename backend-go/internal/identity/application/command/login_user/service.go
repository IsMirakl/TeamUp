package loginuser

import (
	"context"

	database "backend/internal/database/sqlc"
	"backend/internal/identity/application/dto"
	"backend/internal/identity/domain/model"
	auth "backend/internal/pkg/utils"
	sharedErrors "backend/internal/shared/errors"

	"github.com/jackc/pgx/v5"
	"github.com/sirupsen/logrus"
)

type Repository interface {
	GetUserWithPasswordByEmail(ctx context.Context, email string) (database.GetUserWithPasswordByEmailRow, error)
}

type Service struct {
	repository Repository
	log        *logrus.Logger
}

func NewUserService(
	repository Repository, log *logrus.Logger) *Service {
	return &Service{
		repository: repository,
		log:        log,
	}
}

func (s *Service) Login(ctx context.Context, request *dto.LoginUserDTO) (string, error) {
	s.log.WithField("email", request.Email).Info("login request received")

	user, err := s.repository.GetUserWithPasswordByEmail(ctx, request.Email)
	if err != nil {
		if err == pgx.ErrNoRows {
			s.log.Warn("user not found")
			return "", sharedErrors.ErrInvalidCredentials
		}

		s.log.WithError(err).Error("failed to fetch user by email")
		return "", err
	}

	if !model.VerifyPassword(user.PasswordHash, request.Password) {
		s.log.Warn("invalid password")
		return "", sharedErrors.ErrInvalidCredentials
	}

	token, err := auth.CreateToken(user.UserID.String(), s.log)
	if err != nil {
		s.log.WithError(err).Error("failed to create token")
		return "", err
	}

	s.log.WithField("email", user.Email).Info("login successful")

	return token, nil
}
