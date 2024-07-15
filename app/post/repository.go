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
	query := `SELECT id, title, content, created_at
		FROM posts
		WHERE id = $1
	`

	var post Post
	err := p.db.QueryRow(ctx, query, id).Scan(
		&post.ID,
		&post.Title,
		&post.Content,
		&post.CreatedAt,
	)
	if err != nil {
		return Post{}, err
	}

	return post, nil
}

func (p *postRepo) GetPosts(ctx context.Context) ([]Post, error) {
	query := `SELECT id, title, content, created_at
		FROM posts
	`

	rows, err := p.db.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var posts []Post
	for rows.Next() {
		var post Post
		err := rows.Scan(
			&post.ID,
			&post.Title,
			&post.Content,
			&post.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}

	return posts, nil
}
