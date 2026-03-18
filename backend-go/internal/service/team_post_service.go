package service

import (
	teamseekpost "backend/internal/dto/team_seek_post"
	"backend/internal/models"
	"backend/internal/repository"
	"context"

	"gorm.io/gorm"
)

type TeamSeekPostService struct {
	db *gorm.DB
	repository repository.TeamSeekRepository
}

func NewTeamSeekPostService(db * gorm.DB, repository repository.TeamSeekRepository) *TeamSeekPostService {
	return &TeamSeekPostService{
		db: db,
		repository: repository,
	}
}


func (s *TeamSeekPostService) Create(ctx context.Context, dto *teamseekpost.CreateTeamSeekPostDTO) (*models.TeamSeekPost, error) {
	
	tx := s.db.Begin()

	post := &models.TeamSeekPost{
		Title: dto.Title,
		Description: dto.Description,
		Tags: dto.Tags,
	}

	err := s.repository.Create(ctx, tx, post)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	author := &models.Author{
		ID: post.AuthorId,
	}
	
	
	err = s.repository.CreateAuthor(ctx, tx, author)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	
	if err := tx.Commit().Error; err != nil {
		return nil, err
	}

	return post, nil
}


func (s *TeamSeekPostService) Update(ctx context.Context, dto *teamseekpost.UpdateTeamSeekPostDTO) (*models.TeamSeekPost, error) {
	
	tx := s.db.Begin()

	post := &models.TeamSeekPost{
		Title: dto.Title,
		Description: dto.Description,
		Tags: dto.Tags,
	} 

	err := s.repository.Update(ctx, tx, post)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		return nil, err
	}

	return post, nil
}


func (s *TeamSeekPostService) GetPostById(ctx context.Context, ID string) (*models.TeamSeekPost, error) {
	return s.repository.GetPostById(ctx, ID)
}


func (s *TeamSeekPostService) GetAuthorPost(ctx context.Context, authorId string) ([]models.TeamSeekPost, error) {
	return s.repository.GetAuthorPost(ctx, authorId)
}