-- Drop indexes for VerifyEmails table
DROP INDEX IF EXISTS idx_verifyemails_email;

-- Drop VerifyEmails table
DROP TABLE IF EXISTS VerifyEmails;

-- Drop indexes for PostSchedules table
DROP INDEX IF EXISTS idx_postschedules_status;
DROP INDEX IF EXISTS idx_postschedules_scheduled_time;
DROP INDEX IF EXISTS idx_postschedules_post_id;
DROP INDEX IF EXISTS idx_postschedules_user_id;

-- Drop PostSchedules table
DROP TABLE IF EXISTS PostSchedules;

-- Drop indexes for Posts table
DROP INDEX IF EXISTS idx_posts_status;
DROP INDEX IF EXISTS idx_posts_suggestion_id;
DROP INDEX IF EXISTS idx_posts_user_id;

-- Drop Posts table
DROP TABLE IF EXISTS Posts;

-- Drop indexes for PostSuggestions table
DROP INDEX IF EXISTS idx_postsuggestions_text;
DROP INDEX IF EXISTS idx_postsuggestions_prompt_id;

-- Drop PostSuggestions table
DROP TABLE IF EXISTS PostSuggestions;

-- Drop indexes for Prompts table
DROP INDEX IF EXISTS idx_prompts_user_id_text;
DROP INDEX IF EXISTS idx_prompts_text;
DROP INDEX IF EXISTS idx_prompts_user_id;

-- Drop Prompts table
DROP TABLE IF EXISTS Prompts;

-- Drop indexes for SocialAccounts table
DROP INDEX IF EXISTS idx_socialaccounts_platform;
DROP INDEX IF EXISTS idx_socialaccounts_user_id;

-- Drop SocialAccounts table
DROP TABLE IF EXISTS SocialAccounts;

-- Drop indexes for Users table
DROP INDEX IF EXISTS idx_users_username;
DROP INDEX IF EXISTS idx_users_email;

-- Drop Users table
DROP TABLE IF EXISTS Users;

-- Drop pg_trgm extension
DROP EXTENSION IF EXISTS pg_trgm;
