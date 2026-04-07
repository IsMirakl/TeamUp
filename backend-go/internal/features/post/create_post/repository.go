package createpost

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
	return &Repository{
		q: q,
		pool: pool,
	}
}

func (r *Repository) Create(ctx context.Context, arg database.CreatePostParams) (database.Post,error) {
	return r.q.CreatePost(ctx, arg)
}
