package post

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/uptrace/opentelemetry-go-extra/otelzap"
	"go.uber.org/zap"
)

type PostRepo interface {
	CreatePost(ctx context.Context, post Post) error
	GetPost(ctx context.Context, id string) (Post, error)
	GetPosts(ctx context.Context) ([]Post, error)
}

type postUsecase struct {
	repo PostRepo
}

func NewPostUsecase(r PostRepo) *postUsecase {
	return &postUsecase{
		repo: r,
	}
}

func (p *postUsecase) CreatePost(ctx context.Context, arg CreatePost) error {
	newPost := Post{
		ID:        uuid.New().String(),
		Title:     arg.Title,
		Content:   arg.Content,
		CreatedAt: time.Now(),
	}

	if err := p.repo.CreatePost(ctx, newPost); err != nil {
		otelzap.Ctx(ctx).Error("failed to create post", zap.Error(err))
		return err
	}

	return nil
}

func (p *postUsecase) GetPosts(ctx context.Context) ([]Post, error) {
	return p.repo.GetPosts(ctx)
}

func (p *postUsecase) GetPost(ctx context.Context, id string) (Post, error) {
	return p.repo.GetPost(ctx, id)
}
