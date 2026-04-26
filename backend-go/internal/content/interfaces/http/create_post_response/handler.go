package createpostresponse

import (
	createpostresponse "backend/internal/content/application/command/create_post_response"
	"backend/internal/content/application/dto"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgtype"
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
	}

	userID, ok := userIDInterface.(pgtype.UUID)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Invalid user id",
		})
		return
	}

	var req dto.CreatePostResponseDTO
	if err := c.ShouldBindBodyWithJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request",
		})
		return
	}

	if req.PostID == "" {
		req.PostID = c.Param("id")
	}

	response, err := h.service.Create(c.Request.Context(), userID, &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create response"})
		return
	}

	c.JSON(http.StatusCreated, response)

}
