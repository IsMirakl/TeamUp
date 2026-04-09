package getauthorpost

import (
	database "backend/internal/database/sqlc"
	"context"

	"github.com/sirupsen/logrus"
)

type Service struct {
	repository *Repository
	log *logrus.Logger
}

func NewService(repository *Repository, log *logrus.Logger) *Service {
	return &Service{repository: repository, log: log}
}

func (s *Service) GetAuthorPost(ctx context.Context, authorId string) (database.Post, error) {
	s.log.WithField("author_id", authorId).Info("GetByAuthor called");

	author, err := s.repository.GetAuthorPost(ctx, authorId)
	if err != nil {
		s.log.WithField("author_id", authorId).WithError(err).Error("Failed to get author post")
		return database.Post{}, err
	}
	
	return author, nil
}
