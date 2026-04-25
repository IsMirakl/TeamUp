package getmyprofile

import (
	database "backend/internal/database/sqlc"
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/sirupsen/logrus"
)

type Repository interface {
	GetMe(ctx context.Context, userID pgtype.UUID) (database.User, error)
}

type Service struct {
	repository Repository
	log        *logrus.Logger
}

func NewPostService(repository Repository, log *logrus.Logger) *Service {
	return &Service{
		repository: repository,
		log:        log,
	}
}

func (s *Service) GetMe(ctx context.Context, userID string) (*database.User, error) {
	s.log.WithField("userID", userID).Info("GetMe called")

	id, err := uuid.Parse(userID)
	if err != nil {
		s.log.WithError(err).WithField("userID", userID).Error("failed to parse userID")
		return &database.User{}, err
	}

	pgID := pgtype.UUID{Bytes: id, Valid: true}

	user, err := s.repository.GetMe(ctx, pgID)
	if err != nil {
		s.log.WithError(err).WithField("userID", userID).Error("failed to get profile from repository")
		return &database.User{}, err
	}

	s.log.WithField("userID", userID).Info("profile fetched successfully")
	return &user, nil
}
