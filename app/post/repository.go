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

	posts := []Post{}
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

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return posts, nil
}

func (p *postRepo) UpdatePost(ctx context.Context, id string, post Post) error {
	query := `UPDATE posts
		SET title = $1, content = $2
		WHERE id = $3
	`

	_, err := p.db.Exec(ctx, query,
		post.Title,
		post.Content,
		id,
	)
	if err != nil {
		return err
	}

	return nil
}

func (p *postRepo) DeletePost(ctx context.Context, id string) error {
	query := `DELETE FROM posts
		WHERE id = $1
	`

	_, err := p.db.Exec(ctx, query, id)
	if err != nil {
		return err
	}

	return nil
}
