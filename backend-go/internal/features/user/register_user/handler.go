package registeruser

import (
	"backend/internal/features/user/dto"
	"backend/internal/pkg/validation"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *Service
}

func NewUserHandler(service *Service) *Handler {
	return &Handler{
		service: service,
	}
}
func (h *Handler) Handle(c *gin.Context) {

	var request dto.CreateUserDTO

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

	user, err := h.service.Create(c.Request.Context(), &request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	response := dto.ToUserResponse(user)
	c.JSON(http.StatusCreated, response)
}
