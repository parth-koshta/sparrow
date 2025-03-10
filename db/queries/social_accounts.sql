-- name: CreateSocialAccount :one
INSERT INTO social_accounts (
  user_id, platform, name, email, access_token, id_token, token_expires_at, linkedin_sub
) VALUES (
  $1, $2, $3, $4, $5, $6, $7, $8
)
RETURNING *;

-- name: GetSocialAccountByID :one
SELECT platform, user_id, name, access_token, linkedin_sub, token_expires_at, updated_at
FROM social_accounts
WHERE id = $1;

-- name: ListSocialAccountsByUserID :many
SELECT id, user_id, platform, name, email, token_expires_at, created_at, updated_at
FROM social_accounts
WHERE user_id = $1
ORDER BY created_at DESC
LIMIT $2 OFFSET $3;

-- name: UpdateSocialAccount :one
UPDATE social_accounts
SET platform = $2,
    name = $3,
    access_token = $4,
    updated_at = NOW()
WHERE id = $1
RETURNING *;

-- name: UpdateSocialAccountToken :one
UPDATE social_accounts
SET access_token = $2,
    id_token = $3,
    token_expires_at = $4,
    updated_at = NOW()
WHERE user_id = $1
RETURNING *;

-- name: DeleteSocialAccount :one
DELETE FROM social_accounts
WHERE id = $1
RETURNING *;

