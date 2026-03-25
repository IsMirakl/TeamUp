package post

import (
	createpost "backend/internal/features/post/create_post"
	getauthorpost "backend/internal/features/post/get_author_post"
	getbyid "backend/internal/features/post/get_by_id"
	updatepost "backend/internal/features/post/update_post"
	sharedmiddleware "backend/internal/shared/middleware"

	"github.com/gin-gonic/gin"
)

func PostRouter(
	r *gin.RouterGroup,
	createHandler *createpost.Handler,
	updateHandler *updatepost.Handler,
	getByIdHandler *getbyid.Handler,
	getAuthorHandler *getauthorpost.Handler,
	signingKey []byte,
) {
	posts := r.Group("/v1/posts")

	posts.GET("/post/:id", getByIdHandler.Handle)
	posts.GET("/post/author/:authorId", getAuthorHandler.Handle)

	protected := posts.Group("/")
	protected.Use(sharedmiddleware.AuthMiddleware(signingKey))

	protected.POST("/post", createHandler.Handle)
	protected.PUT("/post/:id", updateHandler.Handle)
}
