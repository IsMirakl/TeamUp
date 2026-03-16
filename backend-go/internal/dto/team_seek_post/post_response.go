package teamseekpost

type RepsonseTeamSeekPostDTO struct {
	ID string `json:"id"`
	Title string `json:"title"`
	Description string `json:"description"`
	Tags []string `json:"tags"`
}