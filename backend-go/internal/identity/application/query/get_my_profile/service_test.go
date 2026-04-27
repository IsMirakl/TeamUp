package getmyprofile_test

import (
	"context"
	"errors"
	"testing"

	database "backend/internal/database/sqlc"
	getmyprofile "backend/internal/identity/application/query/get_my_profile"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type stubGetMeRepo struct {
	called bool
	got    pgtype.UUID

	fn func(ctx context.Context, userID pgtype.UUID) (database.User, error)
}

func (s *stubGetMeRepo) GetMe(ctx context.Context, userID pgtype.UUID) (database.User, error) {
	s.called = true
	s.got = userID
	return s.fn(ctx, userID)
}

func TestGetMe_InvalidUUID(t *testing.T) {
	repo := &stubGetMeRepo{fn: func(ctx context.Context, userID pgtype.UUID) (database.User, error) {
		return database.User{}, nil
	}}

	svc := getmyprofile.NewPostService(repo, logrus.New())
	user, err := svc.GetMe(context.Background(), "bad")

	assert.Error(t, err)
	assert.False(t, repo.called)
	assert.Equal(t, &database.User{}, user)
}

func TestGetMe_Success(t *testing.T) {
	userUUID := uuid.New()
	repo := &stubGetMeRepo{fn: func(ctx context.Context, userID pgtype.UUID) (database.User, error) {
		return database.User{UserID: userID, Email: "me@example.com"}, nil
	}}

	svc := getmyprofile.NewPostService(repo, logrus.New())
	user, err := svc.GetMe(context.Background(), userUUID.String())

	require.NoError(t, err)
	require.NotNil(t, user)
	assert.True(t, repo.called)
	assert.True(t, repo.got.Valid)
	assert.Equal(t, [16]byte(userUUID), repo.got.Bytes)
	assert.Equal(t, "me@example.com", user.Email)
}

func TestGetMe_RepositoryError(t *testing.T) {
	expErr := errors.New("db error")
	userUUID := uuid.New()
	repo := &stubGetMeRepo{fn: func(ctx context.Context, userID pgtype.UUID) (database.User, error) {
		return database.User{}, expErr
	}}

	svc := getmyprofile.NewPostService(repo, logrus.New())
	user, err := svc.GetMe(context.Background(), userUUID.String())

	assert.ErrorIs(t, err, expErr)
	assert.Equal(t, &database.User{}, user)
}
