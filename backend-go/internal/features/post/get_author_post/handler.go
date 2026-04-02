package getauthorpost

import (
	"backend/internal/features/post/dto"
	"backend/internal/pkg/validation"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type Handler struct {
	service *Service
	log *logrus.Logger
}

func NewHandler(service *Service, log *logrus.Logger) *Handler {
	return &Handler{service: service, log: log}
}

func (h *Handler) Handle(c *gin.Context) {
	authorId := c.Param("authorId")

	if err := validation.Validate.Var(authorId, "required"); err != nil {

		h.log.WithField("author_id", authorId).WithError(err).Warn("failed to validate request")

		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	posts, err := h.service.GetAuthorPost(c.Request.Context(), authorId)
	if err != nil {
		h.log.WithField("author_ID", authorId).WithError(err).Error("Failed get author post")

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	response := dto.ToPostResponses(posts)
	c.JSON(http.StatusOK, response)
}
