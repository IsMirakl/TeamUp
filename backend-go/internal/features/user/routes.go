package user

import (
	loginuser "backend/internal/features/user/login_user"
	registeruser "backend/internal/features/user/register_user"

	"github.com/gin-gonic/gin"
)

func UserRouter(r *gin.RouterGroup, registerHandler *registeruser.Handler, loginHandler *loginuser.Handler) {

	users := r.Group("/v1/auth")
	users.POST("/register", registerHandler.Handle)
	users.POST("/login", loginHandler.Handle)
}
