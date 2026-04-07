package createpost

import (
	database "backend/internal/database/sqlc"
	"backend/internal/features/post/dto"
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/sirupsen/logrus"
)

type Service struct {
	repository *Repository
	log *logrus.Logger
}

func NewService(repository Repository, log *logrus.Logger) *Service {
	return &Service{
		repository: &repository,
		log: log,
	}
}

func (s *Service) Create(ctx context.Context, dto *dto.CreatePostDTO, userID string) (*database.Post, error) {
	tx, err := s.repository.pool.Begin(ctx)
	if err != nil {
		s.log.WithError(err).Error("failed to begin transaction")
		return nil, err
	}
	defer func() {
		_ = tx.Rollback(ctx)
	}()
	
	
	qtx := s.repository.q.WithTx(tx)
	ID := uuid.New()
	
	s.log.WithField("title_post", dto.Title).Info("Creating post")

	post, err := qtx.CreatePost(ctx, database.CreatePostParams{
		ID:          pgtype.UUID{Bytes: ID, Valid: true},
		Title:       dto.Title,
		Description: dto.Description,
		Tags:        dto.Tags,
		AuthorID:    userID,
	})
	if err != nil {
		s.log.WithField("post_ID", post.ID.String()).WithError(err).Error("Failed create post")
		return nil, err
	}

	if err := tx.Commit(ctx); err != nil {
		s.log.WithField("post_id", ID.String()).
		WithError(err).
		Error("failed to commit transacton")
		return nil, err
	}

	s.log.WithField("post_ID", ID.String()).Info("Post successfully created")

	return &post, nil
}
