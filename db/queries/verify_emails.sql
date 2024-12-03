-- name: CreateVerifyEmail :one
INSERT INTO verify_emails (
    email,
    secret_code
) VALUES (
    $1, $2
) RETURNING *;

-- name: UpdateVerifyEmail :one
UPDATE verify_emails
SET
    is_used = TRUE
WHERE
    email = @email
    AND secret_code = @secret_code
    AND is_used = FALSE
    AND expired_at > now()
RETURNING *;

-- name: GetVerifyEmail :one
SELECT id, email, secret_code, expired_at, is_used, created_at
FROM verify_emails
WHERE email = @email;

-- name: InvalidateVerifyEmail :one
UPDATE verify_emails
SET
    expired_at = now() - interval '1 second'
WHERE
    id = @id
RETURNING *;