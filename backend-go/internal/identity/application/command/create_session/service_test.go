package createsession_test

import (
	database "backend/internal/database/sqlc"
	createsession "backend/internal/identity/application/command/create_session"
	"backend/internal/identity/application/command/create_session/mocks"
	"backend/internal/identity/application/dto"
	"context"
	"errors"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func makeUUID() pgtype.UUID {
	id := uuid.New()

	return pgtype.UUID{
		Bytes: id,
		Valid: true,
	}
}

func TestCreateSession_Success(t *testing.T) {
	repository := &mocks.MockRepository{}

	log := logrus.New()

	service := createsession.NewSessionService(repository, log)

	userID := uuid.New()
	sessionID := makeUUID()
	expiresAt := time.Now()

	req := &dto.CreateSessionDTO{
		UserID:       userID.String(),
		RefreshToken: "refresh-token",
		UserAgent:    "Mozilla/Firefox",
		ClientIp:     "127.0.0.1",
		IsBlocked:    false,
		ExpiresAt:    expiresAt,
	}

	expParams := database.CreateSessionParams{
		UserID: pgtype.UUID{
			Bytes: userID,
			Valid: true,
		},
		RefreshToken: "refresh-token",
		UserAgent:    "Mozilla/Firefox",
		ClientIp:     "127.0.0.1",
		IsBlocked:    false,
		ExpiresAt: pgtype.Timestamptz{
			Time:  expiresAt,
			Valid: true,
		},
	}

	expSession := database.Session{
		ID:           sessionID,
		UserID:       expParams.UserID,
		RefreshToken: "refresh-token",
		UserAgent:    "Mozilla/Firefox",
		ClientIp:     "127.0.0.1",
		IsBlocked:    false,
		ExpiresAt:    expParams.ExpiresAt,
		CreatedAt: pgtype.Timestamptz{
			Time:  time.Now(),
			Valid: true,
		},
	}

	repository.On("CreateSession", context.Background(), expParams).Return(expSession, nil).Once()

	resp, err := service.CreateSession(context.Background(), req)

	require.NoError(t, err)
	require.NotNil(t, resp)

	assert.Equal(t, "refresh-token", resp.RefreshToken)

	repository.AssertExpectations(t)
}

func TestCreateSession_InvalidUserID(t *testing.T) {
	repository := &mocks.MockRepository{}

	log := logrus.New()

	service := createsession.NewSessionService(repository, log)

	req := &dto.CreateSessionDTO{
		UserID:       "bad-user-id",
		RefreshToken: "refresh-token",
		ExpiresAt:    time.Now(),
	}

	resp, err := service.CreateSession(context.Background(), req)

	require.Error(t, err)
	require.Nil(t, resp)

	repository.AssertNotCalled(t, "CreateSession")

}

func TestCreateSession_RepositoryError(t *testing.T) {
	repository := &mocks.MockRepository{}

	log := logrus.New()

	service := createsession.NewSessionService(repository, log)

	userID := uuid.New()
	expiresAt := time.Now()
	expectedErr := errors.New("database error")

	req := &dto.CreateSessionDTO{
		UserID:       userID.String(),
		RefreshToken: "refresh-token",
		UserAgent:    "Mozilla/Firefox",
		ClientIp:     "127.0.0.1",
		IsBlocked:    false,
		ExpiresAt:    expiresAt,
	}

	repository.
		On("CreateSession", mock.Anything, mock.AnythingOfType("database.CreateSessionParams")).
		Return(database.Session{}, expectedErr).
		Once()

	resp, err := service.CreateSession(context.Background(), req)

	require.Error(t, err)
	assert.Nil(t, resp)
	assert.ErrorIs(t, err, expectedErr)

	repository.AssertExpectations(t)
}
