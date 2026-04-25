package getmyprofile

import (
	"backend/internal/identity/application/dto"
	appProfile "backend/internal/identity/application/query/get_my_profile"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type Handler struct {
	service *appProfile.Service
	log     *logrus.Logger
}

func NewProfileHandler(service *appProfile.Service, log *logrus.Logger) *Handler {
	return &Handler{
		service: service,
		log:     log,
	}
}

func (h *Handler) Handle(c *gin.Context) {
	h.log.Info("GET /profile/me")

	rawUserID, exists := c.Get("userID")

	if !exists {
		h.log.Error("User ID not found in context")
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	userID, ok := rawUserID.(string)
	if !ok {
		h.log.WithField("user_id_type", fmt.Sprintf("%T", rawUserID)).Error("Failed to cast userID to string")
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	profile, err := h.service.GetMe(c.Request.Context(), userID)
	if err != nil {
		h.log.WithError(err).Error("Failed to get profile")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get profile"})
		return
	}
	response := dto.ToProfileResponse(profile)
	c.JSON(http.StatusOK, response)
}
