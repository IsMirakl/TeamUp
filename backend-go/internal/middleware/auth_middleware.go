package middleware

import (
	"backend/internal/auth"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware(signingKey []byte) gin.HandlerFunc {
	return func(c *gin.Context){
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "missing Authorization header",
			})
			return 
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		if tokenString == authHeader {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "invalid Authorization header format",
			})
			return
		}

		claims, err := auth.ValidateToken(tokenString, signingKey)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.Set("userID", claims.UserID)

		c.Next()
	}
	
}