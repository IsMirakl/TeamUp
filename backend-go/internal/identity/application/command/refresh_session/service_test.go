package refreshsession_test

import (
	"context"
	"errors"
	"testing"
	"time"

	database "backend/internal/database/sqlc"
	refreshsession "backend/internal/identity/application/command/refresh_session"
	auth "backend/internal/pkg/utils"
	sharedErrors "backend/internal/shared/errors"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type stubRefreshRepo struct {
	calledGet    bool
	calledUpdate bool
	gotToken     string
	gotUpdateArg database.UpdateSessionRefreshTokenParams

	getFn    func(ctx context.Context, refreshToken string) (database.Session, error)
	updateFn func(ctx context.Context, arg database.UpdateSessionRefreshTokenParams) (database.Session, error)
}

func (s *stubRefreshRepo) GetSessionByRefreshToken(ctx context.Context, refreshToken string) (database.Session, error) {
	s.calledGet = true
	s.gotToken = refreshToken
	return s.getFn(ctx, refreshToken)
}

func (s *stubRefreshRepo) UpdateSessionRefreshToken(ctx context.Context, arg database.UpdateSessionRefreshTokenParams) (database.Session, error) {
	s.calledUpdate = true
	s.gotUpdateArg = arg
	return s.updateFn(ctx, arg)
}

type stubRefreshTokenService struct {
	validateRefreshFn func(tokenString string) (*auth.Claims, error)
	genAccessFn       func(userID string) (string, error)
	genRefreshFn      func(userID string) (string, error)
}

func (s *stubRefreshTokenService) GenerateAccessToken(userID string) (string, error) {
	return s.genAccessFn(userID)
}

func (s *stubRefreshTokenService) GenerateRefreshToken(userID string) (string, error) {
	return s.genRefreshFn(userID)
}

func (s *stubRefreshTokenService) ValidateAccessToken(tokenString string) (*auth.Claims, error) {
	return nil, errors.New("not implemented")
}

func (s *stubRefreshTokenService) ValidateRefreshToken(tokenString string) (*auth.Claims, error) {
	return s.validateRefreshFn(tokenString)
}

func TestRefreshSession_InvalidToken_Unauthorized(t *testing.T) {
	log := logrus.New()

	repo := &stubRefreshRepo{}
	tokens := &stubRefreshTokenService{}

	tokens.validateRefreshFn = func(tokenString string) (*auth.Claims, error) {
		return nil, errors.New("bad token")
	}

	svc := refreshsession.NewSessionService(repo, log, tokens)
	resp, err := svc.RefreshSession(context.Background(), "bad")

	assert.Nil(t, resp)
	assert.ErrorIs(t, err, sharedErrors.ErrUnauthorized)
	assert.False(t, repo.calledGet)
	assert.False(t, repo.calledUpdate)
}

func TestRefreshSession_SessionBlocked_Unauthorized(t *testing.T) {
	log := logrus.New()

	repo := &stubRefreshRepo{}
	tokens := &stubRefreshTokenService{}

	userID := uuid.New().String()
	tokens.validateRefreshFn = func(tokenString string) (*auth.Claims, error) { return &auth.Claims{UserID: userID}, nil }
	repo.getFn = func(ctx context.Context, refreshToken string) (database.Session, error) {
		return database.Session{IsBlocked: true}, nil
	}

	svc := refreshsession.NewSessionService(repo, log, tokens)
	resp, err := svc.RefreshSession(context.Background(), "refresh")

	assert.Nil(t, resp)
	assert.ErrorIs(t, err, sharedErrors.ErrUnauthorized)
}

func TestRefreshSession_UserMismatch_Unauthorized(t *testing.T) {
	log := logrus.New()

	repo := &stubRefreshRepo{}
	tokens := &stubRefreshTokenService{}

	tokens.validateRefreshFn = func(tokenString string) (*auth.Claims, error) { return &auth.Claims{UserID: uuid.New().String()}, nil }

	otherID := uuid.New()
	repo.getFn = func(ctx context.Context, refreshToken string) (database.Session, error) {
		return database.Session{
			UserID:    pgtype.UUID{Bytes: otherID, Valid: true},
			IsBlocked: false,
			ExpiresAt: pgtype.Timestamptz{Time: time.Now().Add(1 * time.Hour), Valid: true},
		}, nil
	}

	svc := refreshsession.NewSessionService(repo, log, tokens)
	resp, err := svc.RefreshSession(context.Background(), "refresh")

	assert.Nil(t, resp)
	assert.ErrorIs(t, err, sharedErrors.ErrUnauthorized)
}

func TestRefreshSession_Expired_Unauthorized(t *testing.T) {
	log := logrus.New()

	repo := &stubRefreshRepo{}
	tokens := &stubRefreshTokenService{}

	userUUID := uuid.New()
	tokens.validateRefreshFn = func(tokenString string) (*auth.Claims, error) { return &auth.Claims{UserID: userUUID.String()}, nil }

	repo.getFn = func(ctx context.Context, refreshToken string) (database.Session, error) {
		return database.Session{
			UserID:    pgtype.UUID{Bytes: userUUID, Valid: true},
			IsBlocked: false,
			ExpiresAt: pgtype.Timestamptz{Time: time.Now().Add(-1 * time.Minute), Valid: true},
		}, nil
	}

	svc := refreshsession.NewSessionService(repo, log, tokens)
	resp, err := svc.RefreshSession(context.Background(), "refresh")

	assert.Nil(t, resp)
	assert.ErrorIs(t, err, sharedErrors.ErrUnauthorized)
}

func TestRefreshSession_Success(t *testing.T) {
	log := logrus.New()

	repo := &stubRefreshRepo{}
	tokens := &stubRefreshTokenService{}

	userUUID := uuid.New()
	sessionUUID := uuid.New()

	tokens.validateRefreshFn = func(tokenString string) (*auth.Claims, error) { return &auth.Claims{UserID: userUUID.String()}, nil }
	tokens.genAccessFn = func(userID string) (string, error) {
		require.Equal(t, userUUID.String(), userID)
		return "new-access", nil
	}
	tokens.genRefreshFn = func(userID string) (string, error) {
		require.Equal(t, userUUID.String(), userID)
		return "new-refresh", nil
	}

	repo.getFn = func(ctx context.Context, refreshToken string) (database.Session, error) {
		return database.Session{
			ID:           pgtype.UUID{Bytes: sessionUUID, Valid: true},
			UserID:       pgtype.UUID{Bytes: userUUID, Valid: true},
			RefreshToken: refreshToken,
			IsBlocked:    false,
			ExpiresAt:    pgtype.Timestamptz{Time: time.Now().Add(1 * time.Hour), Valid: true},
		}, nil
	}

	repo.updateFn = func(ctx context.Context, arg database.UpdateSessionRefreshTokenParams) (database.Session, error) {
		require.Equal(t, pgtype.UUID{Bytes: sessionUUID, Valid: true}, arg.ID)
		require.Equal(t, "new-refresh", arg.RefreshToken)
		require.True(t, arg.ExpiresAt.Valid)
		require.True(t, arg.ExpiresAt.Time.After(time.Now()))
		return database.Session{ID: pgtype.UUID{Bytes: sessionUUID, Valid: true}}, nil
	}

	svc := refreshsession.NewSessionService(repo, log, tokens)
	resp, err := svc.RefreshSession(context.Background(), "old-refresh")

	require.NoError(t, err)
	require.NotNil(t, resp)
	assert.Equal(t, sessionUUID.String(), resp.SessionId)
	assert.Equal(t, "new-access", resp.AccessToken)
	assert.Equal(t, "new-refresh", resp.RefreshToken)
	assert.False(t, resp.IsBlocked)
	assert.True(t, resp.ExpiresAt.After(time.Now()))
}

func TestRefreshSession_UpdateFails_Unauthorized(t *testing.T) {
	log := logrus.New()

	repo := &stubRefreshRepo{}
	tokens := &stubRefreshTokenService{}

	userUUID := uuid.New()

	tokens.validateRefreshFn = func(tokenString string) (*auth.Claims, error) { return &auth.Claims{UserID: userUUID.String()}, nil }
	tokens.genAccessFn = func(userID string) (string, error) { return "new-access", nil }
	tokens.genRefreshFn = func(userID string) (string, error) { return "new-refresh", nil }

	repo.getFn = func(ctx context.Context, refreshToken string) (database.Session, error) {
		sessionUUID := uuid.New()
		return database.Session{
			ID:           pgtype.UUID{Bytes: sessionUUID, Valid: true},
			UserID:       pgtype.UUID{Bytes: userUUID, Valid: true},
			RefreshToken: refreshToken,
			IsBlocked:    false,
			ExpiresAt:    pgtype.Timestamptz{Time: time.Now().Add(1 * time.Hour), Valid: true},
		}, nil
	}

	repo.updateFn = func(ctx context.Context, arg database.UpdateSessionRefreshTokenParams) (database.Session, error) {
		return database.Session{}, errors.New("db error")
	}

	svc := refreshsession.NewSessionService(repo, log, tokens)
	resp, err := svc.RefreshSession(context.Background(), "old-refresh")

	assert.Nil(t, resp)
	assert.ErrorIs(t, err, sharedErrors.ErrUnauthorized)
}
