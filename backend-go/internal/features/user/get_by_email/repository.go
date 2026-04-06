package getbyemail

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

func (r *Repository) GetUserByEmail(ctx context.Context, email string) (database.User, error) {
	return r.q.GetUserByEmail(ctx, email)
}
