// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: articles.sql

package database

import (
	"context"
	"database/sql"
)

const createArticle = `-- name: CreateArticle :one
INSERT INTO articles (title, url, content, published_at, polarity, subjectivity, query)
VALUES ($1, $2, $3, $4, $5, $6, $7)
RETURNING id
`

type CreateArticleParams struct {
	Title        string
	Url          string
	Content      sql.NullString
	PublishedAt  sql.NullTime
	Polarity     sql.NullFloat64
	Subjectivity sql.NullFloat64
	Query        sql.NullString
}

func (q *Queries) CreateArticle(ctx context.Context, arg CreateArticleParams) (int32, error) {
	row := q.db.QueryRowContext(ctx, createArticle,
		arg.Title,
		arg.Url,
		arg.Content,
		arg.PublishedAt,
		arg.Polarity,
		arg.Subjectivity,
		arg.Query,
	)
	var id int32
	err := row.Scan(&id)
	return id, err
}

const getArticleByID = `-- name: GetArticleByID :one
SELECT id, title, url, content, published_at, polarity, subjectivity, query, created_at FROM articles
WHERE id = $1
`

func (q *Queries) GetArticleByID(ctx context.Context, id int32) (Article, error) {
	row := q.db.QueryRowContext(ctx, getArticleByID, id)
	var i Article
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Url,
		&i.Content,
		&i.PublishedAt,
		&i.Polarity,
		&i.Subjectivity,
		&i.Query,
		&i.CreatedAt,
	)
	return i, err
}

const listArticlesByQuery = `-- name: ListArticlesByQuery :many
SELECT id, title, url, content, published_at, polarity, subjectivity, query, created_at FROM articles
WHERE query = $1
ORDER BY published_at DESC
`

func (q *Queries) ListArticlesByQuery(ctx context.Context, query sql.NullString) ([]Article, error) {
	rows, err := q.db.QueryContext(ctx, listArticlesByQuery, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Article
	for rows.Next() {
		var i Article
		if err := rows.Scan(
			&i.ID,
			&i.Title,
			&i.Url,
			&i.Content,
			&i.PublishedAt,
			&i.Polarity,
			&i.Subjectivity,
			&i.Query,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}