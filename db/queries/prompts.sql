-- name: CreatePrompt :one
INSERT INTO prompts (
  user_id, prompt_text
) VALUES (
  $1, $2
)
RETURNING *;

-- name: GetPromptByID :one
SELECT id, user_id, prompt_text, created_at, updated_at
FROM prompts
WHERE id = $1;

-- name: ListPromptsByUserID :many
SELECT id, user_id, prompt_text, created_at, updated_at
FROM prompts
WHERE user_id = $1
ORDER BY created_at DESC
LIMIT $2 OFFSET $3;

-- name: UpdatePrompt :one
UPDATE prompts
SET prompt_text = $2,
    updated_at = NOW()
WHERE id = $1
RETURNING *;

-- name: DeletePrompt :one
DELETE FROM prompts
WHERE id = $1
RETURNING *;
