-- name: CreateScheduledPost :one
INSERT INTO scheduledposts (
  user_id, draft_id, scheduled_time, status
) VALUES (
  $1, $2, $3, $4
)
RETURNING *;

-- name: GetScheduledPostByID :one
SELECT id, user_id, draft_id, scheduled_time, status, created_at, updated_at
FROM scheduledposts
WHERE id = $1;

-- name: ListScheduledPostsByUserID :many
SELECT id, user_id, draft_id, scheduled_time, status, created_at, updated_at
FROM scheduledposts
WHERE user_id = $1
ORDER BY scheduled_time DESC
LIMIT $2 OFFSET $3;

-- name: UpdateScheduledPost :one
UPDATE scheduledposts
SET scheduled_time = $2,
    status = $3,
    updated_at = NOW()
WHERE id = $1
RETURNING *;

-- name: DeleteScheduledPost :one
DELETE FROM scheduledposts
WHERE id = $1
RETURNING *;
