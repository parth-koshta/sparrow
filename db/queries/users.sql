-- name: CreateUser :one
INSERT INTO users (
  name, email, password_hash
) VALUES (
  $1, $2, $3
)
RETURNING *;

-- name: GetUserByID :one
SELECT id, name, email, created_at, updated_at, is_email_verified
FROM users 
WHERE id = $1;

-- name: GetUserByEmail :one
SELECT id, name, email, password_hash, created_at, updated_at, is_email_verified 
FROM users 
WHERE email = $1;

-- name: ListUsers :many
SELECT id, name, email, password_hash, created_at, updated_at
FROM users
ORDER BY created_at DESC
LIMIT $1 OFFSET $2;

-- name: UpdateUser :one
UPDATE users
SET
  password_hash = COALESCE(sqlc.narg(password_hash), password_hash),
  name = COALESCE(sqlc.narg(name), name),
  email = COALESCE(sqlc.narg(email), email),
  is_email_verified = COALESCE(sqlc.narg(is_email_verified), is_email_verified)
WHERE
  email = sqlc.arg(email)
RETURNING *;
