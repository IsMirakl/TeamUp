package listposts_test

import (
	"context"
	"errors"
	"testing"

	listposts "backend/internal/content/application/query/list_posts"
	database "backend/internal/database/sqlc"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type stubListPostsRepo struct {
	called bool
	gotL   int32
	gotO   int32

	fn func(ctx context.Context, limit, offset int32) ([]database.ListPostsRow, error)
}

func (s *stubListPostsRepo) ListPosts(ctx context.Context, limit, offset int32) ([]database.ListPostsRow, error) {
	s.called = true
	s.gotL = limit
	s.gotO = offset
	return s.fn(ctx, limit, offset)
}

func TestListPosts_DefaultLimit(t *testing.T) {
	repo := &stubListPostsRepo{fn: func(ctx context.Context, limit, offset int32) ([]database.ListPostsRow, error) {
		return []database.ListPostsRow{}, nil
	}}

	svc := listposts.NewService(repo, logrus.New())
	rows, err := svc.ListPosts(context.Background(), 0, 0)

	require.NoError(t, err)
	assert.True(t, repo.called)
	assert.Equal(t, int32(20), repo.gotL)
	assert.Equal(t, int32(0), repo.gotO)
	assert.NotNil(t, rows)
}

func TestListPosts_ClampLimit(t *testing.T) {
	repo := &stubListPostsRepo{fn: func(ctx context.Context, limit, offset int32) ([]database.ListPostsRow, error) {
		return []database.ListPostsRow{}, nil
	}}

	svc := listposts.NewService(repo, logrus.New())
	_, err := svc.ListPosts(context.Background(), 500, 0)

	require.NoError(t, err)
	assert.Equal(t, int32(100), repo.gotL)
}

func TestListPosts_InvalidLimit(t *testing.T) {
	repo := &stubListPostsRepo{fn: func(ctx context.Context, limit, offset int32) ([]database.ListPostsRow, error) {
		return nil, nil
	}}

	svc := listposts.NewService(repo, logrus.New())
	rows, err := svc.ListPosts(context.Background(), -1, 0)

	assert.Nil(t, rows)
	assert.ErrorIs(t, err, listposts.ErrInvalidLimit)
	assert.False(t, repo.called)
}

func TestListPosts_InvalidOffset(t *testing.T) {
	repo := &stubListPostsRepo{fn: func(ctx context.Context, limit, offset int32) ([]database.ListPostsRow, error) {
		return nil, nil
	}}

	svc := listposts.NewService(repo, logrus.New())
	rows, err := svc.ListPosts(context.Background(), 10, -1)

	assert.Nil(t, rows)
	assert.ErrorIs(t, err, listposts.ErrInvalidOffset)
	assert.False(t, repo.called)
}

func TestListPosts_RepositoryError(t *testing.T) {
	expErr := errors.New("db error")
	repo := &stubListPostsRepo{fn: func(ctx context.Context, limit, offset int32) ([]database.ListPostsRow, error) {
		return nil, expErr
	}}

	svc := listposts.NewService(repo, logrus.New())
	rows, err := svc.ListPosts(context.Background(), 10, 0)

	assert.Nil(t, rows)
	assert.ErrorIs(t, err, expErr)
}
