package getpostresponses

import (
	"errors"
	"net/http"

	getpostresponses "backend/internal/content/application/query/get_post_responses"
	"backend/internal/pkg/validation"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type Handler struct {
	service *getpostresponses.Service
	log     *logrus.Logger
}

func NewHandler(service *getpostresponses.Service, log *logrus.Logger) *Handler {
	return &Handler{
		service: service,
		log:     log,
	}
}

func (h *Handler) Handle(c *gin.Context) {
	postID := c.Param("id")

	h.log.WithField("post_id", postID).Info("GET /post responses")

	if err := validation.Validate.Var(postID, "required"); err != nil {
		h.log.WithField("post_id", postID).WithError(err).Warn("invalid post_id path param")
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid post_id"})
		return
	}

	responses, err := h.service.GetPostResponses(c.Request.Context(), postID)
	if err != nil {
		if errors.Is(err, getpostresponses.ErrInvalidPostID) {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid post_id"})
			return
		}

		h.log.WithField("post_id", postID).WithError(err).Error("failed to get post responses")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}

	c.JSON(http.StatusOK, responses)
}
