package loginuser

import (
	"context"
	"time"

	database "backend/internal/database/sqlc"
	"backend/internal/identity/application/dto"
	"backend/internal/identity/domain/model"
	auth "backend/internal/pkg/utils"
	sharedErrors "backend/internal/shared/errors"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/sirupsen/logrus"
)

type Repository interface {
	GetUserWithPasswordByEmail(ctx context.Context, email string) (database.GetUserWithPasswordByEmailRow, error)
}

type SessionService interface {
	CreateSession(ctx context.Context, request *dto.CreateSessionDTO) (*dto.SessionResponse, error)
}

type Service struct {
	repository     Repository
	log            *logrus.Logger
	sessionService SessionService
	tokenService   auth.TokenService
}

func NewUserService(
	repository Repository, sessionService SessionService, log *logrus.Logger, tokenService auth.TokenService) *Service {
	return &Service{
		repository:     repository,
		log:            log,
		sessionService: sessionService,
		tokenService:   tokenService,
	}
}

func (s *Service) Login(ctx context.Context, request *dto.LoginUserDTO) (*dto.LoginResponse, error) {
	s.log.WithField("email", request.Email).Info("login request received")

	user, err := s.repository.GetUserWithPasswordByEmail(ctx, request.Email)
	if err != nil {
		if err == pgx.ErrNoRows {
			s.log.Warn("user not found")
			return nil, sharedErrors.ErrInvalidCredentials
		}

		s.log.WithError(err).Error("failed to fetch user by email")
		return nil, err
	}

	if !model.VerifyPassword(user.PasswordHash, request.Password) {
		s.log.Warn("invalid password")
		return nil, sharedErrors.ErrInvalidCredentials
	}

	accessToken, err := s.tokenService.GenerateAccessToken(user.UserID.String())
	if err != nil {
		s.log.WithError(err).Error("failed to create access token")
		return nil, err
	}

	refreshToken, err := s.tokenService.GenerateRefreshToken(user.UserID.String())
	if err != nil {
		s.log.WithError(err).Error("failed to create refresh token")
		return nil, err
	}

	session, err := s.sessionService.CreateSession(ctx, &dto.CreateSessionDTO{
		ID:           uuid.New().String(),
		UserID:       user.UserID.String(),
		RefreshToken: refreshToken,
		UserAgent:    request.UserAgent,
		ClientIp:     request.ClientIP,
		IsBlocked:    false,
		ExpiresAt:    time.Now().Add(90 * 24 * time.Hour),
	})
	if err != nil {
		s.log.WithError(err).Error("failed to create session")
		return nil, err
	}

	s.log.WithField("email", user.Email).Info("login successful")

	return &dto.LoginResponse{
		SessionId:    session.ID,
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}
