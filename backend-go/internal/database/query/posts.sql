-- name: CreatePost :one
INSERT INTO posts (
    id,
    title,
    description,
    tags,
    author_id,
    created_at,
    updated_at
) VALUES (
    $1, $2, $3, $4, $5, NOW(), NOW()
) RETURNING *;


-- name: GetAuthorPost :one
SELECT * FROM posts
WHERE author_id = $1 AND deleted_at IS NULL
LIMIT 1;

-- name: GetPostById :one
SELECT * FROM posts
WHERE id = $1 AND deleted_at IS NULL
LIMIT 1;

-- name: UpdatePost :one
UPDATE posts
SET
    title = COALESCE($1, title),
    description = COALESCE($2, description),
    tags = COALESCE($3, tags),
    updated_at = NOW()
WHERE id = $4
RETURNING *;

