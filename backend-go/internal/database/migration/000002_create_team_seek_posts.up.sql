CREATE TABLE posts (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMPTZ,

    title VARCHAR(50) NOT NULL,
    description VARCHAR(750) NOT NULL,
    tags TEXT[],
    author_id TEXT NOT NULL,

    CONSTRAINT min_length_description CHECK (char_length(description) >= 150), 
    CONSTRAINT fk_post_author
        FOREIGN KEY (author_id)
        REFERENCES users(user_id)
        ON DELETE CASCADE
)