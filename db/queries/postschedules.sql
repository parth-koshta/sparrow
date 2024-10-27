-- name: CreatePostSchedule :one
INSERT INTO postschedules (
  user_id, post_id, scheduled_time, status
) VALUES (
  $1, $2, $3, $4
)
RETURNING *;

-- name: GetPostScheduleByID :one
SELECT id, user_id, post_id, scheduled_time, status, created_at, updated_at
FROM postschedules
WHERE id = $1;

-- name: ListPostSchedulesByUserID :many
SELECT id, user_id, post_id, scheduled_time, status, created_at, updated_at
FROM postschedules
WHERE user_id = $1
ORDER BY scheduled_time DESC
LIMIT $2 OFFSET $3;

-- name: UpdatePostSchedule :one
UPDATE postschedules
SET scheduled_time = $2,
    status = $3,
    updated_at = NOW()
WHERE id = $1
RETURNING *;

-- name: DeletePostSchedule :one
DELETE FROM postschedules
WHERE id = $1
RETURNING *;
