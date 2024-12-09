-- Drop indexes for verify_emails table
DROP INDEX IF EXISTS idx_verify_emails_email;

-- Drop verify_emails table
DROP TABLE IF EXISTS verify_emails;

-- Drop indexes for post_schedules table
DROP INDEX IF EXISTS idx_post_schedules_status;
DROP INDEX IF EXISTS idx_post_schedules_scheduled_time;
DROP INDEX IF EXISTS idx_post_schedules_post_id;
DROP INDEX IF EXISTS idx_post_schedules_user_id;

-- Drop post_schedules table
DROP TABLE IF EXISTS post_schedules;

-- Drop indexes for posts table
DROP INDEX IF EXISTS idx_posts_status;
DROP INDEX IF EXISTS idx_posts_suggestion_id;
DROP INDEX IF EXISTS idx_posts_user_id;

-- Drop posts table
DROP TABLE IF EXISTS posts;

-- Drop indexes for post_suggestions table
DROP INDEX IF EXISTS idx_post_suggestions_text;
DROP INDEX IF EXISTS idx_post_suggestions_prompt_id;

-- Drop post_suggestions table
DROP TABLE IF EXISTS post_suggestions;

-- Drop indexes for prompts table
DROP INDEX IF EXISTS idx_prompts_user_id_text;
DROP INDEX IF EXISTS idx_prompts_text;
DROP INDEX IF EXISTS idx_prompts_user_id;

-- Drop prompts table
DROP TABLE IF EXISTS prompts;

-- Drop indexes for social_accounts table
DROP INDEX IF EXISTS idx_social_accounts_platform;
DROP INDEX IF EXISTS idx_social_accounts_user_id;

-- Drop social_accounts table
DROP TABLE IF EXISTS social_accounts;

-- Drop indexes for users table
DROP INDEX IF EXISTS idx_users_name;
DROP INDEX IF EXISTS idx_users_email;

-- Drop users table
DROP TABLE IF EXISTS users;

-- Drop pg_trgm extension
DROP EXTENSION IF EXISTS pg_trgm;
