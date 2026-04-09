package getbyid

import (
	"backend/internal/features/post/dto"
	"backend/internal/pkg/validation"
	appErrors "backend/internal/shared/errors"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type Handler struct {
	service *Service
	log *logrus.Logger
}

func NewHandler(service *Service, log *logrus.Logger) *Handler {
	return &Handler{service: service}
}

func (h *Handler) Handle(c *gin.Context) {
	id := c.Param("id")

	h.log.WithField("id", id).Info("GET /post by ID")

	if err := validation.Validate.Var(id, "required"); err != nil {
		h.log.WithField("id", id).
			WithError(err).
			Warn("invalid id path param")

		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid id",
		})
		return
	}

	post, err := h.service.GetById(c.Request.Context(), id)
	if err != nil {
		if errors.Is(err, appErrors.ErrPostNotFound) {
			h.log.WithField("id", id).Warn("post not found")
			c.JSON(http.StatusNotFound, gin.H{
				"error": "post not found",
			})
			return
		}

		h.log.WithField("id", id).
			WithError(err).
			Error("failed to get post by id")

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "internal server error",
		})
		return
	}

	response := dto.ToPostResponse(*post)
	c.JSON(http.StatusOK, response)
}
