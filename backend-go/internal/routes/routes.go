package routes

import (
	"backend/internal/handlers"

	"github.com/gin-gonic/gin"
)

type Routes struct {
	UserHandler *handlers.UserHandler
	TeamSeekPostHandler *handlers.TeamSeekPostHandler
	signingKey []byte
}

func SetupRouter(r *gin.Engine, h *Routes) {
	
	api := r.Group("/api")

	UserRouter(api, h.UserHandler)
	PostRouter(api, h.TeamSeekPostHandler)
}