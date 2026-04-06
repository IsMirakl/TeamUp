package registeruser

import (
	database "backend/internal/database/sqlc"
	"backend/internal/features/user/dto"
	"backend/internal/features/user/model"
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/sirupsen/logrus"
)

type Service struct {
	repository *userRepository
	log        *logrus.Logger
}

func NewUserService(repository *userRepository, log *logrus.Logger) *Service {
	return &Service{
		repository: repository,
		log:        log,
	}
}

func (s *Service) Create(ctx context.Context, dto *dto.CreateUserDTO) (*database.User, error) {
	tx, err := s.repository.pool.Begin(ctx)
	if err != nil {
		s.log.WithError(err).Error("failed to begin transaction")
		return nil, err
	}
	defer func() {
		_ = tx.Rollback(ctx)
	}()

	qtx := s.repository.q.WithTx(tx)

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

	user, err := qtx.CreateUser(ctx, database.CreateUserParams{
		UserID:           pgtype.UUID{Bytes: userID, Valid: true},
		Email:            dto.Email,
		Name:             dto.Name,
		Avatar:           avatar,
		Role:             "user",
		SubscriptionPlan: "Free",
	})
	if err != nil {
		s.log.WithFields(logrus.Fields{
			"user_id": userID.String(),
			"email":   dto.Email,
		}).WithError(err).Error("failed to create user")

		return nil, err
	}

	err = qtx.CreateAccount(ctx, database.CreateAccountParams{
		UserID:           pgtype.UUID{Bytes: userID, Valid: true},
		PasswordHash: hash,
		Provider:     "local",
	})
	if err != nil {
		s.log.WithField("user_id", userID.String()).
			WithError(err).
			Error("failed to create account")

		return nil, err
	}

	if err := tx.Commit(ctx); err != nil {
		s.log.WithField("user_id", userID.String()).
			WithError(err).
			Error("failed to commit transaction")
		return nil, err
	}

	s.log.WithFields(logrus.Fields{
		"user_id": userID.String(),
		"email":   dto.Email,
	}).Info("user successfully created")

	return &user, nil
}