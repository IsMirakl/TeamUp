package routes

import (
	"backend/internal/handlers"

	"github.com/gin-gonic/gin"
)

type Routes struct {
	UserHandler *handlers.UserHandler
	signingKey []byte
}

func SetupRouter(r *gin.Engine, h *Routes) {
	
	api := r.Group("/api")

	CreateUser(api, h.UserHandler)
	Login(api, h.UserHandler, h.signingKey)
}