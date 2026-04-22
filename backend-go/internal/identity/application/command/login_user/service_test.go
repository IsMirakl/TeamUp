package loginuser_test

import (
	database "backend/internal/database/sqlc"
	loginuser "backend/internal/identity/application/command/login_user"
	"backend/internal/identity/application/command/login_user/mocks"
	"backend/internal/identity/application/dto"
	utilsmocks "backend/internal/pkg/utils/mocks"
	sharedErrors "backend/internal/shared/errors"
	"context"
	"errors"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"golang.org/x/crypto/bcrypt"
)

func makePasswordHash(t *testing.T, password string) string {
	t.Helper()

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		t.Fatalf("Failed to hash password: %v", err)
	}

	return string(hash)
}

func makeUUID() pgtype.UUID {
	id := uuid.New()

	return pgtype.UUID{
		Bytes: id,
		Valid: true,
	}
}

func TestLoginUserNotFound(t *testing.T) {
	repository := &mocks.MockRepository{}
	sessionService := &mocks.MockSessionService{}
	tokenService := &utilsmocks.MockTokenService{}

	log := logrus.New()

	service := loginuser.NewUserService(repository, sessionService, log, tokenService)

	req := &dto.LoginUserDTO{
		Email:    "test@gmail.com",
		Password: "password123",
	}

	repository.On("GetUserWithPasswordByEmail", mock.Anything, req.Email).
		Return(database.GetUserWithPasswordByEmailRow{}, pgx.ErrNoRows)

	resp, err := service.Login(context.Background(), req)

	assert.Nil(t, resp)
	assert.Error(t, err)
	assert.ErrorIs(t, err, sharedErrors.ErrInvalidCredentials)

	repository.AssertExpectations(t)
	sessionService.AssertNotCalled(t, "CreateSession", mock.Anything, mock.Anything)
	tokenService.AssertNotCalled(t, "GenerateAccessToken", mock.Anything)
	tokenService.AssertNotCalled(t, "GenerateRefreshToken", mock.Anything)
}

func TestLoginRepositoryError(t *testing.T) {
	repository := &mocks.MockRepository{}
	sessionService := &mocks.MockSessionService{}
	tokenService := &utilsmocks.MockTokenService{}
	log := logrus.New()

	service := loginuser.NewUserService(repository, sessionService, log, tokenService)

	req := &dto.LoginUserDTO{
		Email:    "test@gmail.com",
		Password: "password123",
	}

	expectedErr := errors.New("database error")

	repository.
		On("GetUserWithPasswordByEmail", mock.Anything, req.Email).
		Return(database.GetUserWithPasswordByEmailRow{}, expectedErr)

	resp, err := service.Login(context.Background(), req)

	assert.Nil(t, resp)
	assert.Error(t, err)
	assert.ErrorIs(t, err, expectedErr)

	repository.AssertExpectations(t)
	sessionService.AssertNotCalled(t, "CreateSession", mock.Anything, mock.Anything)
	tokenService.AssertNotCalled(t, "GenerateAccessToken", mock.Anything)
	tokenService.AssertNotCalled(t, "GenerateRefreshToken", mock.Anything)
}

func TestLoginInvalidPassword(t *testing.T) {
	repository := &mocks.MockRepository{}
	sessionService := &mocks.MockSessionService{}
	tokenService := &utilsmocks.MockTokenService{}
	log := logrus.New()

	service := loginuser.NewUserService(repository, sessionService, log, tokenService)

	req := &dto.LoginUserDTO{
		Email:    "test@gmail.com",
		Password: "password123",
	}

	user := database.GetUserWithPasswordByEmailRow{
		UserID:       makeUUID(),
		Email:        "test@gmail.com",
		PasswordHash: makePasswordHash(t, "correct-password"),
	}

	repository.On("GetUserWithPasswordByEmail", mock.Anything, req.Email).Return(user, nil)

	resp, err := service.Login(context.Background(), req)

	assert.Nil(t, resp)
	assert.Error(t, err)
	assert.ErrorIs(t, err, sharedErrors.ErrInvalidCredentials)

	repository.AssertExpectations(t)
	sessionService.AssertNotCalled(t, "CreateSession", mock.Anything, mock.Anything)
	tokenService.AssertNotCalled(t, "GenerateAccessToken", mock.Anything)
	tokenService.AssertNotCalled(t, "GenerateRefreshToken", mock.Anything)
}

func TestLoginSuccess(t *testing.T) {
	repository := &mocks.MockRepository{}
	sessionService := &mocks.MockSessionService{}
	tokenService := &utilsmocks.MockTokenService{}
	log := logrus.New()

	service := loginuser.NewUserService(repository, sessionService, log, tokenService)

	req := &dto.LoginUserDTO{
		Email:    "test@gmail.com",
		Password: "password123",
	}

	user := database.GetUserWithPasswordByEmailRow{
		UserID:       makeUUID(),
		Email:        "test@gmail.com",
		PasswordHash: makePasswordHash(t, "password123"),
	}

	repository.
		On("GetUserWithPasswordByEmail", mock.Anything, req.Email).
		Return(user, nil)

	tokenService.
		On("GenerateAccessToken", user.UserID.String()).
		Return("access-token", nil)

	tokenService.
		On("GenerateRefreshToken", user.UserID.String()).
		Return("refresh-token", nil)

	sessionService.
		On("CreateSession", mock.Anything, mock.MatchedBy(func(r *dto.CreateSessionDTO) bool {
			return r.UserID == user.UserID.String() &&
				r.RefreshToken == "refresh-token" &&
				r.UserAgent == req.UserAgent &&
				r.ClientIp == req.ClientIP &&
				r.ID != "" &&
				r.ExpiresAt.After(time.Now()) &&
				r.IsBlocked == false
		})).
		Return(&dto.SessionResponse{
			ID: "session-id",
		}, nil)

	resp, err := service.Login(context.Background(), req)

	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, "session-id", resp.SessionId)
	assert.Equal(t, "access-token", resp.AccessToken)
	assert.Equal(t, "refresh-token", resp.RefreshToken)

	repository.AssertExpectations(t)
	tokenService.AssertExpectations(t)
	sessionService.AssertExpectations(t)
}
