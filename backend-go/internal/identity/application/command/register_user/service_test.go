package registeruser_test

import (
	"context"
	"errors"
	"testing"
	"time"

	database "backend/internal/database/sqlc"
	registeruser "backend/internal/identity/application/command/register_user"
	"backend/internal/identity/application/dto"
	auth "backend/internal/pkg/utils"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type stubRegisterRepo struct {
	called           bool
	gotUserParams    database.CreateUserParams
	gotAccountParams database.CreateAccountParams

	createWithAccountFn func(ctx context.Context, userParams database.CreateUserParams, accountParams database.CreateAccountParams) (database.User, error)
}

func (s *stubRegisterRepo) CreateWithAccount(ctx context.Context, userParams database.CreateUserParams, accountParams database.CreateAccountParams) (database.User, error) {
	s.called = true
	s.gotUserParams = userParams
	s.gotAccountParams = accountParams
	return s.createWithAccountFn(ctx, userParams, accountParams)
}

type stubSessionService struct {
	called bool
	gotReq *dto.CreateSessionDTO

	createSessionFn func(ctx context.Context, request *dto.CreateSessionDTO) (*dto.SessionResponse, error)
}

func (s *stubSessionService) CreateSession(ctx context.Context, request *dto.CreateSessionDTO) (*dto.SessionResponse, error) {
	s.called = true
	s.gotReq = request
	return s.createSessionFn(ctx, request)
}

type stubTokenService struct {
	calledGenerateAccess  bool
	calledGenerateRefresh bool
	gotAccessUserID       string
	gotRefreshUserID      string

	generateAccessFn  func(userID string) (string, error)
	generateRefreshFn func(userID string) (string, error)
}

func (s *stubTokenService) GenerateAccessToken(userID string) (string, error) {
	s.calledGenerateAccess = true
	s.gotAccessUserID = userID
	return s.generateAccessFn(userID)
}

func (s *stubTokenService) GenerateRefreshToken(userID string) (string, error) {
	s.calledGenerateRefresh = true
	s.gotRefreshUserID = userID
	return s.generateRefreshFn(userID)
}

func (s *stubTokenService) ValidateAccessToken(tokenString string) (*auth.Claims, error) {
	return nil, errors.New("not implemented")
}

func (s *stubTokenService) ValidateRefreshToken(tokenString string) (*auth.Claims, error) {
	return nil, errors.New("not implemented")
}

func TestRegisterUser_Success_NoAvatar(t *testing.T) {
	log := logrus.New()

	repo := &stubRegisterRepo{}
	sessions := &stubSessionService{}
	tokens := &stubTokenService{}

	repo.createWithAccountFn = func(ctx context.Context, userParams database.CreateUserParams, accountParams database.CreateAccountParams) (database.User, error) {
		require.NotEmpty(t, userParams.Email)
		require.NotEmpty(t, userParams.Name)
		require.True(t, userParams.UserID.Valid)
		require.False(t, userParams.Avatar.Valid)
		require.Equal(t, database.RolesUser, userParams.Role)
		require.Equal(t, database.SubscriptionPlansFree, userParams.SubscriptionPlan)

		require.True(t, accountParams.UserID.Valid)
		require.Equal(t, userParams.UserID, accountParams.UserID)
		require.NotEmpty(t, accountParams.PasswordHash)
		require.Equal(t, database.ProvidersLocal, accountParams.Provider)

		return database.User{
			UserID:           userParams.UserID,
			Email:            userParams.Email,
			Name:             userParams.Name,
			Avatar:           userParams.Avatar,
			Role:             userParams.Role,
			SubscriptionPlan: userParams.SubscriptionPlan,
		}, nil
	}

	tokens.generateAccessFn = func(userID string) (string, error) { return "access-token", nil }
	tokens.generateRefreshFn = func(userID string) (string, error) { return "refresh-token", nil }

	sessions.createSessionFn = func(ctx context.Context, request *dto.CreateSessionDTO) (*dto.SessionResponse, error) {
		require.Equal(t, "refresh-token", request.RefreshToken)
		require.Equal(t, "ua", request.UserAgent)
		require.Equal(t, "ip", request.ClientIp)
		require.False(t, request.IsBlocked)
		require.True(t, request.ExpiresAt.After(time.Now()))
		return &dto.SessionResponse{ID: "session-id"}, nil
	}

	svc := registeruser.NewUserService(repo, sessions, log, tokens)

	req := &dto.CreateUserDTO{
		Email:     "test@example.com",
		Name:      "Test",
		Password:  "password123",
		UserAgent: "ua",
		ClientIP:  "ip",
	}

	resp, err := svc.Create(context.Background(), req)
	require.NoError(t, err)
	require.NotNil(t, resp)

	assert.True(t, repo.called)
	assert.True(t, sessions.called)
	assert.True(t, tokens.calledGenerateAccess)
	assert.True(t, tokens.calledGenerateRefresh)

	assert.Equal(t, "access-token", resp.AccessToken)
	assert.Equal(t, "refresh-token", resp.RefreshToken)
	require.NotNil(t, resp.User)
	assert.Equal(t, req.Email, resp.User.Email)
	assert.Equal(t, req.Name, resp.User.Name)
	assert.Nil(t, resp.User.Avatar)
}

func TestRegisterUser_RepositoryError(t *testing.T) {
	log := logrus.New()

	expErr := errors.New("db error")

	repo := &stubRegisterRepo{}
	sessions := &stubSessionService{}
	tokens := &stubTokenService{}

	repo.createWithAccountFn = func(ctx context.Context, userParams database.CreateUserParams, accountParams database.CreateAccountParams) (database.User, error) {
		return database.User{}, expErr
	}

	svc := registeruser.NewUserService(repo, sessions, log, tokens)
	resp, err := svc.Create(context.Background(), &dto.CreateUserDTO{Email: "test@example.com", Name: "Test", Password: "password123"})

	assert.Nil(t, resp)
	assert.ErrorIs(t, err, expErr)
	assert.False(t, sessions.called)
	assert.False(t, tokens.calledGenerateAccess)
	assert.False(t, tokens.calledGenerateRefresh)
}

func TestRegisterUser_CreateSessionError(t *testing.T) {
	log := logrus.New()

	expErr := errors.New("session error")

	repo := &stubRegisterRepo{}
	sessions := &stubSessionService{}
	tokens := &stubTokenService{}

	repo.createWithAccountFn = func(ctx context.Context, userParams database.CreateUserParams, accountParams database.CreateAccountParams) (database.User, error) {
		return database.User{UserID: userParams.UserID, Email: userParams.Email, Name: userParams.Name}, nil
	}
	tokens.generateAccessFn = func(userID string) (string, error) { return "access-token", nil }
	tokens.generateRefreshFn = func(userID string) (string, error) { return "refresh-token", nil }
	sessions.createSessionFn = func(ctx context.Context, request *dto.CreateSessionDTO) (*dto.SessionResponse, error) {
		return nil, expErr
	}

	svc := registeruser.NewUserService(repo, sessions, log, tokens)
	resp, err := svc.Create(context.Background(), &dto.CreateUserDTO{Email: "test@example.com", Name: "Test", Password: "password123"})

	assert.Nil(t, resp)
	assert.ErrorIs(t, err, expErr)
}
