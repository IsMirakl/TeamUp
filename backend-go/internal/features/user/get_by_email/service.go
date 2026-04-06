package getbyemail

import (
	database "backend/internal/database/sqlc"
	"context"
)

type Service struct {
	repository *Repository
}

func NewService(repository *Repository) *Service {
	return &Service{repository: repository}
}

func (s *Service) GetByEmail(ctx context.Context, email string) (database.User, error) {
	return s.repository.GetUserByEmail(ctx, email)
}
