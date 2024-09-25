-- name: CreateUser :one
INSERT INTO users (
  email, password_hash
) VALUES (
  $1, $2
)
RETURNING *;

-- name: GetUserByID :one
SELECT id, username, email, created_at, updated_at 
FROM users 
WHERE id = $1;

-- name: GetUserByEmail :one
SELECT id, username, email, password_hash, created_at, updated_at 
FROM users 
WHERE email = $1;

-- name: ListUsers :many
SELECT id, username, email, password_hash, created_at, updated_at
FROM users
ORDER BY created_at DESC
LIMIT $1 OFFSET $2;