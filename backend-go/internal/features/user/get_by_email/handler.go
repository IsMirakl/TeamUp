package getbyemail

import (
	"backend/internal/features/user/dto"
	"backend/internal/pkg/validation"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type Handler struct {
	service *Service
	log *logrus.Logger
}

func NewHandler(service *Service, log *logrus.Logger) *Handler {
	return &Handler{
		service: service,
		log: log,
	}
}


func (h *Handler) Handle(c *gin.Context) {
	var request dto.LoginUserDTO

	h.log.WithField("email", request.Email).Info("Get by id request received")
	
	if err := validation.Validate.Var(request.Email, "required"); err != nil {

		h.log.WithField("email", request.Email).WithError(err).Warn("failed to validate request")

		c.JSON(http.StatusBadRequest, gin.H{
			"erro": err.Error(),
		})
		return
	}

	user, err := h.service.GetByEmail(c.Request.Context(), request.Email)
	if err != nil {
		h.log.WithField("email", request.Email).WithError(err).Error("Failed get user by email")
	
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	response := dto.ToUserResponse(user)
	c.JSON(http.StatusOK, response)

}