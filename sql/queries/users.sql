-- name: CreateUser :one
INSERT INTO users (username, email, password_hash, created_at, updated_at)
VALUES ($1, $2, $3, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP)
RETURNING id;

-- name: GetUserByEmail :one
SELECT id, username, email, password_hash, created_at, updated_at
FROM users
WHERE email = $1;

-- name: GetUserByUsername :one
SELECT id, username, email, password_hash, created_at, updated_at
FROM users
WHERE username = $1;

-- name: UpdateUser :one
UPDATE users
SET username = $2, updated_at = CURRENT_TIMESTAMP
WHERE id = $1
RETURNING id;

-- name: UpdateUserPassword :one
UPDATE users
SET password_hash = $2, updated_at = CURRENT_TIMESTAMP
WHERE id = $1
RETURNING id;

-- name: DeleteUser :exec
DELETE FROM users
WHERE id = $1;
