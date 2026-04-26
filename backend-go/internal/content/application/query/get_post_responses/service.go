package getpostresponses

import (
	"context"
	"errors"
	"fmt"

	database "backend/internal/database/sqlc"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/sirupsen/logrus"
)

var ErrInvalidPostID = errors.New("invalid post_id")

type Service struct {
	repository Repository
	log        *logrus.Logger
}

type Repository interface {
	GetPostResponses(ctx context.Context, postID pgtype.UUID) ([]database.GetPostResponsesRow, error)
}

func NewService(repository Repository, log *logrus.Logger) *Service {
	return &Service{
		repository: repository,
		log:        log,
	}
}

func (s *Service) GetPostResponses(ctx context.Context, postID string) ([]database.GetPostResponsesRow, error) {
	s.log.WithField("post_id", postID).Info("GetPostResponses called")

	parsedID, err := uuid.Parse(postID)
	if err != nil {
		s.log.WithError(err).WithField("post_id", postID).Error("failed to parse post_id")
		return nil, fmt.Errorf("%w: %v", ErrInvalidPostID, err)
	}

	pgID := pgtype.UUID{Bytes: parsedID, Valid: true}

	responses, err := s.repository.GetPostResponses(ctx, pgID)
	if err != nil {
		s.log.WithError(err).WithField("post_id", postID).Error("failed to get post responses")
		return nil, err
	}

	return responses, nil
}
