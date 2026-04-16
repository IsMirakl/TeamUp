package createpost

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
	Create(ctx context.Context, arg database.CreatePostParams) (database.Post, error)
}

func NewService(repository Repository, log *logrus.Logger) *Service {
	return &Service{
		repository: repository,
		log:        log,
	}
}

func (s *Service) Create(ctx context.Context, dto *dto.CreatePostDTO, userID string) (*database.Post, error) {
	ID := uuid.New()
	authorUUID, err := uuid.Parse(userID)
	if err != nil {
		s.log.WithField("user_id", userID).WithError(err).Error("invalid user_id")
		return nil, err
	}

	s.log.WithField("title_post", dto.Title).Info("Creating post")

	post, err := s.repository.Create(ctx, database.CreatePostParams{
		ID:          pgtype.UUID{Bytes: ID, Valid: true},
		Title:       dto.Title,
		Description: dto.Description,
		Tags:        dto.Tags,
		AuthorID:    pgtype.UUID{Bytes: authorUUID, Valid: true},
	})
	if err != nil {
		s.log.WithField("post_ID", post.ID.String()).WithError(err).Error("Failed create post")
		return nil, err
	}

	s.log.WithField("post_ID", ID.String()).Info("Post successfully created")

	return &post, nil
}
