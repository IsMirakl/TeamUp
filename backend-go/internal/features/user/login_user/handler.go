package loginuser

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
	var request dto.LoginUserDTO

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

	token, err := h.service.Login(c.Request.Context(), &request)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Invalid credentials",
		})
		return
	}

	response := dto.LoginResponse{AccessToken: token}
	c.JSON(http.StatusOK, response)
}
