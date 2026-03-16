package teamseekpost

type CreateTeamSeekPostDTO struct {
	Title string `json:"title" validate:"required,max=50"`
	Description string `json:"description" validate:"required,min=100,max=750"`
	Tags []string `json:"tags"`
}
