package getbyemail

import (
	"database/sql"
	"errors"
	"net/http"

	"backend/internal/identity/application/dto"
	appGetByEmail "backend/internal/identity/application/query/get_by_email"
	"backend/internal/pkg/validation"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type Handler struct {
	service *appGetByEmail.Service
	log     *logrus.Logger
}

func NewHandler(service *appGetByEmail.Service, log *logrus.Logger) *Handler {
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
