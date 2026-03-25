package createpost

import (
	"backend/internal/features/post/model"
	"context"

	"gorm.io/gorm"
)

type Repository interface {
	Create(ctx context.Context, tx *gorm.DB, post *model.Post) error
}

type postRepository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &postRepository{db: db}
}

func (r *postRepository) Create(ctx context.Context, tx *gorm.DB, post *model.Post) error {
	if tx == nil {
		tx = r.db
	}

	return tx.WithContext(ctx).Create(post).Error
}
