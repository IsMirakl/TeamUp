package registeruser

import (
	database "backend/internal/database/sqlc"
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type userRepository struct {
	q    *database.Queries
	pool *pgxpool.Pool
}

func NewUserRepository(q *database.Queries, pool *pgxpool.Pool) *userRepository {
	return &userRepository{
		q:    q,
		pool: pool,
	}
}

func (r *userRepository) Create(ctx context.Context, arg database.CreateUserParams) (database.User, error) {
	return r.q.CreateUser(ctx, arg)
}

func (r *userRepository) CreateAccount(ctx context.Context, arg database.CreateAccountParams) (error) {
	return r.q.CreateAccount(ctx, arg)
}