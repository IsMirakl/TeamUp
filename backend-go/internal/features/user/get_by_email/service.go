package getbyemail

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

func (s *Service) GetByEmail(ctx context.Context, email string) (*model.User, error) {
	return s.repository.GetUserByEmail(ctx, email)
}
