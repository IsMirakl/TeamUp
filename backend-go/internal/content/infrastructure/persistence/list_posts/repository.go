package listposts

import (
	"context"

	database "backend/internal/database/sqlc"
)

type Repository struct {
	q *database.Queries
}

func NewRepository(q *database.Queries) *Repository {
	return &Repository{q: q}
}

func (r *Repository) ListPosts(ctx context.Context, limit, offset int32) ([]database.Post, error) {
	return r.q.ListPosts(ctx, database.ListPostsParams{Limit: limit, Offset: offset})
}
