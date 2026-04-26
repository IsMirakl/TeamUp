package createpostresponse

import (
	"backend/internal/content/application/dto"
	database "backend/internal/database/sqlc"
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/sirupsen/logrus"
)

type Repository interface {
	Create(ctx context.Context, arg database.CreatePostResponseParams) (database.CreatePostResponseRow, error)
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

func (s *Service) Create(ctx context.Context, userID pgtype.UUID, dto *dto.CreatePostResponseDTO) (*database.CreatePostResponseRow, error) {
	postID, err := uuid.Parse(dto.PostID)
	if err != nil {
		s.log.WithError(err).WithField("post_id", dto.PostID).Error("Invalid post_id")
		return nil, err
	}

	params := database.CreatePostResponseParams{
		PostID: pgtype.UUID{
			Bytes: postID,
			Valid: true,
		},
		UserID:  userID,
		Message: dto.Message,
	}

	response, err := s.repository.Create(ctx, params)
	if err != nil {
		s.log.WithError(err).Error("failed to create post response")
		return nil, err
	}

	return &response, nil
}
