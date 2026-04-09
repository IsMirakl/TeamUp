package getbyemail

import (
	"backend/internal/features/user/dto"
	"backend/internal/pkg/validation"
	"database/sql"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type Handler struct {
	service *Service
	log     *logrus.Logger
}

func NewHandler(service *Service, log *logrus.Logger) *Handler {
	return &Handler{
		service: service,
		log:     log,
	}
}

func (h *Handler) Handle(c *gin.Context) {
	email := c.Param("email")
	if err := validation.Validate.Var(email, "required,email"); err != nil {
		h.log.WithField("email", email).
			WithError(err).
			Warn("invalid email path param")

		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	h.log.WithField("email", email).Info("getting user by email")

	user, err := h.service.GetByEmail(c.Request.Context(), email)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			h.log.WithField("email", email).Warn("user not found")

			c.JSON(http.StatusNotFound, gin.H{
				"error": "user not found",
			})
			return
		}

		h.log.WithField("email", email).
			WithError(err).
			Error("failed to get user by email")

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "internal server error",
		})
		return
	}

	response := dto.ToUserResponse(&user)
	c.JSON(http.StatusOK, response)
}
