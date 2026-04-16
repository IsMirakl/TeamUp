package getauthorpost

import (
	database "backend/internal/database/sqlc"
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

type Repository struct {
	q *database.Queries
}

func NewRepository(q *database.Queries) *Repository {
	return &Repository{q: q}
}

func (r *Repository) GetAuthorPost(ctx context.Context, authorId string) (database.Post, error) {
	authorUUID, err := uuid.Parse(authorId)
	if err != nil {
		return database.Post{}, err
	}
	return r.q.GetAuthorPost(ctx, pgtype.UUID{Bytes: authorUUID, Valid: true})
}
