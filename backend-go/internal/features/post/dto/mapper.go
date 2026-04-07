package dto

import (
	database "backend/internal/database/sqlc"
	"backend/internal/features/post/model"
)

func ToPostResponse(post database.Post) *ResponsePostDTO {
	return &ResponsePostDTO{
		ID:          post.ID.String(),
		Title:       post.Title,
		Description: post.Description,
		Tags:        post.Tags,
	}
}

func ToPostResponses(posts []model.Post) []ResponsePostDTO {
	responses := make([]ResponsePostDTO, 0, len(posts))
	for _, post := range posts {
		responses = append(responses, ResponsePostDTO{
			ID:          post.ID,
			Title:       post.Title,
			Description: post.Description,
			Tags:        post.Tags,
		})
	}

	return responses
}
