package dto

type CreatePostDTO struct {
	Title       string   `json:"title" validate:"required,max=50"`
	Description string   `json:"description" validate:"required,min=100,max=750"`
	Tags        []string `json:"tags"`
}

type UpdatePostDTO struct {
	Title       string   `json:"title" validate:"required,max=50"`
	Description string   `json:"description" validate:"required,min=100,max=750"`
	Tags        []string `json:"tags"`
}

type CreatePostResponseDTO struct {
	PostID  string `json:"post_id" validate:"required"`
	Message string `json:"message" validate:"required,max=1000"`
}
