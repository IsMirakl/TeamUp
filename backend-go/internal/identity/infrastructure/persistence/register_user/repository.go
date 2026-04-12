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

func (r *userRepository) CreateAccount(ctx context.Context, arg database.CreateAccountParams) error {
	return r.q.CreateAccount(ctx, arg)
}

func (r *userRepository) CreateWithAccount(ctx context.Context, userParams database.CreateUserParams, accountParams database.CreateAccountParams) (database.User, error) {
	tx, err := r.pool.Begin(ctx)
	if err != nil {
		return database.User{}, err
	}
	defer func() {
		_ = tx.Rollback(ctx)
	}()

	qtx := r.q.WithTx(tx)

	user, err := qtx.CreateUser(ctx, userParams)
	if err != nil {
		return database.User{}, err
	}

	if err := qtx.CreateAccount(ctx, accountParams); err != nil {
		return database.User{}, err
	}

	if err := tx.Commit(ctx); err != nil {
		return database.User{}, err
	}

	return user, nil
}
