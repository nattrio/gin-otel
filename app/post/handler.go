package post

import (
	"context"

	"github.com/gin-gonic/gin"
)

type PostUsecase interface {
	CreatePost(ctx context.Context, arg CreatePost) error
	GetPost(ctx context.Context, id string) (Post, error)
	GetPosts(ctx context.Context) ([]Post, error)
}

type postHandler struct {
	usecase PostUsecase
}

func NewPostHandler(usecase PostUsecase) *postHandler {
	return &postHandler{
		usecase: usecase,
	}
}

type CreatePost struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

func (h *postHandler) CreatePost(c *gin.Context) {
	var request CreatePost
	ctx := c.Request.Context()

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(400, gin.H{"error": "invalid json"})
		return
	}

	if err := h.usecase.CreatePost(ctx, request); err != nil {
		c.JSON(500, gin.H{"error": "failed to create post"})
		return
	}

	c.JSON(201, gin.H{})
}

func (h *postHandler) GetPost(c *gin.Context) {
	id := c.Param("id")
	ctx := c.Request.Context()

	if post, err := h.usecase.GetPost(ctx, id); err != nil {
		c.JSON(500, gin.H{"error": "failed to get post"})
		return
	} else {
		c.JSON(200, post)
	}
}

func (h *postHandler) GetPosts(c *gin.Context) {
	ctx := c.Request.Context()

	if posts, err := h.usecase.GetPosts(ctx); err != nil {
		c.JSON(500, gin.H{"error": "failed to get post"})
		return
	} else {
		c.JSON(200, posts)
	}
}
