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