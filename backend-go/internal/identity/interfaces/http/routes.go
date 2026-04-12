package identityhttp

import (
	getUserByEmail "backend/internal/identity/interfaces/http/get_by_email"
	getUserById "backend/internal/identity/interfaces/http/get_by_id"
	loginuser "backend/internal/identity/interfaces/http/login_user"
	registeruser "backend/internal/identity/interfaces/http/register_user"

	"github.com/gin-gonic/gin"
)

func UserRouter(r *gin.RouterGroup, registerHandler *registeruser.Handler, loginHandler *loginuser.Handler, getUserById *getUserById.Handler, getUserByEmail *getUserByEmail.Handler) {

	users_auth := r.Group("/v1/auth")
	users := r.Group("/v1")

	users_auth.POST("/register", registerHandler.Handle)
	users_auth.POST("/login", loginHandler.Handle)
	users.GET("/user/:userID", getUserById.Handle)
	users.GET("/user/email/:email", getUserByEmail.Handle)
}
