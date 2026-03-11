package routes

import (
	"backend/internal/handlers"
	"backend/internal/middleware"

	"github.com/gin-gonic/gin"
)

func CreateUser(r *gin.RouterGroup, h *handlers.UserHandler, signingKey []byte){
	
	users := r.Group("/users")

	protected := users.Group("/")
	protected.Use(middleware.AuthMiddleware(signingKey))

	users.POST("/register", h.Create)
}
