package handlers

import (
	postDTO "backend/internal/dto/team_seek_post"
	appErrors "backend/internal/errors"
	"backend/internal/service"

	"backend/internal/validation"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TeamSeekPostHandler struct {
	service *service.TeamSeekPostService
}

func NewTeamSeekPostHandler(service *service.TeamSeekPostService) *TeamSeekPostHandler {
	return &TeamSeekPostHandler{
		service: service,
	}
}

func (h *TeamSeekPostHandler) Create(c *gin.Context) {
	userIDInterface, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "user_id not found in context",
		})
		return
	}

	userID, ok := userIDInterface.(string)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "invalid user_id type",
		})
		return
	}

	var dto postDTO.CreateTeamSeekPostDTO

	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid json",
		})
		return
	}

	if err := validation.Validate.Struct(dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	post, err := h.service.Create(c.Request.Context(), &dto, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	response := postDTO.ToTeamSeekPostResponse(post)
	c.JSON(http.StatusCreated, response)
}


func (h *TeamSeekPostHandler) Update(c *gin.Context) {

	var dto postDTO.UpdateTeamSeekPostDTO

	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid json",
		})
		return
	}

	if err := validation.Validate.Struct(dto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	post, err := h.service.Update(c.Request.Context(), &dto)

	if err = validation.Validate.Struct(dto); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	
	response := postDTO.ToTeamSeekPostResponse(post)
	c.JSON(http.StatusOK, response)
}

func (h *TeamSeekPostHandler) GetPostById(c *gin.Context) {

    id := c.Param("id")

    if err := validation.Validate.Var(id, "required"); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "error": "invalid id",
        })
        return
    }

    post, err := h.service.GetPostById(c.Request.Context(), id)
    if err != nil {
        if errors.Is(err, appErrors.ErrPostNotFound) {
            c.JSON(http.StatusNotFound, gin.H{
                "error": "post not found",
            })
            return
        }

        c.JSON(http.StatusInternalServerError, gin.H{
            "error": "internal server error",
        })
        return
    }

    response := postDTO.ToTeamSeekPostResponse(post)
    c.JSON(http.StatusOK, response)
}

func (h *TeamSeekPostHandler) GetAuthorPost(c *gin.Context) {
	
	authorId := c.Param("authorId")

	if err := validation.Validate.Var(authorId, "required"); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	author, err := h.service.GetPostById(c.Request.Context(), authorId)
	if err = validation.Validate.Var(authorId, "required"); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	response := postDTO.ToTeamSeekPostResponse(author)
	c.JSON(http.StatusOK, response)
}