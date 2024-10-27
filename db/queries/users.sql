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

-- name: UpdateUser :one
UPDATE users
SET
  password_hash = COALESCE(sqlc.narg(password_hash), password_hash),
  username = COALESCE(sqlc.narg(username), username),
  email = COALESCE(sqlc.narg(email), email),
  is_email_verified = COALESCE(sqlc.narg(is_email_verified), is_email_verified)
WHERE
  email = sqlc.arg(email)
RETURNING *;
