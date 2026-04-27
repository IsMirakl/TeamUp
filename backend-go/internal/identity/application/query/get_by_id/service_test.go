package getbyid_test

import (
	"context"
	"errors"
	"testing"

	database "backend/internal/database/sqlc"
	getbyid "backend/internal/identity/application/query/get_by_id"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type stubGetUserByIDRepo struct {
	called bool
	got    pgtype.UUID

	fn func(ctx context.Context, userID pgtype.UUID) (database.User, error)
}

func (s *stubGetUserByIDRepo) GetUserById(ctx context.Context, userID pgtype.UUID) (database.User, error) {
	s.called = true
	s.got = userID
	return s.fn(ctx, userID)
}

func TestGetById_InvalidUUID(t *testing.T) {
	repo := &stubGetUserByIDRepo{fn: func(ctx context.Context, userID pgtype.UUID) (database.User, error) {
		return database.User{}, nil
	}}

	svc := getbyid.NewService(repo, logrus.New())
	user, err := svc.GetById(context.Background(), "bad")

	assert.Error(t, err)
	assert.False(t, repo.called)
	assert.Equal(t, &database.User{}, user)
}

func TestGetById_Success(t *testing.T) {
	userUUID := uuid.New()
	repo := &stubGetUserByIDRepo{fn: func(ctx context.Context, userID pgtype.UUID) (database.User, error) {
		return database.User{UserID: userID, Email: "test@example.com"}, nil
	}}

	svc := getbyid.NewService(repo, logrus.New())
	user, err := svc.GetById(context.Background(), userUUID.String())

	require.NoError(t, err)
	require.NotNil(t, user)
	assert.True(t, repo.called)
	assert.True(t, repo.got.Valid)
	assert.Equal(t, [16]byte(userUUID), repo.got.Bytes)
	assert.Equal(t, "test@example.com", user.Email)
}

func TestGetById_RepositoryError(t *testing.T) {
	expErr := errors.New("db error")
	userUUID := uuid.New()

	repo := &stubGetUserByIDRepo{fn: func(ctx context.Context, userID pgtype.UUID) (database.User, error) {
		return database.User{}, expErr
	}}

	svc := getbyid.NewService(repo, logrus.New())
	user, err := svc.GetById(context.Background(), userUUID.String())

	assert.ErrorIs(t, err, expErr)
	assert.Equal(t, &database.User{}, user)
}
