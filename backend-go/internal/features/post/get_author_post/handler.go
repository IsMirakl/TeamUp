package getauthorpost

import (
	"backend/internal/features/post/dto"
	"backend/internal/pkg/validation"
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
	authorId := c.Param("authorId")
	if err := validation.Validate.Var(authorId, "required"); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	posts, err := h.service.GetAuthorPost(c.Request.Context(), authorId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	response := dto.ToPostResponses(posts)
	c.JSON(http.StatusOK, response)
}
