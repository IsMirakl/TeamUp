package getbyid

import (
	database "backend/internal/database/sqlc"
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/sirupsen/logrus"
)

type Service struct {
	repository Repository
	log        *logrus.Logger
}

type Repository interface {
	GetUserById(ctx context.Context, userID pgtype.UUID) (database.User, error)
}

func NewService(repository Repository, log *logrus.Logger) *Service {
	return &Service{
		repository: repository,
		log:        log,
	}
}

func (s *Service) GetById(ctx context.Context, userID string) (*database.User, error) {
	s.log.WithField("userID", userID).Info("GetById called")

	id, err := uuid.Parse(userID)
	if err != nil {
		s.log.WithError(err).
			WithField("userID", userID).
			Error("failed to parse userID")

		return &database.User{}, err
	}

	pgID := pgtype.UUID{
		Bytes: id,
		Valid: true,
	}

	user, err := s.repository.GetUserById(ctx, pgID)
	if err != nil {
		s.log.WithError(err).
			WithField("userID", userID).
			Error("failed to get user from repository")

		return &database.User{}, err
	}

	s.log.WithField("userID", userID).Info("user fetched successfully")

	return &user, nil
}
