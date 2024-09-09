CREATE EXTENSION IF NOT EXISTS pg_trgm;

-- Create Users table
CREATE TABLE Users (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    username VARCHAR(255) UNIQUE,
    email VARCHAR(255) NOT NULL UNIQUE,
    password_hash VARCHAR(255),
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);

-- Create SocialAccounts table
CREATE TABLE SocialAccounts (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID NOT NULL REFERENCES Users(id) ON DELETE CASCADE,
    platform VARCHAR(255) NOT NULL,
    account_name VARCHAR(255) NOT NULL,
    access_token VARCHAR(255) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);

-- Create Prompts table
CREATE TABLE Prompts (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID NOT NULL REFERENCES Users(id) ON DELETE CASCADE,
    prompt_text TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);

-- Create PostSuggestions table
CREATE TABLE PostSuggestions (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    prompt_id UUID NOT NULL REFERENCES Prompts(id) ON DELETE CASCADE,
    suggestion_text TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);

-- Create Drafts table
CREATE TABLE Drafts (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID NOT NULL REFERENCES Users(id) ON DELETE CASCADE,
    suggestion_id UUID NOT NULL REFERENCES PostSuggestions(id) ON DELETE CASCADE,
    draft_text TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);

-- Create ScheduledPosts table
CREATE TABLE ScheduledPosts (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID NOT NULL REFERENCES Users(id) ON DELETE CASCADE,
    draft_id UUID NOT NULL REFERENCES Drafts(id) ON DELETE CASCADE,
    scheduled_time TIMESTAMP NOT NULL,
    status VARCHAR(50) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);

-- Create indexes
CREATE INDEX idx_users_email ON Users (email);

CREATE INDEX idx_socialaccounts_user_id ON SocialAccounts (user_id);
CREATE INDEX idx_socialaccounts_platform ON SocialAccounts (platform);

CREATE INDEX idx_prompts_user_id ON Prompts (user_id);
CREATE INDEX idx_prompts_prompt_text ON Prompts USING GIN (prompt_text gin_trgm_ops);

CREATE INDEX idx_postsuggestions_prompt_id ON PostSuggestions (prompt_id);
CREATE INDEX idx_postsuggestions_suggestion_text ON PostSuggestions USING GIN (suggestion_text gin_trgm_ops);

CREATE INDEX idx_drafts_user_id ON Drafts (user_id);
CREATE INDEX idx_drafts_suggestion_id ON Drafts (suggestion_id);

CREATE INDEX idx_scheduledposts_user_id ON ScheduledPosts (user_id);
CREATE INDEX idx_scheduledposts_draft_id ON ScheduledPosts (draft_id);
CREATE INDEX idx_scheduledposts_scheduled_time ON ScheduledPosts (scheduled_time);
CREATE INDEX idx_scheduledposts_status ON ScheduledPosts (status);
