package updatepost

import (
	"backend/internal/features/post/dto"
	"backend/internal/features/post/model"
	"context"

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

func (s *Service) Update(ctx context.Context, id string, request *dto.UpdatePostDTO) (*model.Post, error) {
	tx := s.db.Begin()

	post := &model.Post{
		ID:          id,
		Title:       request.Title,
		Description: request.Description,
		Tags:        request.Tags,
	}

	err := s.repository.Update(ctx, tx, post)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		return nil, err
	}

	return post, nil
}
