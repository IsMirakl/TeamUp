package registeruser

import (
	"context"
	"time"

	database "backend/internal/database/sqlc"
	"backend/internal/identity/application/dto"
	"backend/internal/identity/domain/model"
	auth "backend/internal/pkg/utils"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/sirupsen/logrus"
)

type Repository interface {
	CreateWithAccount(ctx context.Context, userParams database.CreateUserParams, accountParams database.CreateAccountParams) (database.User, error)
}

type SessionService interface {
	CreateSession(ctx context.Context, request *dto.CreateSessionDTO) (*dto.SessionResponse, error)
}

type Service struct {
	repository     Repository
	sessionService SessionService
	log            *logrus.Logger
	tokenService   auth.TokenService
}

func NewUserService(repository Repository, sessionService SessionService, log *logrus.Logger, tokenService auth.TokenService) *Service {
	return &Service{
		repository:     repository,
		sessionService: sessionService,
		log:            log,
		tokenService:   tokenService,
	}
}

func (s *Service) Create(ctx context.Context, request *dto.CreateUserDTO) (*dto.RegisterResponse, error) {
	s.log.WithField("email", request.Email).Info("creating user")

	hash, err := model.HashPassword(request.Password)
	if err != nil {
		s.log.WithError(err).Error("failed to hash password")
		return nil, err
	}

	userID := uuid.New()

	var avatar pgtype.Text
	if request.Avatar != nil {
		avatar = pgtype.Text{
			String: *request.Avatar,
			Valid:  true,
		}
	} else {
		avatar = pgtype.Text{Valid: false}
	}

	userParams := database.CreateUserParams{
		UserID:           pgtype.UUID{Bytes: userID, Valid: true},
		Email:            request.Email,
		Name:             request.Name,
		Avatar:           avatar,
		Role:             "user",
		SubscriptionPlan: "Free",
	}

	accountParams := database.CreateAccountParams{
		UserID:       pgtype.UUID{Bytes: userID, Valid: true},
		PasswordHash: hash,
		Provider:     "local",
	}

	user, err := s.repository.CreateWithAccount(ctx, userParams, accountParams)
	if err != nil {
		s.log.WithFields(logrus.Fields{
			"user_id": userID.String(),
			"email":   request.Email,
		}).WithError(err).Error("failed to create user")

		return nil, err
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

	_, err = s.sessionService.CreateSession(ctx, &dto.CreateSessionDTO{
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

	s.log.WithFields(logrus.Fields{
		"user_id": userID.String(),
		"email":   request.Email,
	}).Info("user successfully created")

	return &dto.RegisterResponse{
		User:         dto.ToUserResponse(&user),
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}
