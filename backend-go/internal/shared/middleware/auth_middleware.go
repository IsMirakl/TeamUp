package middleware

import (
	auth "backend/internal/pkg/utils"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func AuthMiddleware(tokenService auth.TokenService, log *logrus.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Info("auth middleware triggered")

		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			log.Warn("missing Authorization header")

			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "missing Authorization header",
			})
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		if tokenString == authHeader {
			log.Warn("invalid Authorization header format")

			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "invalid Authorization header format",
			})
			return
		}

		claims, err := tokenService.ValidateAccessToken(tokenString)
		if err != nil {
			log.WithError(err).Error("failed to validate token")

			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.Set("userID", claims.UserID)

		c.Next()
	}

}
