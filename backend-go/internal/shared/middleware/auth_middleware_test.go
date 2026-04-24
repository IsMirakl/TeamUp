package middleware_test

import (
	auth "backend/internal/pkg/utils"
	"backend/internal/pkg/utils/mocks"
	"backend/internal/shared/middleware"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAuthMiddleware_Success(t *testing.T) {
	tokenService := &mocks.MockTokenService{}

	log := logrus.New()
	router := gin.New()

	router.GET("/test", middleware.AuthMiddleware(tokenService, log), func(c *gin.Context) {
		userID, exists := c.Get("userID")
		require.True(t, exists)
		require.Equal(t, "user-123", userID)

		c.Status(http.StatusOK)
	})

	tokenService.On("ValidateAccessToken", "valid-token").Return(&auth.Claims{UserID: "user-123"}, nil)

	req := httptest.NewRequest(http.MethodGet, "/test", nil)
	req.Header.Set("Authorization", "Bearer valid-token")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestAuthMiddleware_NoHeader(t *testing.T) {
	tokenService := &mocks.MockTokenService{}

	log := logrus.New()
	router := gin.New()

	router.GET("/test", middleware.AuthMiddleware(tokenService, log), func(c *gin.Context) {
		c.Status(http.StatusOK)
	})

	req := httptest.NewRequest(http.MethodGet, "/test", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusUnauthorized, w.Code)
}

func TestAuthMiddleware_NoCorrectFormat(t *testing.T) {
	tokenService := &mocks.MockTokenService{}

	log := logrus.New()
	router := gin.New()

	router.GET("/test", middleware.AuthMiddleware(tokenService, log), func(c *gin.Context) {
		userID, exists := c.Get("userID")
		require.True(t, exists)
		require.Equal(t, "user-123", userID)

		c.Status(http.StatusOK)
	})

	tokenService.On("ValidateAccessToken", "valid-token").Return(&auth.Claims{UserID: "user-123"}, nil)

	req := httptest.NewRequest(http.MethodGet, "/test", nil)
	req.Header.Set("Authorization", "invalid-token")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusUnauthorized, w.Code)
}

func TestAuthMiddleware_InvalidToke(t *testing.T) {
	tokenService := &mocks.MockTokenService{}

	log := logrus.New()
	router := gin.New()

	router.GET("/test", middleware.AuthMiddleware(tokenService, log), func(c *gin.Context) {
		userID, exists := c.Get("userID")
		require.True(t, exists)
		require.Equal(t, "user-123", userID)

		c.Status(http.StatusOK)
	})

	tokenService.On("ValidateAccessToken", "bad-token").Return(nil, errors.New("invalid-token"))

	req := httptest.NewRequest(http.MethodGet, "/test", nil)
	req.Header.Set("Authorization", "invalid-token")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusUnauthorized, w.Code)
}
