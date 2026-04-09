package getbyid

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
	return &Handler{service: service, log: log}
}


func (h *Handler) Handle(c *gin.Context) {
	userID := c.Param("userID")
	h.log.WithField("user_id", userID).Info("GET /user by id")

	if err := validation.Validate.Var(userID, "required"); err != nil {
		
		h.log.WithError(err).Warn("handler error")

		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	user, err := h.service.GetById(c.Request.Context(), userID)
	if err != nil {
		
		h.log.WithError(err).Error("handler failed get user by id")

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	responses := dto.ToUserResponse(user)
	c.JSON(http.StatusOK, responses)
}