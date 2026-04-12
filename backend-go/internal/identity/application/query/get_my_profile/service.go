package getmyprofile

import (
	database "backend/internal/database/sqlc"
	"context"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/sirupsen/logrus"
)



type Repository interface {
	GetMe(ctx context.Context, userID pgtype.UUID) (database.User, error)
}

type Service struct {
	repository Repository
	log *logrus.Logger
}

func NewPostService(repository Repository, log *logrus.Logger) *Service {
	return &Service{
		repository: repository,
		log: log,
	}
}

func (s *Service) GetMe(ctx context.Context, userID pgtype.UUID) (*database.User, error) {
	s.log.Infof("Getting user profile for userID: %v", userID)
	user, err := s.repository.GetMe(ctx, userID)
	if err != nil {
		s.log.Errorf("Failed to get user: %v", err)
		return nil, err
	}
	s.log.Infof("Successfully retrieved user profile: %v", user)
	return &user, nil
}