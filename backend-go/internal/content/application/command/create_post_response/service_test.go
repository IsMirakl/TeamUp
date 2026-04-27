package createpostresponse_test

import (
	"context"
	"errors"
	"testing"

	createpostresponse "backend/internal/content/application/command/create_post_response"
	"backend/internal/content/application/dto"
	database "backend/internal/database/sqlc"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type stubCreatePostResponseRepo struct {
	called bool
	got    database.CreatePostResponseParams

	fn func(ctx context.Context, arg database.CreatePostResponseParams) (database.CreatePostResponseRow, error)
}

func (s *stubCreatePostResponseRepo) Create(ctx context.Context, arg database.CreatePostResponseParams) (database.CreatePostResponseRow, error) {
	s.called = true
	s.got = arg
	return s.fn(ctx, arg)
}

func TestCreatePostResponse_InvalidUserID(t *testing.T) {
	repo := &stubCreatePostResponseRepo{fn: func(ctx context.Context, arg database.CreatePostResponseParams) (database.CreatePostResponseRow, error) {
		return database.CreatePostResponseRow{}, nil
	}}

	svc := createpostresponse.NewService(repo, logrus.New())
	resp, err := svc.Create(context.Background(), "bad", &dto.CreatePostResponseDTO{PostID: uuid.New().String(), Message: "m", Telegram: "t"})

	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.False(t, repo.called)
}

func TestCreatePostResponse_InvalidPostID(t *testing.T) {
	repo := &stubCreatePostResponseRepo{fn: func(ctx context.Context, arg database.CreatePostResponseParams) (database.CreatePostResponseRow, error) {
		return database.CreatePostResponseRow{}, nil
	}}

	svc := createpostresponse.NewService(repo, logrus.New())
	resp, err := svc.Create(context.Background(), uuid.New().String(), &dto.CreatePostResponseDTO{PostID: "bad", Message: "m", Telegram: "t"})

	assert.Error(t, err)
	assert.Nil(t, resp)
	assert.False(t, repo.called)
}

func TestCreatePostResponse_Success_TelegramOptional(t *testing.T) {
	userUUID := uuid.New()
	postUUID := uuid.New()

	repo := &stubCreatePostResponseRepo{fn: func(ctx context.Context, arg database.CreatePostResponseParams) (database.CreatePostResponseRow, error) {
		return database.CreatePostResponseRow{ResponseID: pgtype.UUID{Bytes: uuid.New(), Valid: true}}, nil
	}}

	svc := createpostresponse.NewService(repo, logrus.New())
	resp, err := svc.Create(context.Background(), userUUID.String(), &dto.CreatePostResponseDTO{PostID: postUUID.String(), Message: "hello", Telegram: ""})

	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.True(t, repo.called)
	assert.Equal(t, pgtype.UUID{Bytes: userUUID, Valid: true}, repo.got.UserID)
	assert.Equal(t, pgtype.UUID{Bytes: postUUID, Valid: true}, repo.got.PostID)
	assert.Equal(t, "hello", repo.got.Message)
	assert.False(t, repo.got.Telegram.Valid)
}

func TestCreatePostResponse_RepositoryError(t *testing.T) {
	expErr := errors.New("db error")
	userUUID := uuid.New()
	postUUID := uuid.New()

	repo := &stubCreatePostResponseRepo{fn: func(ctx context.Context, arg database.CreatePostResponseParams) (database.CreatePostResponseRow, error) {
		return database.CreatePostResponseRow{}, expErr
	}}

	svc := createpostresponse.NewService(repo, logrus.New())
	resp, err := svc.Create(context.Background(), userUUID.String(), &dto.CreatePostResponseDTO{PostID: postUUID.String(), Message: "m", Telegram: "t"})

	assert.Nil(t, resp)
	assert.ErrorIs(t, err, expErr)
}
