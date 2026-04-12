package loginuser

import (
	"net/http"

	appLogin "backend/internal/identity/application/command/login_user"
	"backend/internal/identity/application/dto"
	"backend/internal/pkg/validation"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type Handler struct {
	service *appLogin.Service
	log     *logrus.Logger
}

func NewUserHandler(service *appLogin.Service, log *logrus.Logger) *Handler {
	return &Handler{
		service: service,
		log:     log,
	}
}

func (h *Handler) Handle(c *gin.Context) {
	var request dto.LoginUserDTO

	h.log.WithField("email", request.Email).Info("GET /login")

	if err := c.ShouldBindJSON(&request); err != nil {
		h.log.WithError(err).Error("failed to bind json")

		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid json",
		})
		return
	}

	if err := validation.Validate.Struct(request); err != nil {
		h.log.WithError(err).Warn("failed to validate request")

		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	token, err := h.service.Login(c.Request.Context(), &request)
	if err != nil {
		h.log.WithField("email", request.Email).Error("failed to login user")

		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Invalid credentials",
		})
		return
	}

	response := dto.LoginResponse{AccessToken: token}
	c.JSON(http.StatusOK, response)
}
