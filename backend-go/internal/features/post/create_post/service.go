package createpost

import (
	"backend/internal/features/post/dto"
	"backend/internal/features/post/model"
	"context"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Service struct {
	db         *gorm.DB
	repository Repository
}

func NewService(db *gorm.DB, repository Repository) *Service {
	return &Service{
		db:         db,
		repository: repository,
	}
}

func (s *Service) Create(ctx context.Context, request *dto.CreatePostDTO, userID string) (*model.Post, error) {
	tx := s.db.Begin()

	post := &model.Post{
		ID:          uuid.NewString(),
		Title:       request.Title,
		Description: request.Description,
		Tags:        request.Tags,
		AuthorID:    userID,
	}

	err := s.repository.Create(ctx, tx, post)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		return nil, err
	}

	return post, nil
}
