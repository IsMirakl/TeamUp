package getauthorpost_test

import (
	"context"
	"errors"
	"testing"

	getauthorpost "backend/internal/content/application/query/get_author_post"
	database "backend/internal/database/sqlc"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type stubGetAuthorPostRepo struct {
	called bool
	got    string

	fn func(ctx context.Context, authorId string) (database.Post, error)
}

func (s *stubGetAuthorPostRepo) GetAuthorPost(ctx context.Context, authorId string) (database.Post, error) {
	s.called = true
	s.got = authorId
	return s.fn(ctx, authorId)
}

func TestGetAuthorPost_Success(t *testing.T) {
	repo := &stubGetAuthorPostRepo{fn: func(ctx context.Context, authorId string) (database.Post, error) {
		return database.Post{Title: "t"}, nil
	}}

	svc := getauthorpost.NewService(repo, logrus.New())
	post, err := svc.GetAuthorPost(context.Background(), "author")

	require.NoError(t, err)
	assert.True(t, repo.called)
	assert.Equal(t, "author", repo.got)
	assert.Equal(t, "t", post.Title)
}

func TestGetAuthorPost_RepositoryError(t *testing.T) {
	expErr := errors.New("db error")
	repo := &stubGetAuthorPostRepo{fn: func(ctx context.Context, authorId string) (database.Post, error) {
		return database.Post{}, expErr
	}}

	svc := getauthorpost.NewService(repo, logrus.New())
	post, err := svc.GetAuthorPost(context.Background(), "author")

	assert.ErrorIs(t, err, expErr)
	assert.Equal(t, database.Post{}, post)
}
