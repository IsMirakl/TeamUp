package revokesession

import (
	database "backend/internal/database/sqlc"
	"backend/internal/shared/errors"
	"context"

	"github.com/sirupsen/logrus"
)

type Repository interface {
	RevokeSessionByRefreshToken(ctx context.Context, refreshToken string) (database.Session, error)
}

type Service struct {
	repository Repository
	log        *logrus.Logger
}

func NewService(repository Repository, log *logrus.Logger) *Service {
	return &Service{
		repository: repository,
		log:        log,
	}
}

func (s *Service) RevokeSession(ctx context.Context, refreshToken string) error {
	if refreshToken == "" {
		s.log.Error("refresh token is empty")
		return errors.ErrEmptyRefreshToken
	}

	_, err := s.repository.RevokeSessionByRefreshToken(ctx, refreshToken)
	if err != nil {
		s.log.WithError(err).Error("failed to revoke session")
		return err
	}

	return nil
}
