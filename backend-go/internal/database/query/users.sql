-- name: CreateUser :one
INSERT INTO users (
    user_id,
    email,
    email_verifed,
    name,
    avatar,
    role,
    subscription_plan,
    created_at,
    updated_at
    ) VALUES (
    $1, $2, false, $3, $4, $5, $6, NOW(), NOW()
    )
    RETURNING *;

-- name: GetUserByID :one
SELECT * FROM users
WHERE user_id = $1 AND deleted_at IS NULL 
LIMIT 1;


-- name: GetUserByEmail :one
SELECT * FROM users
WHERE email = $1 AND deleted_at IS NULL
LIMIT 1;

-- name: GetUserWithPasswordByEmail :one
SELECT 
    u.user_id,
    u.email,
    u.name,
    u.avatar,
    u.role,
    u.subscription_plan,
    a.password_hash
FROM users u
JOIN accounts a ON a.user_id = u.user_id
WHERE u.email = $1
LIMIT 1;