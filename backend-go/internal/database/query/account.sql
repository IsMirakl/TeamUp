-- name: CreateAccount :exec
INSERT INTO accounts (
    user_id,
    password_hash,
    refresh_token,
    provider
) VALUES (
    $1, $2, $3, $4
);