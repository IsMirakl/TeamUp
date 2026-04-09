package updatepost

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

func NewService(repository *Repository, log *logrus.Logger) *Service {
	return &Service{
		repository: repository,
		log: log,
	}
}

func (s *Service) Update(ctx context.Context, postID string ,dto *dto.UpdatePostDTO) (*database.UpdatePostRow, error) {
	parsedID, err := uuid.Parse(postID)
	if err != nil {
		s.log.WithField("post_id", postID).WithError(err).Error("invalid post id")
		return nil, err
	}

	pgID := pgtype.UUID{
		Bytes: parsedID,
		Valid: true,
	}
	
	tx, err := s.repository.pool.Begin(ctx)
	if err != nil {
		s.log.WithError(err).Error("failed to begin transaction")
		return nil, err
	}
	defer func() {
		_ = tx.Rollback(ctx)
	}()

	qtx := s.repository.q.WithTx(tx)

	s.log.WithField("post_id", postID).Info("updating post")

	post, err := qtx.UpdatePost(ctx, database.UpdatePostParams{
		ID: pgID,
		Title: dto.Title,
		Description: dto.Description,
		Tags: dto.Tags,
	})

	if err != nil {
		s.log.WithField("post_id", postID).WithError(err).Error("failed update post");
		return nil, err
	}

	if err := tx.Commit(ctx); err != nil {
		s.log.WithField("post_id", postID).WithError(err).Error("failed to commit transaction")
		return nil, err
	}

	return &post, nil
}
