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

func (r *Repository) GetPostById(ctx context.Context, id pgtype.UUID) (database.Post, error) {
	return r.q.GetPostById(ctx, id)
}
