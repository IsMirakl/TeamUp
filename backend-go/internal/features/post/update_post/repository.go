package updatepost

import (
	"backend/internal/features/post/model"
	"context"

	"gorm.io/gorm"
)

type Repository interface {
	Update(ctx context.Context, tx *gorm.DB, post *model.Post) error
}

type postRepository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &postRepository{db: db}
}

func (r *postRepository) Update(ctx context.Context, tx *gorm.DB, post *model.Post) error {
	if tx == nil {
		tx = r.db
	}

	return tx.WithContext(ctx).Model(&model.Post{}).Where("id = ?", post.ID).Updates(post).Error
}
