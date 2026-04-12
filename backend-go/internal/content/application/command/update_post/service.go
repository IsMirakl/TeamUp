package updatepost

import (
	"context"

	"backend/internal/content/application/dto"
	database "backend/internal/database/sqlc"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/sirupsen/logrus"
)

type Service struct {
	repository Repository
	log        *logrus.Logger
}

type Repository interface {
	Update(ctx context.Context, arg database.UpdatePostParams) (database.UpdatePostRow, error)
}

func NewService(repository Repository, log *logrus.Logger) *Service {
	return &Service{
		repository: repository,
		log:        log,
	}
}

func (s *Service) Update(ctx context.Context, postID string, dto *dto.UpdatePostDTO) (*database.UpdatePostRow, error) {
	parsedID, err := uuid.Parse(postID)
	if err != nil {
		s.log.WithField("post_id", postID).WithError(err).Error("invalid post id")
		return nil, err
	}

	pgID := pgtype.UUID{
		Bytes: parsedID,
		Valid: true,
	}

	s.log.WithField("post_id", postID).Info("updating post")

	post, err := s.repository.Update(ctx, database.UpdatePostParams{
		ID:          pgID,
		Title:       dto.Title,
		Description: dto.Description,
		Tags:        dto.Tags,
	})

	if err != nil {
		s.log.WithField("post_id", postID).WithError(err).Error("failed update post")
		return nil, err
	}

	return &post, nil
}
