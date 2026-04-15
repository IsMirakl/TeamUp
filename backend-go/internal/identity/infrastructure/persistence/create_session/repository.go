package createsession

import (
	database "backend/internal/database/sqlc"
	"context"

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

func (r *Repository) CreateSession(ctx context.Context, arg database.CreateSessionParams) (database.Session, error) {
	return r.q.CreateSession(ctx, arg)
}
