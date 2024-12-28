--
-- PostgreSQL database dump
--

-- Dumped from database version 16.3
-- Dumped by pg_dump version 16.3

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

--
-- Name: pg_trgm; Type: EXTENSION; Schema: -; Owner: -
--

CREATE EXTENSION IF NOT EXISTS pg_trgm WITH SCHEMA public;


--
-- Name: EXTENSION pg_trgm; Type: COMMENT; Schema: -; Owner: -
--

COMMENT ON EXTENSION pg_trgm IS 'text similarity measurement and index searching based on trigrams';


SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: post_schedules; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.post_schedules (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    user_id uuid NOT NULL,
    post_id uuid NOT NULL,
    social_account_id uuid NOT NULL,
    scheduled_time timestamp without time zone NOT NULL,
    executed_time timestamp without time zone,
    status character varying(50) NOT NULL,
    created_at timestamp without time zone DEFAULT now() NOT NULL,
    updated_at timestamp without time zone DEFAULT now() NOT NULL
);


--
-- Name: post_suggestions; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.post_suggestions (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    prompt_id uuid NOT NULL,
    text text NOT NULL,
    status character varying(50) DEFAULT 'pending'::character varying NOT NULL,
    created_at timestamp without time zone DEFAULT now() NOT NULL,
    updated_at timestamp without time zone DEFAULT now() NOT NULL
);


--
-- Name: posts; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.posts (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    user_id uuid NOT NULL,
    suggestion_id uuid,
    text text NOT NULL,
    status character varying(50) NOT NULL,
    created_at timestamp without time zone DEFAULT now() NOT NULL,
    updated_at timestamp without time zone DEFAULT now() NOT NULL
);


--
-- Name: prompts; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.prompts (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    user_id uuid NOT NULL,
    text text NOT NULL,
    created_at timestamp without time zone DEFAULT now() NOT NULL,
    updated_at timestamp without time zone DEFAULT now() NOT NULL
);


--
-- Name: schema_migrations; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.schema_migrations (
    version bigint NOT NULL,
    dirty boolean NOT NULL
);


--
-- Name: social_accounts; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.social_accounts (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    user_id uuid NOT NULL,
    platform character varying(255) NOT NULL,
    name character varying(255) NOT NULL,
    email character varying(255) NOT NULL,
    access_token text NOT NULL,
    id_token text NOT NULL,
    token_expires_at timestamp without time zone NOT NULL,
    created_at timestamp without time zone DEFAULT now() NOT NULL,
    updated_at timestamp without time zone DEFAULT now() NOT NULL,
    linkedin_sub character varying(255)
);


--
-- Name: users; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.users (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    name character varying(255),
    email character varying(255) NOT NULL,
    password_hash character varying(255),
    is_email_verified boolean DEFAULT false NOT NULL,
    created_at timestamp without time zone DEFAULT now() NOT NULL,
    updated_at timestamp without time zone DEFAULT now() NOT NULL
);


--
-- Name: verify_emails; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.verify_emails (
    id bigint NOT NULL,
    email character varying(255) NOT NULL,
    secret_code character varying(255) NOT NULL,
    is_used boolean DEFAULT false NOT NULL,
    created_at timestamp without time zone DEFAULT now() NOT NULL,
    expired_at timestamp without time zone DEFAULT (now() + '00:15:00'::interval) NOT NULL
);


--
-- Name: verify_emails_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE public.verify_emails_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: verify_emails_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE public.verify_emails_id_seq OWNED BY public.verify_emails.id;


--
-- Name: verify_emails id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.verify_emails ALTER COLUMN id SET DEFAULT nextval('public.verify_emails_id_seq'::regclass);


--
-- Name: post_schedules post_schedules_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.post_schedules
    ADD CONSTRAINT post_schedules_pkey PRIMARY KEY (id);


--
-- Name: post_suggestions post_suggestions_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.post_suggestions
    ADD CONSTRAINT post_suggestions_pkey PRIMARY KEY (id);


--
-- Name: posts posts_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.posts
    ADD CONSTRAINT posts_pkey PRIMARY KEY (id);


--
-- Name: prompts prompts_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.prompts
    ADD CONSTRAINT prompts_pkey PRIMARY KEY (id);


--
-- Name: schema_migrations schema_migrations_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.schema_migrations
    ADD CONSTRAINT schema_migrations_pkey PRIMARY KEY (version);


--
-- Name: social_accounts social_accounts_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.social_accounts
    ADD CONSTRAINT social_accounts_pkey PRIMARY KEY (id);


--
-- Name: posts unique_post_suggestion_id; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.posts
    ADD CONSTRAINT unique_post_suggestion_id UNIQUE (suggestion_id);


--
-- Name: users users_email_key; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_email_key UNIQUE (email);


--
-- Name: users users_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);


--
-- Name: verify_emails verify_emails_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.verify_emails
    ADD CONSTRAINT verify_emails_pkey PRIMARY KEY (id);


--
-- Name: idx_post_schedules_post_id; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_post_schedules_post_id ON public.post_schedules USING btree (post_id);


--
-- Name: idx_post_schedules_scheduled_time; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_post_schedules_scheduled_time ON public.post_schedules USING btree (scheduled_time);


--
-- Name: idx_post_schedules_status; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_post_schedules_status ON public.post_schedules USING btree (status);


--
-- Name: idx_post_schedules_user_id; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_post_schedules_user_id ON public.post_schedules USING btree (user_id);


--
-- Name: idx_post_suggestions_prompt_id; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_post_suggestions_prompt_id ON public.post_suggestions USING btree (prompt_id);


--
-- Name: idx_post_suggestions_text; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_post_suggestions_text ON public.post_suggestions USING gin (text public.gin_trgm_ops);


--
-- Name: idx_posts_status; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_posts_status ON public.posts USING btree (status);


--
-- Name: idx_posts_suggestion_id; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_posts_suggestion_id ON public.posts USING btree (suggestion_id);


--
-- Name: idx_posts_user_id; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_posts_user_id ON public.posts USING btree (user_id);


--
-- Name: idx_prompts_text; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_prompts_text ON public.prompts USING gin (text public.gin_trgm_ops);


--
-- Name: idx_prompts_user_id; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_prompts_user_id ON public.prompts USING btree (user_id);


--
-- Name: idx_prompts_user_id_text; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_prompts_user_id_text ON public.prompts USING btree (user_id, text);


--
-- Name: idx_social_accounts_platform; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_social_accounts_platform ON public.social_accounts USING btree (platform);


--
-- Name: idx_social_accounts_user_id; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_social_accounts_user_id ON public.social_accounts USING btree (user_id);


--
-- Name: idx_users_email; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_users_email ON public.users USING btree (email);


--
-- Name: idx_users_name; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_users_name ON public.users USING btree (name);


--
-- Name: idx_verify_emails_email; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_verify_emails_email ON public.verify_emails USING btree (email);


--
-- Name: post_schedules post_schedules_post_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.post_schedules
    ADD CONSTRAINT post_schedules_post_id_fkey FOREIGN KEY (post_id) REFERENCES public.posts(id) ON DELETE CASCADE;


--
-- Name: post_schedules post_schedules_social_account_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.post_schedules
    ADD CONSTRAINT post_schedules_social_account_id_fkey FOREIGN KEY (social_account_id) REFERENCES public.social_accounts(id) ON DELETE CASCADE;


--
-- Name: post_schedules post_schedules_user_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.post_schedules
    ADD CONSTRAINT post_schedules_user_id_fkey FOREIGN KEY (user_id) REFERENCES public.users(id) ON DELETE CASCADE;


--
-- Name: post_suggestions post_suggestions_prompt_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.post_suggestions
    ADD CONSTRAINT post_suggestions_prompt_id_fkey FOREIGN KEY (prompt_id) REFERENCES public.prompts(id) ON DELETE CASCADE;


--
-- Name: posts posts_suggestion_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.posts
    ADD CONSTRAINT posts_suggestion_id_fkey FOREIGN KEY (suggestion_id) REFERENCES public.post_suggestions(id) ON DELETE SET NULL;


--
-- Name: posts posts_user_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.posts
    ADD CONSTRAINT posts_user_id_fkey FOREIGN KEY (user_id) REFERENCES public.users(id) ON DELETE CASCADE;


--
-- Name: prompts prompts_user_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.prompts
    ADD CONSTRAINT prompts_user_id_fkey FOREIGN KEY (user_id) REFERENCES public.users(id) ON DELETE CASCADE;


--
-- Name: social_accounts social_accounts_user_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.social_accounts
    ADD CONSTRAINT social_accounts_user_id_fkey FOREIGN KEY (user_id) REFERENCES public.users(id) ON DELETE CASCADE;


--
-- Name: verify_emails verify_emails_email_fkey; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.verify_emails
    ADD CONSTRAINT verify_emails_email_fkey FOREIGN KEY (email) REFERENCES public.users(email) ON DELETE CASCADE;


--
-- PostgreSQL database dump complete
--

