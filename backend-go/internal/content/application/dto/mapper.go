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
		Author:      "",
	}
}

func ToPostByIDResponse(post database.GetPostByIdRow) *ResponsePostDTO {
	return &ResponsePostDTO{
		ID:          post.ID.String(),
		Title:       post.Title,
		Description: post.Description,
		Tags:        post.Tags,
		Author:      post.AuthorName,
	}
}

func ToPostUpdateResponse(post database.UpdatePostRow) *ResponsePostDTO {
	return &ResponsePostDTO{
		ID:          post.ID.String(),
		Title:       post.Title,
		Description: post.Description,
		Tags:        post.Tags,
		Author:      "",
	}
}

func ToPostResponses(posts database.Post) *ResponsePostDTO {
	return &ResponsePostDTO{
		ID:          posts.ID.String(),
		Title:       posts.Title,
		Description: posts.Description,
		Tags:        posts.Tags,
		Author:      "",
	}
}

func ToPostListResponse(posts []database.ListPostsRow) []ResponsePostDTO {
	result := make([]ResponsePostDTO, 0, len(posts))
	for _, post := range posts {
		result = append(result, ResponsePostDTO{
			ID:          post.ID.String(),
			Title:       post.Title,
			Description: post.Description,
			Tags:        post.Tags,
			Author:      post.AuthorName,
		})
	}
	return result
}
