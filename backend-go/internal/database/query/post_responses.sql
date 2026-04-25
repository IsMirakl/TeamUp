-- name: CreatePostResponse :one 
INSERT INTO post_responses (
    post_id,
    user_id,
    message
) VALUES (
    $1, $2, $3
) RETURNING response_id, post_id, user_id, message, status, created_at, updated_at;

-- name: GetPostResponses :many
SELECT 
    r.response_id, 
    r.post_id, 
    r.user_id, 
    r.message, 
    r.status, 
    r.created_at, 
    r.updated_at,
    u.email,
    u.name,
    u.avatar
FROM post_responses r
JOIN users u ON u.user_id = r.user_id
WHERE r.post_id = $1
ORDER BY r.created_at DESC;