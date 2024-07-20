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
	UpdatePost(ctx context.Context, id string, post Post) error
	DeletePost(ctx context.Context, id string) error
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
	posts, err := p.repo.GetPosts(ctx)
	if err != nil {
		return nil, err
	}

	return posts, nil
}

func (p *postUsecase) GetPost(ctx context.Context, id string) (Post, error) {
	post, err := p.repo.GetPost(ctx, id)
	if err != nil {
		return Post{}, err
	}

	return post, nil
}

func (p *postUsecase) UpdatePost(ctx context.Context, id string, arg UpdatePost) error {
	currentPost, err := p.repo.GetPost(ctx, id)
	if err != nil {
		return err
	}

	updatePost := Post{
		ID:        currentPost.ID,
		Title:     arg.Title,
		Content:   arg.Content,
		CreatedAt: currentPost.CreatedAt,
	}

	if err := p.repo.UpdatePost(ctx, id, updatePost); err != nil {
		return err
	}

	return nil
}

func (p *postUsecase) DeletePost(ctx context.Context, id string) error {
	if err := p.repo.DeletePost(ctx, id); err != nil {
		return err
	}

	return nil
}
