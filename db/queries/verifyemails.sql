-- name: CreateVerifyEmail :one
INSERT INTO verifyemails (
    email,
    secret_code
) VALUES (
    $1, $2
) RETURNING *;

-- name: UpdateVerifyEmail :one
UPDATE verifyemails
SET
    is_used = TRUE
WHERE
    id = @id
    AND secret_code = @secret_code
    AND is_used = FALSE
    AND expired_at > now()
RETURNING *;