-- name: InsertRefreshToken :one
INSERT INTO refresh_tokens (id, user_id, token, expires_at, created_at, revoked)
VALUES ($1, $2, $3, $4, $5, $6) RETURNING id;

-- name: GetRefreshToken :one
SELECT id, user_id, token, expires_at, created_at, revoked
FROM refresh_tokens
WHERE token = $1;

-- name: GetRefreshTokenById :one
SELECT id, user_id, token, expires_at, created_at, revoked
FROM refresh_tokens
WHERE id = $1;

-- name: InvalidateRefreshToken :exec
UPDATE refresh_tokens
SET revoked = true
WHERE token = $1;

-- name: InvalidateRefreshTokenById :exec
UPDATE refresh_tokens
SET revoked = true
WHERE id = $1;

-- name: InvalidateAllTokensForUser :exec
UPDATE refresh_tokens
SET revoked = true
WHERE user_id = $1 AND revoked = false;
