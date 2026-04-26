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
SELECT
    p.id,
    p.title,
    p.description,
    p.tags,
    p.author_id,
    COALESCE(NULLIF(u.name, ''), ('user-' || LEFT(p.author_id::text, 8)))::text AS author_name
FROM posts p
JOIN users u ON u.user_id = p.author_id
WHERE p.id = $1 AND p.deleted_at IS NULL
LIMIT 1;

-- name: UpdatePost :one
UPDATE posts
SET
    title = COALESCE($1, title),
    description = COALESCE($2, description),
    tags = COALESCE($3, tags),
    updated_at = NOW()
WHERE id = $4
RETURNING id, title, description, tags;


-- name: ListPosts :many
SELECT
    p.id,
    p.title,
    p.description,
    p.tags,
    p.author_id,
    COALESCE(NULLIF(u.name, ''), ('user-' || LEFT(p.author_id::text, 8)))::text AS author_name
FROM posts p
JOIN users u ON u.user_id = p.author_id
WHERE p.deleted_at IS NULL
ORDER BY p.created_at DESC
LIMIT $1 OFFSET $2;
