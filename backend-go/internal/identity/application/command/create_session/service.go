package createsession

import (
	database "backend/internal/database/sqlc"
	"backend/internal/identity/application/dto"
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/sirupsen/logrus"
)

type Repository interface {
	CreateSession(ctx context.Context, sessionParams database.CreateSessionParams) (database.Session, error)
}

type Service struct {
	repository Repository
	log        *logrus.Logger
}

func NewSesssionService(repository Repository, log *logrus.Logger) *Service {
	return &Service{
		repository: repository,
		log:        log,
	}
}

func (s *Service) CreateSession(ctx context.Context, request *dto.CreateSessionDTO) (*dto.SessionResponse, error) {
	userID, err := uuid.Parse(request.UserID)
	if err != nil {
		s.log.WithField("user_id", request.UserID).WithError(err).Error("invalid user_id")
		return nil, err
	}

	sessionParams := database.CreateSessionParams{
		UserID:       pgtype.UUID{Bytes: userID, Valid: true},
		RefreshToken: request.RefreshToken,
		UserAgent:    request.UserAgent,
		ClientIp:     request.ClientIp,
		IsBlocked:    request.IsBlocked,
		ExpiresAt: pgtype.Timestamptz{
			Time:  request.ExpiresAt,
			Valid: true,
		},
	}

	session, err := s.repository.CreateSession(ctx, sessionParams)
	if err != nil {
		s.log.WithError(err).Error("failed to create session")
		return nil, err
	}

	return dto.ToSessionMapper(&session), nil
}
