-- name: CreatSession :one
INSERT INTO sessions (
    id,
    user_id,
    refresh_token,
    user_agent,
    client_ip,
    is_blocked,
    expires_at,
    created_at,
    revoked_at
) VALUES (
    $1, $2, $3, $4, $5, false, $6, NOW(), NOW()
) RETURNING *;