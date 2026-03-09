package service

import (
	"backend/internal/models"
	"backend/internal/repository"
	"context"
)

type UserService struct {
	repository repository.UserRepository
}

func NewUserService(repository repository.UserRepository) *UserService {
	return &UserService{
		repository: repository,
	}
}

func (s *UserService) GetUserById(ctx context.Context, UserID uint) (*models.User, error) {
	return s.repository.GetUserById(ctx, UserID, nil)
}