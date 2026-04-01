package registeruser

import (
	"backend/internal/features/user/dto"
	"backend/internal/pkg/validation"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type Handler struct {
	service *Service
	log     *logrus.Logger
}

func NewUserHandler(service *Service, log *logrus.Logger) *Handler {
	return &Handler{
		service: service,
		log: log,
	}
}
func (h *Handler) Handle(c *gin.Context) {

	var request dto.CreateUserDTO

	h.log.WithFields(logrus.Fields{
		"path":   c.FullPath(),
		"method": c.Request.Method,
	}).Info("register user request received")

	if err := c.ShouldBindJSON(&request); err != nil {
		h.log.WithFields(logrus.Fields{
			"path": c.FullPath(),
			"method": c.Request.Method,
		}).WithError(err).Warn("failed to bind json")

		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid json",
		})
		return
	}

	if err := validation.Validate.Struct(request); err != nil {
		h.log.WithFields(logrus.Fields{
			"email": request.Email,
		}).Warn("register user validation failed")

		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	user, err := h.service.Create(c.Request.Context(), &request)
	
	if err != nil {
		h.log.WithFields(logrus.Fields{
			"email": request.Email,
		}).WithError(err).Error("failed to create user")

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	h.log.WithFields(logrus.Fields{
		"user_id": user.UserID,
		"email": user.Email,
	}).Info("User created successfully")

	response := dto.ToUserResponse(user)
	c.JSON(http.StatusCreated, response)
}
