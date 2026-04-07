package updatepost

import (
	database "backend/internal/database/sqlc"
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Repository struct {
	q *database.Queries
	pool *pgxpool.Pool
}

func NewRepository(q *database.Queries, pool *pgxpool.Pool) *Repository {
	return &Repository{q: q, pool: pool}
}

func (r *Repository) Update(ctx context.Context, arg database.UpdatePostParams) (database.Post ,error) {
	return r.q.UpdatePost(ctx, arg)
}
