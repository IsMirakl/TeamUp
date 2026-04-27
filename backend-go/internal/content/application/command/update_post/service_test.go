package updatepost_test

import (
	"context"
	"errors"
	"testing"

	updatepost "backend/internal/content/application/command/update_post"
	"backend/internal/content/application/dto"
	database "backend/internal/database/sqlc"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type stubUpdatePostRepo struct {
	called bool
	got    database.UpdatePostParams

	fn func(ctx context.Context, arg database.UpdatePostParams) (database.UpdatePostRow, error)
}

func (s *stubUpdatePostRepo) Update(ctx context.Context, arg database.UpdatePostParams) (database.UpdatePostRow, error) {
	s.called = true
	s.got = arg
	return s.fn(ctx, arg)
}

func TestUpdatePost_InvalidPostID(t *testing.T) {
	repo := &stubUpdatePostRepo{fn: func(ctx context.Context, arg database.UpdatePostParams) (database.UpdatePostRow, error) {
		return database.UpdatePostRow{}, nil
	}}

	svc := updatepost.NewService(repo, logrus.New())
	post, err := svc.Update(context.Background(), "bad", &dto.UpdatePostDTO{Title: "t", Description: "d"})

	assert.Error(t, err)
	assert.Nil(t, post)
	assert.False(t, repo.called)
}

func TestUpdatePost_Success(t *testing.T) {
	postUUID := uuid.New()

	repo := &stubUpdatePostRepo{fn: func(ctx context.Context, arg database.UpdatePostParams) (database.UpdatePostRow, error) {
		return database.UpdatePostRow{ID: arg.ID, Title: arg.Title, Description: arg.Description, Tags: arg.Tags}, nil
	}}

	svc := updatepost.NewService(repo, logrus.New())
	req := &dto.UpdatePostDTO{Title: "title", Description: "desc", Tags: []string{"x"}}

	post, err := svc.Update(context.Background(), postUUID.String(), req)

	require.NoError(t, err)
	require.NotNil(t, post)
	assert.True(t, repo.called)
	assert.Equal(t, pgtype.UUID{Bytes: postUUID, Valid: true}, repo.got.ID)
	assert.Equal(t, req.Title, repo.got.Title)
	assert.Equal(t, req.Description, repo.got.Description)
	assert.Equal(t, req.Tags, repo.got.Tags)
}

func TestUpdatePost_RepositoryError(t *testing.T) {
	expErr := errors.New("db error")
	postUUID := uuid.New()

	repo := &stubUpdatePostRepo{fn: func(ctx context.Context, arg database.UpdatePostParams) (database.UpdatePostRow, error) {
		return database.UpdatePostRow{}, expErr
	}}

	svc := updatepost.NewService(repo, logrus.New())
	post, err := svc.Update(context.Background(), postUUID.String(), &dto.UpdatePostDTO{Title: "t", Description: "d"})

	assert.Nil(t, post)
	assert.ErrorIs(t, err, expErr)
}
