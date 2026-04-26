package dto

type ResponsePostDTO struct {
	ID          string   `json:"id"`
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Tags        []string `json:"tags"`
	AuthorID    string   `json:"author_id,omitempty"`
	Author      string   `json:"author,omitempty"`
}
