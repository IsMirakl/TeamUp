package dto

import (
	database "backend/internal/database/sqlc"
)

func ToPostResponse(post database.Post) *ResponsePostDTO {
	return &ResponsePostDTO{
		ID:          post.ID.String(),
		Title:       post.Title,
		Description: post.Description,
		Tags:        post.Tags,
	}
}

func ToPostUpdateResponse(post database.UpdatePostRow) *ResponsePostDTO {
	return &ResponsePostDTO{
		ID:          post.ID.String(),
		Title:       post.Title,
		Description: post.Description,
		Tags:        post.Tags,
	}
}

func ToPostResponses(posts database.Post) *ResponsePostDTO {
	return &ResponsePostDTO{
		ID:          posts.ID.String(),
		Title:       posts.Title,
		Description: posts.Description,
		Tags:        posts.Tags,
	}
}

