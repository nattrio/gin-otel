package post

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type postRepo struct {
	db *pgxpool.Pool
}

func NewPostRepo(db *pgxpool.Pool) *postRepo {
	return &postRepo{
		db: db,
	}
}

func (p *postRepo) CreatePost(ctx context.Context, post Post) error {
	query := `INSERT INTO posts (id, title, content, created_at)
		VALUES ($1, $2, $3, $4)
	`

	_, err := p.db.Exec(ctx, query,
		post.ID,
		post.Title,
		post.Content,
		post.CreatedAt,
	)
	if err != nil {
		return err
	}

	return nil
}

func (p *postRepo) GetPost(ctx context.Context, id string) (Post, error) {
	return Post{}, nil
}

func (p *postRepo) GetPosts(ctx context.Context) ([]Post, error) {
	return []Post{}, nil
}
