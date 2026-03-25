package getauthorpost

import (
	"backend/internal/features/post/model"
	"context"

	"gorm.io/gorm"
)

type Repository interface {
	GetAuthorPost(ctx context.Context, authorId string) ([]model.Post, error)
}

type postRepository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &postRepository{db: db}
}

func (r *postRepository) GetAuthorPost(ctx context.Context, authorId string) ([]model.Post, error) {
	var posts []model.Post

	err := r.db.WithContext(ctx).Where("author_id = ?", authorId).Find(&posts).Error
	if err != nil {
		return nil, err
	}

	return posts, nil
}
