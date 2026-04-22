package auth

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/sirupsen/logrus"
)

type Claims struct {
	UserID string `json:"user_id"`
	jwt.RegisteredClaims
}

type TokenService interface {
	GenerateAccessToken(userID string) (string, error)
	GenerateRefreshToken(userID string) (string, error)
	ValidateAccessToken(tokenString string) (*Claims, error)
	ValidateRefreshToken(tokenString string) (*Claims, error)
}

type tokenService struct {
	accessSecret  []byte
	refreshSecret []byte
	log           *logrus.Logger
	issuer        string
}

func NewTokenService(
	accessSecret string,
	refreshSecret string,
	issuer string,
	log *logrus.Logger,
) TokenService {
	return &tokenService{
		accessSecret:  []byte(accessSecret),
		refreshSecret: []byte(refreshSecret),
		log:           log,
		issuer:        issuer,
	}
}

func (s *tokenService) GenerateAccessToken(userID string) (string, error) {
	claims := Claims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(72 * time.Hour)),
			Issuer:    s.issuer,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(s.accessSecret)
	if err != nil {
		return "", fmt.Errorf("Sign access token: %w", err)
	}

	return tokenString, nil
}

func (s *tokenService) GenerateRefreshToken(userID string) (string, error) {
	claims := Claims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(90 * 24 * time.Hour)),
			Issuer:    s.issuer,
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	refreshToken, err := token.SignedString(s.refreshSecret)
	if err != nil {
		s.log.WithFields(logrus.Fields{
			"user_id": userID,
		}).WithError(err).Error("Failed to sign refresh token")

		return "", fmt.Errorf("sign refresh token: %w", err)
	}

	s.log.WithFields(logrus.Fields{
		"user_id": userID,
	}).Info("refresh token generated")

	return refreshToken, nil
}

func (s *tokenService) ValidateAccessToken(tokenString string) (*Claims, error) {
	return validateToken(tokenString, s.accessSecret)
}

func (s *tokenService) ValidateRefreshToken(tokenString string) (*Claims, error) {
	return validateToken(tokenString, s.refreshSecret)
}

func validateToken(tokenString string, signingKey []byte) (*Claims, error) {
	token, err := jwt.ParseWithClaims(
		tokenString,
		&Claims{},
		func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method")
			}
			return signingKey, nil
		},
	)

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*Claims)
	if !ok || !token.Valid {
		return nil, fmt.Errorf("invalid token")

	}
	return claims, nil
}
