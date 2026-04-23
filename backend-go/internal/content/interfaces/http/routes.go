package contenthttp

import (
	createpost "backend/internal/content/interfaces/http/create_post"
	getauthorpost "backend/internal/content/interfaces/http/get_author_post"
	getbyid "backend/internal/content/interfaces/http/get_by_id"
	listposts "backend/internal/content/interfaces/http/list_posts"
	updatepost "backend/internal/content/interfaces/http/update_post"
	auth "backend/internal/pkg/utils"
	"backend/internal/shared/middleware"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type RouterParams struct {
	tokenService auth.TokenService
	log          *logrus.Logger
}

func NewRouterParams(tokenService auth.TokenService, log *logrus.Logger) RouterParams {
	return RouterParams{
		tokenService: tokenService,
		log:          log,
	}
}

func PostRouter(
	r *gin.RouterGroup,
	createHandler *createpost.Handler,
	updateHandler *updatepost.Handler,
	getByIdHandler *getbyid.Handler,
	getAuthorHandler *getauthorpost.Handler,
	listPostsHandler *listposts.Handler,
	params RouterParams,
	log *logrus.Logger,
) {
	posts := r.Group("/v1/posts")

	posts.GET("/post", listPostsHandler.Handle)
	posts.GET("/post/:id", getByIdHandler.Handle)
	posts.GET("/post/author/:authorId", getAuthorHandler.Handle)

	protected := posts.Group("/")
	protected.Use(middleware.AuthMiddleware(params.tokenService, params.log))

	protected.POST("/post", createHandler.Handle)
	protected.PUT("/post/:id", updateHandler.Handle)
}
