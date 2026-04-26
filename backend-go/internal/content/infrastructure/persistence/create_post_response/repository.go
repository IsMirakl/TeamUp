package createpostresponse

import (
	database "backend/internal/database/sqlc"
	"context"
)

type Repository struct {
	q *database.Queries
}

func NewRepository(q *database.Queries) *Repository {
	return &Repository{
		q: q,
	}
}

func (r *Repository) Create(ctx context.Context, arg database.CreatePostResponseParams) (database.CreatePostResponseRow, error) {
	return r.q.CreatePostResponse(ctx, arg)
}
