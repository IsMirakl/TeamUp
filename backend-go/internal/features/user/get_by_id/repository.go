package getbyid

import (
	database "backend/internal/database/sqlc"
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

type Repository struct {
	q *database.Queries
}


func NewRepository(q *database.Queries) *Repository {
	return &Repository{q: q}
}

func (r *Repository) GetUserById(ctx context.Context, userID pgtype.UUID) (database.User, error) {
	return r.q.GetUserByID(ctx, userID)
}
