package routes

import (
	"backend/internal/handlers"
	// "backend/internal/middleware"

	"github.com/gin-gonic/gin"
)

func UserRouter(r *gin.RouterGroup, h *handlers.UserHandler){
	
	users := r.Group("/v1/auth")
	users.POST("/register", h.Create)
	users.POST("/login", h.Login)
}