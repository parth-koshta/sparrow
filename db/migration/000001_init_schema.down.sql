-- Drop indexes
DROP INDEX IF EXISTS idx_scheduledposts_status;
DROP INDEX IF EXISTS idx_scheduledposts_scheduled_time;
DROP INDEX IF EXISTS idx_scheduledposts_draft_id;
DROP INDEX IF EXISTS idx_scheduledposts_user_id;

DROP INDEX IF EXISTS idx_drafts_suggestion_id;
DROP INDEX IF EXISTS idx_drafts_user_id;

DROP INDEX IF EXISTS idx_postsuggestions_suggestion_text;
DROP INDEX IF EXISTS idx_postsuggestions_prompt_id;

DROP INDEX IF EXISTS idx_prompts_prompt_text;
DROP INDEX IF EXISTS idx_prompts_user_id;

DROP INDEX IF EXISTS idx_socialaccounts_platform;
DROP INDEX IF EXISTS idx_socialaccounts_user_id;

DROP INDEX IF EXISTS idx_users_email;

-- Drop tables
DROP TABLE IF EXISTS OAuthProviders;
DROP TABLE IF EXISTS ScheduledPosts;
DROP TABLE IF EXISTS Drafts;
DROP TABLE IF EXISTS PostSuggestions;
DROP TABLE IF EXISTS Prompts;
DROP TABLE IF EXISTS SocialAccounts;
DROP TABLE IF EXISTS Users;
