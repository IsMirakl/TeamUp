package routes

import (
	"backend/internal/handlers"

	"github.com/gin-gonic/gin"
)

// import (
// 	"backend/internal/handlers"

// 	"github.com/gin-gonic/gin"
// )


func PostRouter(r *gin.RouterGroup, h *handlers.TeamSeekPostHandler) {

	posts := r.Group("/v1/posts")
	posts.POST("/post", h.Create)
	posts.PUT("/post/:id", h.Update)
	posts.GET("/post/:id", h.GetPostById)
	posts.GET("/post/author/:id", h.GetAuthorPost)

}