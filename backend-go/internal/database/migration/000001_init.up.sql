create type roles as enum ('user', 'admin', 'team_lead');
create type subscription_plans as enum ('Free', 'Pro', 'Enterprise');

CREATE TABLE users (
    user_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMPTZ,

    email TEXT NOT NULL UNIQUE,
    email_verified BOOLEAN NOT NULL DEFAULT FALSE,

    name VARCHAR(25) NOT NULL,
    avatar TEXT,

    role roles NOT NULL DEFAULT 'user',
    subscription_plan subscription_plans NOT NULL DEFAULT 'Free'
);

create type providers as enum ('local', 'github', 'google');

CREATE TABLE accounts (
    user_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    password_hash VARCHAR(255) NOT NULL,
    refresh_token  TEXT,
    provider providers NOT NULL DEFAULT 'local',

    CONSTRAINT fk_accounts_user
        FOREIGN KEY (user_id)
        REFERENCES users(user_id)
        ON DELETE CASCADE
);