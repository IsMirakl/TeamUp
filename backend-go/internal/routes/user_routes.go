package routes

import (
	"backend/internal/handlers"
	// "backend/internal/middleware"

	"github.com/gin-gonic/gin"
)

func CreateUser(r *gin.RouterGroup, h *handlers.UserHandler){
	
	users := r.Group("/v1/auth")
	users.POST("/register", h.Create)
}

func Login(r *gin.RouterGroup, h *handlers.UserHandler, signingKey []byte) {
	
	users := r.Group("/v1/auth")

	users.POST("/login", h.Login)
	// protected := users.Group("/", middleware.AuthMiddleware(signingKey))
	// {
	// 	protected.GET("/login", h.Login)
	// }


}