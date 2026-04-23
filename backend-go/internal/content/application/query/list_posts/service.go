package listposts

import (
	"context"

	database "backend/internal/database/sqlc"

	"github.com/sirupsen/logrus"
)

type Service struct {
	repository Repository
	log        *logrus.Logger
}

type Repository interface {
	ListPosts(ctx context.Context, limit, offset int32) ([]database.ListPostsRow, error)
}

func NewService(repository Repository, log *logrus.Logger) *Service {
	return &Service{repository: repository, log: log}
}

func (s *Service) ListPosts(ctx context.Context, limit, offset int32) ([]database.ListPostsRow, error) {
	s.log.WithFields(logrus.Fields{"limit": limit, "offset": offset}).Info("ListPosts called")

	posts, err := s.repository.ListPosts(ctx, limit, offset)
	if err != nil {
		s.log.WithError(err).Error("Failed to list posts")
		return nil, err
	}

	return posts, nil
}
