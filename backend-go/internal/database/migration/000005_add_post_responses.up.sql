create type STATUS_RESPONSES as enum ('pending', 'accepted', 'rejected');

CREATE TABLE post_responses (
    response_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMPTZ,

    post_id UUID NOT NULL REFERENCES team_seek_posts(id) ON DELETE CASCADE,
    user_id UUID NOT NULL REFERENCES users(user_id) ON DELETE CASCADE,

    message TEXT NOT NULL,

    status STATUS_RESPONSES NOT NULL DEFAULT 'pending',

    CONSTRAINT unique_user_post_response UNIQUE (post_id, user_id)
);

CREATE INDEX idx_post_responses_post_id ON post_responses(post_id);
CREATE INDEX idx_post_responses_user_id ON post_responses(user_id);