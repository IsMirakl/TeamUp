package getauthorpost

import (
	database "backend/internal/database/sqlc"
	"context"
)

type Repository struct {
	q *database.Queries
}

func NewRepository(q *database.Queries) *Repository {
	return &Repository{q: q}
}

func (r *Repository) GetAuthorPost(ctx context.Context, authorId string) (database.Post, error) {
	return r.q.GetAuthorPost(ctx, authorId)
}
