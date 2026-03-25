package getbyid

import (
	"backend/internal/features/user/model"
	"context"
)

type Service struct {
	repository Repository
}

func NewService(repository Repository) *Service {
	return &Service{repository: repository}
}

func (s *Service) GetById(ctx context.Context, userID string) (*model.User, error) {
	return s.repository.GetUserById(ctx, userID)
}
