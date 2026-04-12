package updatepost

import (
	"net/http"

	appUpdatePost "backend/internal/content/application/command/update_post"
	"backend/internal/content/application/dto"
	"backend/internal/pkg/validation"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *appUpdatePost.Service
}

func NewHandler(service *appUpdatePost.Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) Handle(c *gin.Context) {
	post_id := c.Param("id")
	if err := validation.Validate.Var(post_id, "required"); err != nil {
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

	post, err := h.service.Update(c.Request.Context(), post_id, &request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	response := dto.ToPostUpdateResponse(*post)
	c.JSON(http.StatusOK, response)
}
