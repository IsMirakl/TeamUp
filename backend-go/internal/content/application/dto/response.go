package dto

type ResponsePostDTO struct {
	ID          string   `json:"id"`
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Tags        []string `json:"tags"`
	Author      string   `json:"author,omitempty"`
}
