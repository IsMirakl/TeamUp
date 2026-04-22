package auth_test

import (
	auth "backend/internal/pkg/utils"
	"testing"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func createTokenService() auth.TokenService {
	log := logrus.New()

	return auth.NewTokenService(
		"access-secret-key",
		"refresh-secret-key",
		"TestIssuer",
		log,
	)
}

func createOtherTokenService() auth.TokenService {
	log := logrus.New()

	return auth.NewTokenService(
		"another-secret-key",
		"another-secret-key",
		"TestIssuer",
		log,
	)
}

func TestGenerateAccessToken_Success(t *testing.T) {
	service := createTokenService()

	token, err := service.GenerateAccessToken("user-123")

	require.NoError(t, err)
	require.NotEmpty(t, token)
}

func TestValidateAccessToken_Success(t *testing.T) {
	service := createTokenService()

	token, err := service.GenerateAccessToken("user-123")

	require.NoError(t, err)
	require.NotEmpty(t, token)

	claims, err := service.ValidateAccessToken(token)

	require.NoError(t, err)
	require.NotNil(t, claims)

	assert.Equal(t, "user-123", claims.UserID)
	assert.NotNil(t, claims.ExpiresAt)
	assert.True(t, claims.ExpiresAt.Time.After(time.Now()))
}

func TestGenerateRefreshToken_Success(t *testing.T) {
	service := createTokenService()

	token, err := service.GenerateRefreshToken("user-123")

	require.NoError(t, err)
	require.NotEmpty(t, token)
}

func TestValidateRefreshToken_Success(t *testing.T) {
	service := createTokenService()

	token, err := service.GenerateRefreshToken("user-123")

	require.NoError(t, err)
	require.NotEmpty(t, token)

	claims, err := service.ValidateRefreshToken(token)

	require.NoError(t, err)
	require.NotEmpty(t, claims)

	assert.Equal(t, "user-123", claims.UserID)
	assert.NotNil(t, claims.ExpiresAt)
	assert.True(t, claims.ExpiresAt.Time.After(time.Now()))
}

func TestValidateAccessToken_InvalidSignature(t *testing.T) {
	service := createTokenService()

	token, err := service.GenerateAccessToken("user-123")
	require.NoError(t, err)

	otherService := createOtherTokenService()
	claims, err := otherService.ValidateAccessToken(token)

	assert.Error(t, err)
	assert.Nil(t, claims)
}

func TestValidateRefreshToken_InvalidSignature(t *testing.T) {
	service := createTokenService()

	token, err := service.GenerateRefreshToken("user-123")
	require.NoError(t, err)

	otherService := createOtherTokenService()
	claims, err := otherService.ValidateRefreshToken(token)

	assert.Error(t, err)
	assert.Nil(t, claims)
}
