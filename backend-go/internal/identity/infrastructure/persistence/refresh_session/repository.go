package refreshsession

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
	return &Repository{q: q, log: log}
}

func (r *Repository) GetSessionByRefreshToken(ctx context.Context, refreshToken string) (database.Session, error) {
	return r.q.GetSessionByRefreshToken(ctx, refreshToken)
}

func (r *Repository) UpdateSessionRefreshToken(ctx context.Context, arg database.UpdateSessionRefreshTokenParams) (database.Session, error) {
	return r.q.UpdateSessionRefreshToken(ctx, arg)
}
