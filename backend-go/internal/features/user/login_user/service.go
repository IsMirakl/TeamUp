package loginuser

import (
	"backend/internal/features/user/dto"
	"backend/internal/features/user/model"
	auth "backend/internal/pkg/utils"
	"context"

	"github.com/sirupsen/logrus"
)

type UserFetcher interface {
	GetByEmail(ctx context.Context, email string) (*model.User, error)
}

type Service struct {
	userFetcher UserFetcher
	log         *logrus.Logger
}

func NewUserService(userFetcher UserFetcher, log *logrus.Logger) *Service {
	return &Service{userFetcher: userFetcher, log: log}
}

func (s *Service) Login(ctx context.Context, request *dto.LoginUserDTO) (string, error) {

	s.log.WithField("email", request.Email).Info("login request received")

	user, err := s.userFetcher.GetByEmail(ctx, request.Email)
	if err != nil {
		s.log.WithError(err).Error("failed to fetch user by email")

		return "", err
	}

	if user == nil || user.Account == nil {
		s.log.Warn("user not found or invalid data")

		return "", nil
	}

	if !model.VerifyPassword(user.Account.PasswordHash, request.Password) {
		s.log.Warn("invalid password")

		return "", nil
	}

	token, err := auth.CreateToken(user.UserID, s.log)
	if err != nil {
		s.log.WithError(err).Error("failed to create token")
		
		return "", err
	}

	return token, nil
}
