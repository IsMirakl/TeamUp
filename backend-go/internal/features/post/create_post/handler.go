package createpost

import (
	"backend/internal/features/post/dto"
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
	userIDInterface, exists := c.Get("userID")

	if !exists {
		h.log.WithFields(logrus.Fields{
			"user_id": userIDInterface,
		}).Error("UserID not found in context")

		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "user_id not found in context",
		})
		return
	}

	userID, ok := userIDInterface.(string)
	if !ok {
		h.log.WithField("iser_id", userID).Error("Invalid user_id type")

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "invalid user_id type",
		})
		return
	}

	var request dto.CreatePostDTO
	
	h.log.WithFields(logrus.Fields{
		"path":   c.FullPath(),
		"method": c.Request.Method,
	}).Info("create post request received")

	if err := c.ShouldBindJSON(&request); err != nil {
		h.log.WithFields(logrus.Fields{
			"path": c.FullPath(),
			"method": c.Request.Method,
		}).WithError(err).Warn("failed to bind json")

		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid json",
		})
		return
	}

	if err := validation.Validate.Struct(request); err != nil {
		h.log.Warn("create post validation failed")

		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	post, err := h.service.Create(c.Request.Context(), &request, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	response := dto.ToPostResponse(*post)
	c.JSON(http.StatusCreated, response)
}
