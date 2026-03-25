package updatepost

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
	id := c.Param("id")
	if err := validation.Validate.Var(id, "required"); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid id",
		})
		return
	}

	var request dto.UpdatePostDTO
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid json",
		})
		return
	}

	if err := validation.Validate.Struct(request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	post, err := h.service.Update(c.Request.Context(), id, &request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	response := dto.ToPostResponse(post)
	c.JSON(http.StatusOK, response)
}
