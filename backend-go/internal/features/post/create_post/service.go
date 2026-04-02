package createpost

import (
	"backend/internal/features/post/dto"
	"backend/internal/features/post/model"
	"context"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type Service struct {
	db         *gorm.DB
	repository Repository
	log *logrus.Logger
}

func NewService(db *gorm.DB, repository Repository, log *logrus.Logger) *Service {
	return &Service{
		db:         db,
		repository: repository,
		log: log,
	}
}

func (s *Service) Create(ctx context.Context, dto *dto.CreatePostDTO, userID string) (*model.Post, error) {
	
	tx := s.db.Begin()
	
	s.log.WithField("title_post", dto.Title).Info("Creating post")

	post := &model.Post{
		ID:          uuid.NewString(),
		Title:       dto.Title,
		Description: dto.Description,
		Tags:        dto.Tags,
		AuthorID:    userID,
	}

	err := s.repository.Create(ctx, tx, post)
	if err != nil {
		tx.Rollback()

		s.log.WithField("post_ID", post.ID).WithError(err).Error("Failed create post")
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		return nil, err
	}

	s.log.WithField("post_ID", post.ID).Info("Post successfully created")

	return post, nil
}
