-- name: CreatePostSuggestion :one
INSERT INTO postsuggestions (
  prompt_id, suggestion_text
) VALUES (
  $1, $2
)
RETURNING *;

-- name: GetPostSuggestionByID :one
SELECT id, prompt_id, suggestion_text, created_at, updated_at
FROM postsuggestions
WHERE id = $1;

-- name: ListPostSuggestionsByPromptID :many
SELECT id, prompt_id, suggestion_text, created_at, updated_at
FROM postsuggestions
WHERE prompt_id = $1
ORDER BY created_at DESC
LIMIT $2 OFFSET $3;

-- name: UpdatePostSuggestion :one
UPDATE postsuggestions
SET suggestion_text = $2,
    updated_at = NOW()
WHERE id = $1
RETURNING *;

-- name: DeletePostSuggestion :one
DELETE FROM postsuggestions
WHERE id = $1
RETURNING *;
