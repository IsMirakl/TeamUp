package listposts

import (
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
	limit := int32(20)
	offset := int32(0)

	if v := c.Query("limit"); v != "" {
		n, err := strconv.Atoi(v)
		if err != nil || n <= 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid limit"})
			return
		}
		if n > 100 {
			n = 100
		}
		limit = int32(n)
	}

	if v := c.Query("offset"); v != "" {
		n, err := strconv.Atoi(v)
		if err != nil || n < 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid offset"})
			return
		}
		offset = int32(n)
	}

	posts, err := h.service.ListPosts(c.Request.Context(), limit, offset)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, dto.ToPostListResponse(posts))
}
