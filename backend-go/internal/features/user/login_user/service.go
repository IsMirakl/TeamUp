package loginuser

import (
	"backend/internal/features/user/dto"
	"backend/internal/features/user/model"
	auth "backend/internal/pkg/utils"
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/sirupsen/logrus"
)


type Service struct {
	repository *Repository
	log        *logrus.Logger
}

func NewUserService(
	repository *Repository, log *logrus.Logger) *Service {
	return &Service{
		repository: repository,
		log:       log,
	}
}

func (s *Service) Login(ctx context.Context, request *dto.LoginUserDTO) (string, error) {
	s.log.WithField("email", request.Email).Info("login request received")

	user, err := s.repository.GetUserWithPasswordByEmail(ctx, request.Email)
	if err != nil {
		if err == pgx.ErrNoRows {
			s.log.Warn("user not found")
			return "", nil
		}

		s.log.WithError(err).Error("failed to fetch user by email")
		return "", err
	}

	if !model.VerifyPassword(user.PasswordHash, request.Password) {
		s.log.Warn("invalid password")
		return "", nil
	}

	token, err := auth.CreateToken(user.UserID.String(), s.log)
	if err != nil {
		s.log.WithError(err).Error("failed to create token")
		return "", err
	}

	return token, nil
}