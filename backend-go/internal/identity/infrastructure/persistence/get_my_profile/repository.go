package getmyprofile

import (
	database "backend/internal/database/sqlc"
	"context"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/sirupsen/logrus"
)

type Repository struct {
	q   *database.Queries
	log *logrus.Logger
}

func NewRepository(q *database.Queries, log *logrus.Logger) *Repository {
	return &Repository{
		q:   q,
		log: log,
	}
}

func (r *Repository) GetMe(ctx context.Context, userID pgtype.UUID) (database.User, error) {
	return r.q.GetUserByID(ctx, userID)
}
