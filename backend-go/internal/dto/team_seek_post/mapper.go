package teamseekpost

import "backend/internal/models"


func ToTeamSeekPostResponse(post * models.TeamSeekPost) *RepsonseTeamSeekPostDTO {
	return &RepsonseTeamSeekPostDTO{
		ID: post.ID,
		Title: post.Title,
		Description: post.Description,
		Tags: post.Tags,
	}
}