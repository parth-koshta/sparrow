-- Drop indexes

DROP INDEX IF EXISTS idx_verifyemails_email;

DROP INDEX IF EXISTS idx_scheduledposts_status;
DROP INDEX IF EXISTS idx_scheduledposts_scheduled_time;
DROP INDEX IF EXISTS idx_scheduledposts_draft_id;
DROP INDEX IF EXISTS idx_scheduledposts_user_id;

DROP INDEX IF EXISTS idx_posts_suggestion_id;
DROP INDEX IF EXISTS idx_posts_user_id;

DROP INDEX IF EXISTS idx_postsuggestions_text;
DROP INDEX IF EXISTS idx_postsuggestions_prompt_id;

DROP INDEX IF EXISTS idx_prompts_text;
DROP INDEX IF EXISTS idx_prompts_user_id;

DROP INDEX IF EXISTS idx_socialaccounts_platform;
DROP INDEX IF EXISTS idx_socialaccounts_user_id;

DROP INDEX IF EXISTS idx_users_email;
DROP INDEX IF EXISTS idx_users_username;

-- Drop tables
DROP TABLE IF EXISTS ScheduledPosts;
DROP TABLE IF EXISTS Posts;
DROP TABLE IF EXISTS PostSuggestions;
DROP TABLE IF EXISTS Prompts;
DROP TABLE IF EXISTS SocialAccounts;
DROP TABLE IF EXISTS Users;
DROP TABLE IF EXISTS VerifyEmails;
