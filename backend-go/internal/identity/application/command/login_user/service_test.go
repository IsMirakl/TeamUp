package loginuser_test

import (
	database "backend/internal/database/sqlc"
	loginuser "backend/internal/identity/application/command/login_user"
	"backend/internal/identity/application/command/login_user/mocks"
	"backend/internal/identity/application/dto"
	sharedErrors "backend/internal/shared/errors"
	"context"
	"errors"
	"testing"

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
	// фейк объекты репозитория и сервиса
	repository := &mocks.MockRepository{}
	sessionService := &mocks.MockSessionService{}
	log := logrus.New()

	service := loginuser.NewUserService(repository, sessionService, log)

	req := &dto.LoginUserDTO{
		Email:    "test@gmail.com",
		Password: "password123",
	}

	repository.On("GetUserWithPasswordByEmail", mock.Anything, req.Email).
		Return(database.GetUserWithPasswordByEmailRow{}, pgx.ErrNoRows)

	// act = вызов того, что тестируем
	resp, err := service.Login(context.Background(), req)

	assert.Nil(t, resp)
	assert.Error(t, err)
	assert.ErrorIs(t, err, sharedErrors.ErrInvalidCredentials)

	repository.AssertExpectations(t)
	sessionService.AssertNotCalled(t, "CreateSession", mock.Anything, mock.Anything)
}

func TestLoginRepositoryError(t *testing.T) {
	repository := &mocks.MockRepository{}
	sessionService := &mocks.MockSessionService{}
	log := logrus.New()

	service := loginuser.NewUserService(repository, sessionService, log)

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
	assert.ErrorIs(t, expectedErr, err)

	repository.AssertExpectations(t)
	sessionService.AssertNotCalled(t, "CreateSession", mock.Anything, mock.Anything)

}
