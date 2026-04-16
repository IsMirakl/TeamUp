package refreshsession

import (
	"net/http"

	appRefresh "backend/internal/identity/application/command/refresh_session"
	"backend/internal/identity/application/dto"
	"backend/internal/pkg/validation"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type Handler struct {
	service *appRefresh.Service
	log     *logrus.Logger
}

func NewHandler(service *appRefresh.Service, log *logrus.Logger) *Handler {
	return &Handler{service: service, log: log}
}

func (h *Handler) Handle(c *gin.Context) {
	var request dto.RefreshSessionRequest

	h.log.WithFields(logrus.Fields{
		"path":   c.FullPath(),
		"method": c.Request.Method,
	}).Info("refresh session")

	if err := c.ShouldBindJSON(&request); err != nil {
		h.log.WithError(err).Warn("failed to bind json")
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid json"})
		return
	}

	if err := validation.Validate.Struct(request); err != nil {
		h.log.WithError(err).Warn("refresh session validation failed")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp, err := h.service.RefreshSession(c.Request.Context(), request.RefreshToken)
	if err != nil {
		h.log.WithError(err).Warn("refresh session failed")
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return
	}

	c.JSON(http.StatusOK, resp)
}
