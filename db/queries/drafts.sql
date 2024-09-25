-- name: CreateDraft :one
INSERT INTO drafts (
  user_id, suggestion_id, draft_text
) VALUES (
  $1, $2, $3
)
RETURNING *;

-- name: GetDraftByID :one
SELECT id, user_id, suggestion_id, draft_text, created_at, updated_at
FROM drafts
WHERE id = $1;

-- name: ListDraftsByUserID :many
SELECT id, user_id, suggestion_id, draft_text, created_at, updated_at
FROM drafts
WHERE user_id = $1
ORDER BY created_at DESC
LIMIT $2 OFFSET $3;

-- name: UpdateDraft :one
UPDATE drafts
SET draft_text = $2,
    updated_at = NOW()
WHERE id = $1
RETURNING *;

-- name: DeleteDraft :one
DELETE FROM drafts
WHERE id = $1
RETURNING *;
