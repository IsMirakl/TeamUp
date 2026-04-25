package identityhttp

import (
	getUserByEmail "backend/internal/identity/interfaces/http/get_by_email"
	getUserById "backend/internal/identity/interfaces/http/get_by_id"
	getMyProfile "backend/internal/identity/interfaces/http/get_my_profile"
	loginuser "backend/internal/identity/interfaces/http/login_user"
	refreshsession "backend/internal/identity/interfaces/http/refresh_session"
	registeruser "backend/internal/identity/interfaces/http/register_user"
	revokesession "backend/internal/identity/interfaces/http/revoke_session"
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

func UserRouter(
	r *gin.RouterGroup,
	registerHandler *registeruser.Handler,
	loginHandler *loginuser.Handler,
	refreshHandler *refreshsession.Handler,
	revokeSessionHandler *revokesession.Handler,
	getUserById *getUserById.Handler,
	getUserByEmail *getUserByEmail.Handler,
	params RouterParams,
	getMyProfile *getMyProfile.Handler,
) {

	users_auth := r.Group("/v1/auth")
	users := r.Group("/v1")

	users_auth.POST("/register", registerHandler.Handle)
	users_auth.POST("/login", loginHandler.Handle)
	users_auth.POST("/refresh", refreshHandler.Handle)
	users.GET("/user/:userID", getUserById.Handle)
	users.GET("/user/email/:email", getUserByEmail.Handle)

	profile := r.Group("/v1/profile")
	profile.Use(middleware.AuthMiddleware(params.tokenService, params.log))
	profile.GET("/me", getMyProfile.Handle)
}
