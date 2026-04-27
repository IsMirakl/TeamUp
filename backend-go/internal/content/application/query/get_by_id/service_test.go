package getbyid_test

import (
	"context"
	"errors"
	"testing"

	getbyid "backend/internal/content/application/query/get_by_id"
	database "backend/internal/database/sqlc"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type stubGetPostByIDRepo struct {
	called bool
	got    pgtype.UUID

	fn func(ctx context.Context, id pgtype.UUID) (database.GetPostByIdRow, error)
}

func (s *stubGetPostByIDRepo) GetPostById(ctx context.Context, id pgtype.UUID) (database.GetPostByIdRow, error) {
	s.called = true
	s.got = id
	return s.fn(ctx, id)
}

func TestGetPostById_InvalidUUID(t *testing.T) {
	repo := &stubGetPostByIDRepo{fn: func(ctx context.Context, id pgtype.UUID) (database.GetPostByIdRow, error) {
		return database.GetPostByIdRow{}, nil
	}}

	svc := getbyid.NewService(repo, logrus.New())
	post, err := svc.GetById(context.Background(), "bad")

	assert.Error(t, err)
	assert.False(t, repo.called)
	assert.Equal(t, &database.GetPostByIdRow{}, post)
}

func TestGetPostById_Success(t *testing.T) {
	postUUID := uuid.New()
	repo := &stubGetPostByIDRepo{fn: func(ctx context.Context, id pgtype.UUID) (database.GetPostByIdRow, error) {
		return database.GetPostByIdRow{ID: id, Title: "t"}, nil
	}}

	svc := getbyid.NewService(repo, logrus.New())
	post, err := svc.GetById(context.Background(), postUUID.String())

	require.NoError(t, err)
	require.NotNil(t, post)
	assert.True(t, repo.called)
	assert.Equal(t, [16]byte(postUUID), repo.got.Bytes)
	assert.Equal(t, "t", post.Title)
}

func TestGetPostById_RepositoryError(t *testing.T) {
	expErr := errors.New("db error")
	postUUID := uuid.New()

	repo := &stubGetPostByIDRepo{fn: func(ctx context.Context, id pgtype.UUID) (database.GetPostByIdRow, error) {
		return database.GetPostByIdRow{}, expErr
	}}

	svc := getbyid.NewService(repo, logrus.New())
	post, err := svc.GetById(context.Background(), postUUID.String())

	assert.ErrorIs(t, err, expErr)
	assert.Equal(t, &database.GetPostByIdRow{}, post)
}
