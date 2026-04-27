package revokesession_test

import (
	"context"
	"errors"
	"testing"

	database "backend/internal/database/sqlc"
	revokesession "backend/internal/identity/application/command/revoke_session"
	sharedErrors "backend/internal/shared/errors"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

type stubRevokeRepo struct {
	called bool
	gotTok string

	fn func(ctx context.Context, refreshToken string) (database.Session, error)
}

func (s *stubRevokeRepo) RevokeSessionByRefreshToken(ctx context.Context, refreshToken string) (database.Session, error) {
	s.called = true
	s.gotTok = refreshToken
	return s.fn(ctx, refreshToken)
}

func TestRevokeSession_EmptyToken(t *testing.T) {
	repo := &stubRevokeRepo{}
	svc := revokesession.NewService(repo, logrus.New())

	err := svc.RevokeSession(context.Background(), "")

	assert.ErrorIs(t, err, sharedErrors.ErrEmptyRefreshToken)
	assert.False(t, repo.called)
}

func TestRevokeSession_RepositoryError(t *testing.T) {
	expErr := errors.New("db error")

	repo := &stubRevokeRepo{fn: func(ctx context.Context, refreshToken string) (database.Session, error) {
		return database.Session{}, expErr
	}}

	svc := revokesession.NewService(repo, logrus.New())
	err := svc.RevokeSession(context.Background(), "token")

	assert.ErrorIs(t, err, expErr)
	assert.True(t, repo.called)
	assert.Equal(t, "token", repo.gotTok)
}

func TestRevokeSession_Success(t *testing.T) {
	repo := &stubRevokeRepo{fn: func(ctx context.Context, refreshToken string) (database.Session, error) {
		return database.Session{}, nil
	}}

	svc := revokesession.NewService(repo, logrus.New())
	err := svc.RevokeSession(context.Background(), "token")

	assert.NoError(t, err)
	assert.True(t, repo.called)
}
