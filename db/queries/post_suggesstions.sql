-- name: CreatePostSuggestion :one
INSERT INTO post_suggestions (
  prompt_id, text
) VALUES (
  $1, $2
)
RETURNING *;

-- name: GetPostSuggestionByID :one
SELECT id, prompt_id, text, created_at, updated_at
FROM post_suggestions
WHERE id = $1;

-- name: ListPostSuggestionsByPromptID :many
SELECT id, prompt_id, text, created_at, updated_at
FROM post_suggestions
WHERE prompt_id = $1
ORDER BY created_at DESC
LIMIT $2 OFFSET $3;

-- name: DeletePostSuggestion :one
DELETE FROM post_suggestions
WHERE id = $1
RETURNING *;


-- name: BulkCreatePostSuggestions :many
INSERT INTO post_suggestions (prompt_id, text)
SELECT @prompt_id, unnest(@suggestions::text[])
ON CONFLICT (prompt_id, text) DO NOTHING
RETURNING id, prompt_id, text, created_at;

-- name: UpdatePostSuggestionStatus :one
UPDATE post_suggestions
SET status = $2
WHERE id = $1
RETURNING *;