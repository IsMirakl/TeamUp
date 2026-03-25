package getbyid

import (
	"backend/internal/features/post/model"
	"context"
	"errors"

	"gorm.io/gorm"
)

type Repository interface {
	GetPostById(ctx context.Context, id string) (*model.Post, error)
}

type postRepository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &postRepository{db: db}
}

func (r *postRepository) GetPostById(ctx context.Context, id string) (*model.Post, error) {
	var post model.Post

	err := r.db.WithContext(ctx).Where("id = ?", id).First(&post).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	if err != nil {
		return nil, err
	}

	return &post, nil
}
