package getpostresponses

import (
	"context"

	database "backend/internal/database/sqlc"
	"github.com/jackc/pgx/v5/pgtype"
)

type Repository struct {
	q *database.Queries
}

func NewRepository(q *database.Queries) *Repository {
	return &Repository{
		q: q,
	}
}

func (r *Repository) GetPostResponses(ctx context.Context, postID pgtype.UUID) ([]database.GetPostResponsesRow, error) {
	return r.q.GetPostResponses(ctx, postID)
}
