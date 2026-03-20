package repository

import (
	"backend/internal/models"
	"context"
	"errors"

	"gorm.io/gorm"
)


type TeamSeekRepository interface {
	Create(ctx context.Context, tx *gorm.DB, post *models.TeamSeekPost) error
	Update(ctx context.Context, tx *gorm.DB, post *models.TeamSeekPost) error
	GetPostById(ctx context.Context, ID string) (*models.TeamSeekPost, error)
	GetAuthorPost(ctx context.Context, authorId string) ([]models.TeamSeekPost, error)
}


type teamSeekPostRepository struct {
	db *gorm.DB
}

func NewTeamSeekPostRepository(db *gorm.DB) TeamSeekRepository {
	return &teamSeekPostRepository{db: db}
}

func (r *teamSeekPostRepository) Create(ctx context.Context, tx *gorm.DB, post *models.TeamSeekPost) (error) {
	
	if tx == nil {
		tx = r.db
	}

	return tx.WithContext(ctx).Create(post).Error
}

func (r *teamSeekPostRepository) Update(ctx context.Context, tx *gorm.DB, post *models.TeamSeekPost) (error) {
	
	if tx == nil {
		tx = r.db
	}

	return tx.WithContext(ctx).Select("*").Updates(post).Error
}


func (r *teamSeekPostRepository) GetPostById(ctx context.Context, ID string) (*models.TeamSeekPost, error) {
	
	var post models.TeamSeekPost
	
	err := r.db.WithContext(ctx).Where("id = ?", ID).First(&post).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	if err != nil {
		return nil, err
	}

	return &post, nil
}


func (r *teamSeekPostRepository) GetAuthorPost(ctx context.Context, authorId string) ([]models.TeamSeekPost, error) {
	
	var posts []models.TeamSeekPost
	
	err := r.db.WithContext(ctx).
	Where("author_id = ?", authorId).
	Preload("Author").
	Preload("Author.User").
	Find(&posts).Error

	if err != nil {
		return nil, err
	}

	return posts, nil
}