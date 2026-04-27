package createpost_test

import (
	"context"
	"errors"
	"testing"

	createpost "backend/internal/content/application/command/create_post"
	"backend/internal/content/application/dto"
	database "backend/internal/database/sqlc"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type stubCreatePostRepo struct {
	called bool
	got    database.CreatePostParams

	fn func(ctx context.Context, arg database.CreatePostParams) (database.Post, error)
}

func (s *stubCreatePostRepo) Create(ctx context.Context, arg database.CreatePostParams) (database.Post, error) {
	s.called = true
	s.got = arg
	return s.fn(ctx, arg)
}

func TestCreatePost_InvalidUserID(t *testing.T) {
	repo := &stubCreatePostRepo{fn: func(ctx context.Context, arg database.CreatePostParams) (database.Post, error) {
		return database.Post{}, nil
	}}

	svc := createpost.NewService(repo, logrus.New())
	post, err := svc.Create(context.Background(), &dto.CreatePostDTO{Title: "t", Description: "d", Tags: []string{"a"}}, "bad")

	assert.Error(t, err)
	assert.Nil(t, post)
	assert.False(t, repo.called)
}

func TestCreatePost_Success(t *testing.T) {
	authorUUID := uuid.New()

	repo := &stubCreatePostRepo{fn: func(ctx context.Context, arg database.CreatePostParams) (database.Post, error) {
		return database.Post{ID: arg.ID, Title: arg.Title, Description: arg.Description, Tags: arg.Tags, AuthorID: arg.AuthorID}, nil
	}}

	svc := createpost.NewService(repo, logrus.New())
	req := &dto.CreatePostDTO{Title: "title", Description: "desc", Tags: []string{"go", "api"}}

	post, err := svc.Create(context.Background(), req, authorUUID.String())

	require.NoError(t, err)
	require.NotNil(t, post)
	assert.True(t, repo.called)
	assert.True(t, repo.got.ID.Valid)
	assert.NotEqual(t, uuid.Nil, repo.got.ID.Bytes)
	assert.Equal(t, req.Title, repo.got.Title)
	assert.Equal(t, req.Description, repo.got.Description)
	assert.Equal(t, req.Tags, repo.got.Tags)
	assert.Equal(t, pgtype.UUID{Bytes: authorUUID, Valid: true}, repo.got.AuthorID)
}

func TestCreatePost_RepositoryError(t *testing.T) {
	expErr := errors.New("db error")
	authorUUID := uuid.New()

	repo := &stubCreatePostRepo{fn: func(ctx context.Context, arg database.CreatePostParams) (database.Post, error) {
		return database.Post{}, expErr
	}}

	svc := createpost.NewService(repo, logrus.New())
	post, err := svc.Create(context.Background(), &dto.CreatePostDTO{Title: "t", Description: "d"}, authorUUID.String())

	assert.Nil(t, post)
	assert.ErrorIs(t, err, expErr)
}
