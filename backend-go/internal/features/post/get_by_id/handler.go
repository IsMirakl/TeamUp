package getbyid

import (
	"backend/internal/features/post/dto"
	"backend/internal/pkg/validation"
	appErrors "backend/internal/shared/errors"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) Handle(c *gin.Context) {
	id := c.Param("id")
	
	if err := validation.Validate.Var(id, "required"); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid id",
		})
		return
	}

	post, err := h.service.GetById(c.Request.Context(), id)
	if err != nil {
		if errors.Is(err, appErrors.ErrPostNotFound) {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "post not found",
			})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "internal server error",
		})
		return
	}

	response := dto.ToPostResponse(post)
	c.JSON(http.StatusOK, response)
}
