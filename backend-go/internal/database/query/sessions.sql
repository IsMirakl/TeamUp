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
