package getpostresponses_test

import (
	"context"
	"errors"
	"testing"

	getpostresponses "backend/internal/content/application/query/get_post_responses"
	database "backend/internal/database/sqlc"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type stubGetPostResponsesRepo struct {
	called bool
	got    pgtype.UUID

	fn func(ctx context.Context, postID pgtype.UUID) ([]database.GetPostResponsesRow, error)
}

func (s *stubGetPostResponsesRepo) GetPostResponses(ctx context.Context, postID pgtype.UUID) ([]database.GetPostResponsesRow, error) {
	s.called = true
	s.got = postID
	return s.fn(ctx, postID)
}

func TestGetPostResponses_InvalidPostID(t *testing.T) {
	repo := &stubGetPostResponsesRepo{fn: func(ctx context.Context, postID pgtype.UUID) ([]database.GetPostResponsesRow, error) {
		return nil, nil
	}}

	svc := getpostresponses.NewService(repo, logrus.New())
	rows, err := svc.GetPostResponses(context.Background(), "bad")

	assert.Nil(t, rows)
	assert.Error(t, err)
	assert.ErrorIs(t, err, getpostresponses.ErrInvalidPostID)
	assert.False(t, repo.called)
}

func TestGetPostResponses_Success(t *testing.T) {
	postUUID := uuid.New()
	repo := &stubGetPostResponsesRepo{fn: func(ctx context.Context, postID pgtype.UUID) ([]database.GetPostResponsesRow, error) {
		return []database.GetPostResponsesRow{{Message: "m"}}, nil
	}}

	svc := getpostresponses.NewService(repo, logrus.New())
	rows, err := svc.GetPostResponses(context.Background(), postUUID.String())

	require.NoError(t, err)
	assert.True(t, repo.called)
	assert.Equal(t, [16]byte(postUUID), repo.got.Bytes)
	assert.Len(t, rows, 1)
}

func TestGetPostResponses_RepositoryError(t *testing.T) {
	expErr := errors.New("db error")
	postUUID := uuid.New()
	repo := &stubGetPostResponsesRepo{fn: func(ctx context.Context, postID pgtype.UUID) ([]database.GetPostResponsesRow, error) {
		return nil, expErr
	}}

	svc := getpostresponses.NewService(repo, logrus.New())
	rows, err := svc.GetPostResponses(context.Background(), postUUID.String())

	assert.Nil(t, rows)
	assert.ErrorIs(t, err, expErr)
}
