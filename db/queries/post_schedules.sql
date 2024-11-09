-- name: SchedulePost :one
INSERT INTO post_schedules (
  user_id, post_id, social_account_id, scheduled_time, status
) VALUES (
  $1, $2, $3, $4, $5
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

-- name: DeleteSchedule :one
DELETE FROM post_schedules
WHERE id = $1
RETURNING *;

-- name: GetScheduledPostsWithinTimeframe :many
SELECT p.id, 
       p.user_id, 
       p.suggestion_id, 
       p.text, 
       p.status, 
       p.created_at, 
       p.updated_at, 
       ps.scheduled_time,
       ps.id AS schedule_id
FROM posts p
JOIN post_schedules ps ON p.id = ps.post_id
WHERE p.status = 'scheduled'
  AND ps.scheduled_time BETWEEN NOW() - (sqlc.arg(hours_from)::int * INTERVAL '1 hour') 
  AND NOW() + (sqlc.arg(hours_till)::int * INTERVAL '1 hour')
ORDER BY ps.scheduled_time ASC;

-- name: UpdatePostScheduleExectued :one
UPDATE post_schedules
SET executed_time = NOW(),
    status = 'executed',
    updated_at = NOW()
WHERE id = $1
RETURNING *;