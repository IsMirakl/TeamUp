package registeruser

import (
	"context"

	database "backend/internal/database/sqlc"
	"backend/internal/identity/application/dto"
	"backend/internal/identity/domain/model"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/sirupsen/logrus"
)

type Repository interface {
	CreateWithAccount(ctx context.Context, userParams database.CreateUserParams, accountParams database.CreateAccountParams) (database.User, error)
}

type Service struct {
	repository Repository
	log        *logrus.Logger
}

func NewUserService(repository Repository, log *logrus.Logger) *Service {
	return &Service{
		repository: repository,
		log:        log,
	}
}

func (s *Service) Create(ctx context.Context, dto *dto.CreateUserDTO) (*database.User, error) {
	s.log.WithField("email", dto.Email).Info("creating user")

	hash, err := model.HashPassword(dto.Password)
	if err != nil {
		s.log.WithError(err).Error("failed to hash password")
		return nil, err
	}

	userID := uuid.New()

	var avatar pgtype.Text
	if dto.Avatar != nil {
		avatar = pgtype.Text{
			String: *dto.Avatar,
			Valid:  true,
		}
	} else {
		avatar = pgtype.Text{
			Valid: false,
		}
	}

	userParams := database.CreateUserParams{
		UserID:           pgtype.UUID{Bytes: userID, Valid: true},
		Email:            dto.Email,
		Name:             dto.Name,
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
			"email":   dto.Email,
		}).WithError(err).Error("failed to create user")

		return nil, err
	}

	s.log.WithFields(logrus.Fields{
		"user_id": userID.String(),
		"email":   dto.Email,
	}).Info("user successfully created")

	return &user, nil
}
