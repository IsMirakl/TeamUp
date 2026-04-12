package getmyprofile

import (
	appProfile "backend/internal/identity/application/query/get_my_profile"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/sirupsen/logrus"
)

type Handler struct {
	service *appProfile.Service
	log     *logrus.Logger
}

func NewProfileHandler(service *appProfile.Service, log *logrus.Logger) *Handler {
	return &Handler{
		service: service,
		log:     log,
	}
}

func (h *Handler) Handle(c *gin.Context) {
	h.log.Info("GET /profile/me")

	rawUserID, exists := c.Get("id")

	if !exists {
		h.log.Error("User ID not found in context")
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	userID, err := parseUserID(rawUserID)
	if err != nil {
		h.log.WithError(err).Error("Failed to parse userID")
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	profile, err := h.service.GetMe(c.Request.Context(), userID)
	if err != nil {
		h.log.WithError(err).Error("Failed to get profile")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get profile"})
		return
	}
	c.JSON(http.StatusOK, profile)
}

func parseUserID(rawUserID interface{}) (pgtype.UUID, error) {
	switch v := rawUserID.(type) {
	case pgtype.UUID:
		return v, nil
	case uuid.UUID:
		return pgtype.UUID{Bytes: v, Valid: true}, nil
	case string:
		parsed, err := uuid.Parse(v)
		if err != nil {
			return pgtype.UUID{}, err
		}
		return pgtype.UUID{Bytes: parsed, Valid: true}, nil
	default:
		return pgtype.UUID{}, errors.New("unsupported userID type")
	}
}
