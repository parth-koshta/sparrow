-- name: CreatePostSuggestion :one
INSERT INTO postsuggestions (
  prompt_id, text
) VALUES (
  $1, $2
)
RETURNING *;

-- name: GetPostSuggestionByID :one
SELECT id, prompt_id, text, created_at, updated_at
FROM postsuggestions
WHERE id = $1;

-- name: ListPostSuggestionsByPromptID :many
SELECT id, prompt_id, text, created_at, updated_at
FROM postsuggestions
WHERE prompt_id = $1
ORDER BY created_at DESC
LIMIT $2 OFFSET $3;

-- name: UpdatePostSuggestion :one
UPDATE postsuggestions
SET text = $2,
    updated_at = NOW()
WHERE id = $1
RETURNING *;

-- name: DeletePostSuggestion :one
DELETE FROM postsuggestions
WHERE id = $1
RETURNING *;


-- name: BulkCreatePostSuggestions :many
INSERT INTO postsuggestions (prompt_id, text)
SELECT @prompt_id, unnest(@suggestions::text[])
ON CONFLICT (prompt_id, text) DO NOTHING
RETURNING id, prompt_id, text, created_at;