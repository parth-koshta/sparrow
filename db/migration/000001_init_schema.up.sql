-- Ensure the pg_trgm extension is enabled
CREATE EXTENSION IF NOT EXISTS pg_trgm;

-- Create Users table
CREATE TABLE IF NOT EXISTS Users (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    username VARCHAR(255) UNIQUE,
    email VARCHAR(255) NOT NULL UNIQUE,
    password_hash VARCHAR(255),
    is_email_verified BOOLEAN NOT NULL DEFAULT FALSE,  -- New column added
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);

-- Create SocialAccounts table
CREATE TABLE IF NOT EXISTS SocialAccounts (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID NOT NULL REFERENCES Users(id) ON DELETE CASCADE,
    platform VARCHAR(255) NOT NULL,
    account_name VARCHAR(255) NOT NULL,
    account_email VARCHAR(255) NOT NULL,
    access_token TEXT NOT NULL,
    id_token TEXT NOT NULL,
    token_expires_at TIMESTAMP NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);

-- Create Prompts table
CREATE TABLE IF NOT EXISTS Prompts (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID NOT NULL REFERENCES Users(id) ON DELETE CASCADE,
    text TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);

-- Create PostSuggestions table
CREATE TABLE IF NOT EXISTS PostSuggestions (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    prompt_id UUID NOT NULL REFERENCES Prompts(id) ON DELETE CASCADE,
    text TEXT NOT NULL,
    status VARCHAR(50) NOT NULL DEFAULT 'pending',
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    CONSTRAINT unique_prompt_suggestion UNIQUE (prompt_id, text)
);

-- Create Posts table
CREATE TABLE IF NOT EXISTS Posts (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID NOT NULL REFERENCES Users(id) ON DELETE CASCADE,
    suggestion_id UUID REFERENCES PostSuggestions(id) ON DELETE SET NULL,
    text TEXT NOT NULL,
    status VARCHAR(50) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),
    CONSTRAINT unique_post_suggestion_id UNIQUE (suggestion_id)
);

-- Create PostSchedules table
CREATE TABLE IF NOT EXISTS PostSchedules (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID NOT NULL REFERENCES Users(id) ON DELETE CASCADE,
    post_id UUID NOT NULL REFERENCES Posts(id) ON DELETE CASCADE,
    scheduled_time TIMESTAMP NOT NULL,
    executed_time TIMESTAMP,
    status VARCHAR(50) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);

-- Create VerifyEmails table
CREATE TABLE IF NOT EXISTS VerifyEmails (
    id BIGSERIAL PRIMARY KEY,
    email VARCHAR(255) NOT NULL REFERENCES Users(email) ON DELETE CASCADE,
    secret_code VARCHAR(255) NOT NULL,
    is_used BOOLEAN NOT NULL DEFAULT FALSE,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    expired_at TIMESTAMP NOT NULL DEFAULT (NOW() + INTERVAL '15 minutes')
);

-- Create indexes for Users table
CREATE INDEX IF NOT EXISTS idx_users_email ON Users (email);
CREATE INDEX IF NOT EXISTS idx_users_username ON Users (username);

-- Create indexes for SocialAccounts table
CREATE INDEX IF NOT EXISTS idx_socialaccounts_user_id ON SocialAccounts (user_id);
CREATE INDEX IF NOT EXISTS idx_socialaccounts_platform ON SocialAccounts (platform);

-- Create indexes for Prompts table
CREATE INDEX IF NOT EXISTS idx_prompts_user_id ON Prompts (user_id);
CREATE INDEX IF NOT EXISTS idx_prompts_text ON Prompts USING GIN (text gin_trgm_ops);
CREATE INDEX IF NOT EXISTS idx_prompts_user_id_text ON Prompts (user_id, text);

-- Create indexes for PostSuggestions table
CREATE INDEX IF NOT EXISTS idx_postsuggestions_prompt_id ON PostSuggestions (prompt_id);
CREATE INDEX IF NOT EXISTS idx_postsuggestions_text ON PostSuggestions USING GIN (text gin_trgm_ops);

-- Create indexes for Posts table
CREATE INDEX IF NOT EXISTS idx_posts_user_id ON Posts (user_id);
CREATE INDEX IF NOT EXISTS idx_posts_suggestion_id ON Posts (suggestion_id);
CREATE INDEX IF NOT EXISTS idx_posts_status ON Posts (status);

-- Create indexes for PostSchedules table
CREATE INDEX IF NOT EXISTS idx_postschedules_user_id ON PostSchedules (user_id);
CREATE INDEX IF NOT EXISTS idx_postschedules_post_id ON PostSchedules (post_id);
CREATE INDEX IF NOT EXISTS idx_postschedules_scheduled_time ON PostSchedules (scheduled_time);
CREATE INDEX IF NOT EXISTS idx_postschedules_status ON PostSchedules (status);

-- Create indexes for VerifyEmails table
CREATE INDEX IF NOT EXISTS idx_verifyemails_email ON VerifyEmails (email);
