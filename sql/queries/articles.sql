-- name: CreateArticle :one
INSERT INTO articles (title, url, content, published_at, polarity, subjectivity, query)
VALUES ($1, $2, $3, $4, $5, $6, $7)
RETURNING id;

-- name: ListArticlesByQuery :many
SELECT * FROM articles
WHERE query = $1
ORDER BY published_at DESC;

-- name: GetArticleByID :one
SELECT * FROM articles
WHERE id = $1;
