package routes

import (
	"backend/internal/handlers"
	authMiddleware "backend/internal/middleware"

	"github.com/gin-gonic/gin"
)

func PostRouter(r *gin.RouterGroup, h *handlers.TeamSeekPostHandler, signingKey []byte) {
	posts := r.Group("/v1/posts")

	posts.GET("/post/:id", h.GetPostById)
	posts.GET("/post/author/:id", h.GetAuthorPost)

	protected := posts.Group("/")
	protected.Use(authMiddleware.AuthMiddleware(signingKey))

	protected.POST("/post", h.Create)
	protected.PUT("/post/:id", h.Update)
}