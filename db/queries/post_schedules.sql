-- name: CreatePostSchedule :one
INSERT INTO post_schedules (
  user_id, post_id, scheduled_time, status
) VALUES (
  $1, $2, $3, $4
)
RETURNING *;

-- name: GetPostScheduleByID :one
SELECT id, user_id, post_id, scheduled_time, status, created_at, updated_at
FROM post_schedules
WHERE id = $1;

-- name: ListPostSchedulesByUserID :many
SELECT id, user_id, post_id, scheduled_time, status, created_at, updated_at
FROM post_schedules
WHERE user_id = $1
ORDER BY scheduled_time DESC
LIMIT $2 OFFSET $3;

-- name: UpdatePostSchedule :one
UPDATE post_schedules
SET scheduled_time = $2,
    status = $3,
    updated_at = NOW()
WHERE id = $1
RETURNING *;

-- name: DeletePostSchedule :one
DELETE FROM post_schedules
WHERE id = $1
RETURNING *;
