package dto

import "backend/internal/features/post/model"

func ToPostResponse(post *model.Post) *ResponsePostDTO {
	return &ResponsePostDTO{
		ID:          post.ID,
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
