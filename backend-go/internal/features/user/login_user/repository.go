package loginuser

import (
	database "backend/internal/database/sqlc"
	"context"
)

type Repository struct {
	q *database.Queries
}

func NewRepository(q *database.Queries) *Repository{
	return &Repository{
		q: q,
	}
}


func (r *Repository) GetUserWithPasswordByEmail(ctx context.Context, email string) (database.GetUserWithPasswordByEmailRow, error) {
	return r.q.GetUserWithPasswordByEmail(ctx, email)
}