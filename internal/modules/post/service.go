package post

import (
	"context"
	"errors"
	db "myshelf/internal/db/sqlc/generated"
	"myshelf/pkg/logger"

	"go.uber.org/zap"
)

type PostService struct {
	Repo PostRepository
}

func (s *PostService) CreatePost(ctx context.Context, userID int32, title, content string) error {
	logger.Log.Info("Service: Creating post", zap.Int32("userID", userID))

	if title == "" || content == "" {
		err := errors.New("title and content are required")
		logger.Log.Warn("Validation failed", zap.Error(err))
		return err
	}

	err := s.Repo.CreatePost(ctx, userID, title, content)
	if err != nil {
		logger.Log.Error("Failed to create post in repo", zap.Error(err))
	}

	return err
}

func (s *PostService) GetPostByID(ctx context.Context, id int32) (db.Post, error) {
	return s.Repo.GetPostByID(ctx, id)
}

func (s *PostService) UpdatePost(ctx context.Context, id int32, content string) error {
	if content == "" {
		return errors.New("content cannot be empty")
	}
	return s.Repo.UpdatePost(ctx, id, content)
}

func (s *PostService) DeletePost(ctx context.Context, id int32) error {
	if id == 0 {
		return errors.New("invalid post ID")
	}
	return s.Repo.DeletePost(ctx, id)
}
