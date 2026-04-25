package revokesession

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

func (r *Repository) RevokeSessionByRefreshToken(ctx context.Context, refreshToken string) (database.Session, error) {
	session, _ := r.q.RevokeSessionByRefreshToken(ctx, refreshToken)

	return session, nil
}
