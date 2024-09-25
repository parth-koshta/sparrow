-- name: CreateSocialAccount :one
INSERT INTO socialaccounts (
  user_id, platform, account_name, access_token
) VALUES (
  $1, $2, $3, $4
)
RETURNING *;

-- name: GetSocialAccountByID :one
SELECT id, user_id, platform, account_name, access_token, created_at, updated_at
FROM socialaccounts
WHERE id = $1;

-- name: ListSocialAccountsByUserID :many
SELECT id, user_id, platform, account_name, access_token, created_at, updated_at
FROM socialaccounts
WHERE user_id = $1
ORDER BY created_at DESC
LIMIT $2 OFFSET $3;

-- name: UpdateSocialAccount :one
UPDATE socialaccounts
SET platform = $2,
    account_name = $3,
    access_token = $4,
    updated_at = NOW()
WHERE id = $1
RETURNING *;

-- name: DeleteSocialAccount :one
DELETE FROM socialaccounts
WHERE id = $1
RETURNING *;

