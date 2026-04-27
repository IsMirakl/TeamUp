package getbyemail_test

import (
	"context"
	"errors"
	"testing"

	database "backend/internal/database/sqlc"
	getbyemail "backend/internal/identity/application/query/get_by_email"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type stubGetByEmailRepo struct {
	called bool
	got    string

	fn func(ctx context.Context, email string) (database.User, error)
}

func (s *stubGetByEmailRepo) GetUserByEmail(ctx context.Context, email string) (database.User, error) {
	s.called = true
	s.got = email
	return s.fn(ctx, email)
}

func TestGetByEmail_Success(t *testing.T) {
	repo := &stubGetByEmailRepo{fn: func(ctx context.Context, email string) (database.User, error) {
		return database.User{Email: email, Name: "Name"}, nil
	}}

	svc := getbyemail.NewService(repo, logrus.New())
	user, err := svc.GetByEmail(context.Background(), "test@example.com")

	require.NoError(t, err)
	assert.True(t, repo.called)
	assert.Equal(t, "test@example.com", repo.got)
	assert.Equal(t, "test@example.com", user.Email)
}

func TestGetByEmail_RepositoryError(t *testing.T) {
	expErr := errors.New("db error")
	repo := &stubGetByEmailRepo{fn: func(ctx context.Context, email string) (database.User, error) {
		return database.User{}, expErr
	}}

	svc := getbyemail.NewService(repo, logrus.New())
	user, err := svc.GetByEmail(context.Background(), "test@example.com")

	assert.ErrorIs(t, err, expErr)
	assert.Equal(t, database.User{}, user)
}
