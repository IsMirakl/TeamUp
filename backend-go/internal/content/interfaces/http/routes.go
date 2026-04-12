package contenthttp

import (
	createpost "backend/internal/content/interfaces/http/create_post"
	getauthorpost "backend/internal/content/interfaces/http/get_author_post"
	getbyid "backend/internal/content/interfaces/http/get_by_id"
	updatepost "backend/internal/content/interfaces/http/update_post"
	sharedmiddleware "backend/internal/shared/middleware"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func PostRouter(
	r *gin.RouterGroup,
	createHandler *createpost.Handler,
	updateHandler *updatepost.Handler,
	getByIdHandler *getbyid.Handler,
	getAuthorHandler *getauthorpost.Handler,
	signingKey []byte,
	log *logrus.Logger,
) {
	posts := r.Group("/v1/posts")

	posts.GET("/post/:id", getByIdHandler.Handle)
	posts.GET("/post/author/:authorId", getAuthorHandler.Handle)

	protected := posts.Group("/")
	protected.Use(sharedmiddleware.AuthMiddleware(signingKey, log))

	protected.POST("/post", createHandler.Handle)
	protected.PUT("/post/:id", updateHandler.Handle)
}
