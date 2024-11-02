-- Ensure the pg_trgm extension is enabled
CREATE EXTENSION IF NOT EXISTS pg_trgm;

-- Create users table
CREATE TABLE IF NOT EXISTS users (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    username VARCHAR(255) UNIQUE,
    email VARCHAR(255) NOT NULL UNIQUE,
    password_hash VARCHAR(255),
    is_email_verified BOOLEAN NOT NULL DEFAULT FALSE,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);

-- Create social_accounts table
CREATE TABLE IF NOT EXISTS social_accounts (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    platform VARCHAR(255) NOT NULL,
    account_name VARCHAR(255) NOT NULL,
    account_email VARCHAR(255) NOT NULL,
    access_token TEXT NOT NULL,
    id_token TEXT NOT NULL,
    token_expires_at TIMESTAMP NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);

-- Create prompts table
CREATE TABLE IF NOT EXISTS prompts (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    text TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);

-- Create post_suggestions table
CREATE TABLE IF NOT EXISTS post_suggestions (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    prompt_id UUID NOT NULL REFERENCES prompts(id) ON DELETE CASCADE,
    text TEXT NOT NULL,
    status VARCHAR(50) NOT NULL DEFAULT 'pending',
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    CONSTRAINT unique_prompt_suggestion UNIQUE (prompt_id, text)
);

-- Create posts table
CREATE TABLE IF NOT EXISTS posts (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    suggestion_id UUID REFERENCES post_suggestions(id) ON DELETE SET NULL,
    text TEXT NOT NULL,
    status VARCHAR(50) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    CONSTRAINT unique_post_suggestion_id UNIQUE (suggestion_id)
);

-- Create post_schedules table
CREATE TABLE IF NOT EXISTS post_schedules (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    post_id UUID NOT NULL REFERENCES posts(id) ON DELETE CASCADE,
    scheduled_time TIMESTAMP NOT NULL,
    executed_time TIMESTAMP,
    status VARCHAR(50) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);

-- Create verify_emails table
CREATE TABLE IF NOT EXISTS verify_emails (
    id BIGSERIAL PRIMARY KEY,
    email VARCHAR(255) NOT NULL REFERENCES users(email) ON DELETE CASCADE,
    secret_code VARCHAR(255) NOT NULL,
    is_used BOOLEAN NOT NULL DEFAULT FALSE,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    expired_at TIMESTAMP NOT NULL DEFAULT (NOW() + INTERVAL '15 minutes')
);

-- Create indexes for users table
CREATE INDEX IF NOT EXISTS idx_users_email ON users (email);
CREATE INDEX IF NOT EXISTS idx_users_username ON users (username);

-- Create indexes for social_accounts table
CREATE INDEX IF NOT EXISTS idx_social_accounts_user_id ON social_accounts (user_id);
CREATE INDEX IF NOT EXISTS idx_social_accounts_platform ON social_accounts (platform);

-- Create indexes for prompts table
CREATE INDEX IF NOT EXISTS idx_prompts_user_id ON prompts (user_id);
CREATE INDEX IF NOT EXISTS idx_prompts_text ON prompts USING GIN (text gin_trgm_ops);
CREATE INDEX IF NOT EXISTS idx_prompts_user_id_text ON prompts (user_id, text);

-- Create indexes for post_suggestions table
CREATE INDEX IF NOT EXISTS idx_post_suggestions_prompt_id ON post_suggestions (prompt_id);
CREATE INDEX IF NOT EXISTS idx_post_suggestions_text ON post_suggestions USING GIN (text gin_trgm_ops);

-- Create indexes for posts table
CREATE INDEX IF NOT EXISTS idx_posts_user_id ON posts (user_id);
CREATE INDEX IF NOT EXISTS idx_posts_suggestion_id ON posts (suggestion_id);
CREATE INDEX IF NOT EXISTS idx_posts_status ON posts (status);

-- Create indexes for post_schedules table
CREATE INDEX IF NOT EXISTS idx_post_schedules_user_id ON post_schedules (user_id);
CREATE INDEX IF NOT EXISTS idx_post_schedules_post_id ON post_schedules (post_id);
CREATE INDEX IF NOT EXISTS idx_post_schedules_scheduled_time ON post_schedules (scheduled_time);
CREATE INDEX IF NOT EXISTS idx_post_schedules_status ON post_schedules (status);

-- Create indexes for verify_emails table
CREATE INDEX IF NOT EXISTS idx_verify_emails_email ON verify_emails (email);
