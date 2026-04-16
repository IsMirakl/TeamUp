-- name: CreateSession :one
INSERT INTO sessions (
    user_id,
    refresh_token,
    user_agent,
    client_ip,
    is_blocked,
    expires_at
) VALUES (
    $1, $2, $3, $4, $5, $6
)
RETURNING id, user_id, refresh_token, expires_at, created_at, revoked_at, user_agent, client_ip, is_blocked;

-- name: GetSessionByRefreshToken :one
SELECT id, user_id, refresh_token, expires_at, created_at, revoked_at, user_agent, client_ip, is_blocked
FROM sessions
WHERE refresh_token = $1 AND revoked_at IS NULL;

-- name: UpdateSessionRefreshToken :one
UPDATE sessions
SET refresh_token = $2,
    expires_at = $3
WHERE id = $1 AND revoked_at IS NULL
RETURNING id, user_id, refresh_token, expires_at, created_at, revoked_at, user_agent, client_ip, is_blocked;
