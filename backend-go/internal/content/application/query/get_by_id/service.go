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
	GetPostById(ctx context.Context, id pgtype.UUID) (database.GetPostByIdRow, error)
}

func NewService(repository Repository, log *logrus.Logger) *Service {
	return &Service{repository: repository, log: log}
}

func (s *Service) GetById(ctx context.Context, id string) (*database.GetPostByIdRow, error) {
	s.log.WithField("post_id", id).Info("GetById called")

	postID, err := uuid.Parse(id)
	if err != nil {
		s.log.WithError(err).
			WithField("userID", id).
			Error("failed to parse userID")

		return &database.GetPostByIdRow{}, err
	}

	pgID := pgtype.UUID{
		Bytes: postID,
		Valid: true,
	}

	post, err := s.repository.GetPostById(ctx, pgID)
	if err != nil {
		s.log.WithError(err).
			WithField("post_id", id).
			Error("failed to get post from repository")

		return &database.GetPostByIdRow{}, err
	}

	s.log.WithField("post_id", id).Info("post fetched successfully")

	return &post, nil
}
