-- name: CreatePrompt :one
INSERT INTO prompts (
  user_id, text
) VALUES (
  $1, $2
)
RETURNING *;

-- name: GetPromptByID :one
SELECT id, user_id, text, created_at, updated_at
FROM prompts
WHERE id = $1;

-- name: GetPromptByUserIDAndText :one
SELECT id, user_id, text, created_at, updated_at
FROM prompts
WHERE user_id = $1 AND text = $2
LIMIT 1;

-- name: ListPromptsByUserID :many
SELECT id, user_id, text, created_at, updated_at
FROM prompts
WHERE user_id = $1
ORDER BY created_at DESC
LIMIT $2 OFFSET $3;

-- name: UpdatePrompt :one
UPDATE prompts
SET text = $2,
    updated_at = NOW()
WHERE id = $1
RETURNING *;

-- name: DeletePrompt :one
DELETE FROM prompts
WHERE id = $1
RETURNING *;
