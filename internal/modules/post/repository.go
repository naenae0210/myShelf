package post

import (
	"context"
	"errors"
	db "myshelf/internal/db/sqlc/generated"
	"myshelf/pkg/logger"

	"go.uber.org/zap"
)

type PostRepository interface {
	CreatePost(ctx context.Context, userID int32, title, content string) error
	GetPostByID(ctx context.Context, id int32) (db.Post, error)
	UpdatePost(ctx context.Context, id int32, content string) error
	DeletePost(ctx context.Context, id int32) error
}

type postRepository struct {
	queries *db.Queries
}

func NewPostRepository(db *db.Queries) PostRepository {
	return &postRepository{
		queries: db,
	}
}

func (r *postRepository) CreatePost(ctx context.Context, userID int32, title, content string) error {
	logger.Log.Info("creating post", zap.Int32("userID", userID), zap.String("title", title))

	params := db.CreatePostParams{
		UserID:  userID,
		Title:   title,
		Content: content,
	}

	err := r.queries.CreatePost(ctx, params)
	if err != nil {
		logger.Log.Error("Failed to create post", zap.Error(err))
	}

	return err
}

func (r *postRepository) GetPostByID(ctx context.Context, id int32) (db.Post, error) {

	logger.Log.Info("Fetching post", zap.Int32("id", id))

	post, err := r.queries.GetPostByID(ctx, id)
	if err != nil {
		logger.Log.Error("Post not found", zap.Int32("id", id), zap.Error(err))
		return db.Post{}, err
	}
	return post, nil
}

func (r *postRepository) UpdatePost(ctx context.Context, id int32, content string) error {
	params := db.UpdatePostParams{
		Content: content,
		ID:      id,
	}
	return r.queries.UpdatePost(ctx, params)
}

func (r *postRepository) DeletePost(ctx context.Context, id int32) error {
	if id == 0 {
		return errors.New("invalid post ID")
	}
	return r.queries.DeletePost(ctx, id)
}
