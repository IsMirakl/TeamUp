package getbyid

import (
	database "backend/internal/database/sqlc"
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

type Service struct {
	repository *Repository
}

func NewService(repository *Repository) *Service {
	return &Service{
		repository: repository,
	}
}

func (s *Service) GetById(ctx context.Context, userID string) (database.User, error) {
	id, err := uuid.Parse(userID)
	if err != nil {
		return database.User{}, err
	}

	pgID := pgtype.UUID{
		Bytes: id,
		Valid: true,
	}

	return s.repository.GetUserById(ctx, pgID)
}