package loginuser

import (
	"backend/internal/features/user/dto"
	"backend/internal/features/user/model"
	auth "backend/internal/pkg/utils"
	"context"
	"fmt"
)

type UserFetcher interface {
	GetByEmail(ctx context.Context, email string) (*model.User, error)
}

type Service struct {
	userFetcher UserFetcher
}

func NewUserService(userFetcher UserFetcher) *Service {
	return &Service{userFetcher: userFetcher}
}

func (s *Service) Login(ctx context.Context, request *dto.LoginUserDTO) (string, error) {
	user, err := s.userFetcher.GetByEmail(ctx, request.Email)
	if err != nil {
		return "", err
	}

	if user == nil || user.Account == nil {
		return "", fmt.Errorf("user not found or invalid data")
	}

	if !model.VerifyPassword(user.Account.PasswordHash, request.Password) {
		return "", fmt.Errorf("invalid password")
	}

	token, err := auth.CreateToken(user.UserID)
	if err != nil {
		return "", err
	}

	return token, nil
}
