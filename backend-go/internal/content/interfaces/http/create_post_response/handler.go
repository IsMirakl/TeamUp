package createpostresponse

import (
	createpostresponse "backend/internal/content/application/command/create_post_response"
	"backend/internal/content/application/dto"
	"backend/internal/pkg/validation"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type Handler struct {
	service *createpostresponse.Service
	log     *logrus.Logger
}

func NewHandler(service *createpostresponse.Service, log *logrus.Logger) *Handler {
	return &Handler{
		service: service,
		log:     log,
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
		h.log.WithFields(logrus.Fields{
			"user_id_value": userIDInterface,
			"user_id_type":  fmt.Sprintf("%T", userIDInterface),
		}).Error("Invalid user_id type")

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "invalid user_id type",
		})
		return
	}

	var req dto.CreatePostResponseDTO
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request",
		})
		return
	}

	if req.PostID == "" {
		req.PostID = c.Param("id")
	}

	if err := validation.Validate.Struct(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	response, err := h.service.Create(c.Request.Context(), userID, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create response"})
		return
	}

	c.JSON(http.StatusCreated, response)

}
