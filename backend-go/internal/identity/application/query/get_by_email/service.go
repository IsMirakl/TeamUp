package getbyemail

import (
	database "backend/internal/database/sqlc"
	"context"

	"github.com/sirupsen/logrus"
)

type Service struct {
	repository Repository
	log        *logrus.Logger
}

type Repository interface {
	GetUserByEmail(ctx context.Context, email string) (database.User, error)
}

func NewService(repository Repository, log *logrus.Logger) *Service {
	return &Service{repository: repository, log: log}
}

func (s *Service) GetByEmail(ctx context.Context, email string) (database.User, error) {
	s.log.WithField("email", email).Info("getting user by email")

	user, err := s.repository.GetUserByEmail(ctx, email)
	if err != nil {
		s.log.WithError(err).Warn("failed to get user by email")
		return database.User{}, err
	}

	return user, nil
}
