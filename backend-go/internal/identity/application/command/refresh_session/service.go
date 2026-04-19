package refreshsession

import (
	database "backend/internal/database/sqlc"
	"backend/internal/identity/application/dto"
	"backend/internal/pkg/config"
	auth "backend/internal/pkg/utils"
	"backend/internal/shared/errors"
	"context"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/sirupsen/logrus"
)

type Repository interface {
	GetSessionByRefreshToken(ctx context.Context, refreshToken string) (database.Session, error)
	UpdateSessionRefreshToken(ctx context.Context, arg database.UpdateSessionRefreshTokenParams) (database.Session, error)
}

type Service struct {
	repository Repository
	log        *logrus.Logger
}

func NewSessionService(repository Repository, log *logrus.Logger) *Service {
	return &Service{
		repository: repository,
		log:        log,
	}
}

func (s *Service) RefreshSession(ctx context.Context, refreshToken string) (*dto.LoginResponse, error) {
	conf, err := config.New(s.log)
	if err != nil {
		s.log.WithError(err).Error("Failed to create config")
		return nil, errors.ErrUnauthorized
	}

	signingRefreshKey := []byte(conf.SECRET_KEY.REFRESH_SECRET)

	claims, err := auth.ValidateRefreshToken(refreshToken, signingRefreshKey)
	if err != nil {
		s.log.WithError(err).Warn("Invalid refresh token")
		return nil, errors.ErrUnauthorized
	}

	session, err := s.repository.GetSessionByRefreshToken(ctx, refreshToken)
	if err != nil {
		s.log.WithError(err).Error("Failed to refresh session")
		return nil, errors.ErrUnauthorized
	}
	if session.IsBlocked {
		s.log.Warn("session is blocked")
		return nil, errors.ErrUnauthorized
	}

	if session.UserID.String() != claims.UserID {
		s.log.WithFields(logrus.Fields{
			"token_user_id":   claims.UserID,
			"session_user_id": session.UserID.String(),
		}).Warn("refresh token user mismatch")
		return nil, errors.ErrUnauthorized
	}

	if !session.ExpiresAt.Valid || session.ExpiresAt.Time.Before(time.Now()) {
		s.log.Warn("refresh token expired")
		return nil, errors.ErrUnauthorized
	}

	accessToken, err := auth.CreateToken(claims.UserID, s.log)
	if err != nil {
		s.log.WithError(err).Error("failed to create access token")
		return nil, err
	}

	newRefreshToken, err := auth.GenerateRefreshToken(claims.UserID, s.log)
	if err != nil {
		s.log.WithError(err).Error("failed to create refresh token")
		return nil, err
	}

	newExpiresAt := time.Now().Add(90 * 24 * time.Hour)
	_, err = s.repository.UpdateSessionRefreshToken(ctx, database.UpdateSessionRefreshTokenParams{
		ID:           session.ID,
		RefreshToken: newRefreshToken,
		ExpiresAt: pgtype.Timestamptz{
			Time:  newExpiresAt,
			Valid: true,
		},
	})
	if err != nil {
		s.log.WithError(err).Error("failed to update session refresh token")
		return nil, errors.ErrUnauthorized
	}

	return &dto.LoginResponse{
		SessionId:    session.ID.String(),
		AccessToken:  accessToken,
		RefreshToken: newRefreshToken,
		IsBlocked:    false,
		ExpiresAt:    newExpiresAt,
	}, nil
}
