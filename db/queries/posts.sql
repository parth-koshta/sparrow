-- name: CreatePost :one
INSERT INTO posts (
  user_id, suggestion_id, text, status
) VALUES (
  $1, $2, $3, $4
)
RETURNING *;

-- name: GetPostByID :one
SELECT id, user_id, suggestion_id, text, created_at, updated_at
FROM posts
WHERE id = $1;

-- name: ListPostsByUserID :many
SELECT id, user_id, suggestion_id, text, created_at, updated_at
FROM posts
WHERE user_id = $1
ORDER BY created_at DESC
LIMIT $2 OFFSET $3;

-- name: UpdatePost :one
UPDATE posts
SET text = $2,
    updated_at = NOW()
WHERE id = $1
RETURNING *;

-- name: DeletePost :one
DELETE FROM posts
WHERE id = $1
RETURNING *;


-- name: UpdatePostStatus :one
UPDATE posts
SET status = $2,
    updated_at = NOW()
WHERE id = $1
RETURNING *;