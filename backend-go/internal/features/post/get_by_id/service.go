package getbyid

import (
	"backend/internal/features/post/model"
	appErrors "backend/internal/shared/errors"
	"context"
	"errors"

	"gorm.io/gorm"
)

type Service struct {
	repository Repository
}

func NewService(repository Repository) *Service {
	return &Service{repository: repository}
}

func (s *Service) GetById(ctx context.Context, id string) (*model.Post, error) {
	post, err := s.repository.GetPostById(ctx, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, appErrors.ErrPostNotFound
		}
		return nil, err
	}

	if post == nil {
		return nil, appErrors.ErrPostNotFound
	}

	return post, nil
}
