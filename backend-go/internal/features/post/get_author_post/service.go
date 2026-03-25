package getauthorpost

import (
	"backend/internal/features/post/model"
	"context"
)

type Service struct {
	repository Repository
}

func NewService(repository Repository) *Service {
	return &Service{repository: repository}
}

func (s *Service) GetAuthorPost(ctx context.Context, authorId string) ([]model.Post, error) {
	return s.repository.GetAuthorPost(ctx, authorId)
}
