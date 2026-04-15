package auth

import (
	"backend/internal/pkg/config"
	"fmt"
	"log"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

type Claims struct {
	UserID string `json:"user_id"`
	jwt.RegisteredClaims
	log *logrus.Logger
}

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("No .env file found")
	}
}

func CreateToken(userID string, log *logrus.Logger) (string, error) {
	conf := config.New(log)

	signingKey := []byte(conf.SECRET_KEY.JWT_SECRET)

	claims := Claims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(72 * time.Hour)),
			Issuer:    "TeamUP",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(signingKey)
	if err != nil {
		return "", fmt.Errorf("Sign token: %w", err)
	}

	return tokenString, nil
}

func ValidateToken(tokenString string, signingKey []byte) (*Claims, error) {

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
	if !ok {
		return nil, fmt.Errorf("invalid token")

	}
	return claims, nil
}

func GenerateRefreshToken(userID string, log *logrus.Logger) (string, error) {
	conf := config.New(log)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userID,
		"exp": time.Now().Add(time.Hour * 24 * 90).Unix(),
	})

	signingRefreshKey := []byte(conf.SECRET_KEY.REFRESH_SECRET)

	refreshToken, err := token.SignedString(signingRefreshKey)
	if err != nil {
		log.WithFields(logrus.Fields{
			"user_id": userID,
		}).WithError(err).Error("Failed to sign refresh token")

		return "", fmt.Errorf("sign refresh token: %w", err)
	}

	log.WithFields(logrus.Fields{
		"user_id": userID,
	}).Info("refresh token generated")

	return refreshToken, nil
}


func ValidateRefreshToken(refreshToken string, signingRefreshKey []byte) (*Claims, error) {
	token, err := jwt.ParseWithClaims(
		refreshToken,
		&Claims{},
		func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method")
			}
			return signingRefreshKey, nil
		},
	)

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*Claims)
	if !ok {
		return nil, fmt.Errorf("invalid token")
	}
	return claims, nil
}