package listposts

import (
	"errors"
	"net/http"
	"strconv"

	"backend/internal/content/application/dto"
	appListPosts "backend/internal/content/application/query/list_posts"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type Handler struct {
	service *appListPosts.Service
	log     *logrus.Logger
}

func NewHandler(service *appListPosts.Service, log *logrus.Logger) *Handler {
	return &Handler{service: service, log: log}
}

func (h *Handler) Handle(c *gin.Context) {
	var limit int32
	var offset int32

	if v := c.Query("limit"); v != "" {
		n, err := strconv.Atoi(v)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid limit"})
			return
		}
		limit = int32(n)
	}

	if v := c.Query("offset"); v != "" {
		n, err := strconv.Atoi(v)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid offset"})
			return
		}
		offset = int32(n)
	}

	posts, err := h.service.ListPosts(c.Request.Context(), limit, offset)
	if err != nil {
		if errors.Is(err, appListPosts.ErrInvalidLimit) || errors.Is(err, appListPosts.ErrInvalidOffset) {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, dto.ToPostListResponse(posts))
}
