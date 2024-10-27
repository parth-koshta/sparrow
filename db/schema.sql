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
-- Name: postschedules; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.postschedules (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    user_id uuid NOT NULL,
    post_id uuid NOT NULL,
    scheduled_time timestamp without time zone NOT NULL,
    executed_time timestamp without time zone,
    status character varying(50) NOT NULL,
    created_at timestamp without time zone DEFAULT now() NOT NULL,
    updated_at timestamp without time zone DEFAULT now() NOT NULL
);


--
-- Name: postsuggestions; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.postsuggestions (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    prompt_id uuid NOT NULL,
    text text NOT NULL,
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
-- Name: scheduledposts; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.scheduledposts (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    user_id uuid NOT NULL,
    draft_id uuid NOT NULL,
    scheduled_time timestamp without time zone NOT NULL,
    status character varying(50) NOT NULL,
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
-- Name: socialaccounts; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.socialaccounts (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    user_id uuid NOT NULL,
    platform character varying(255) NOT NULL,
    account_name character varying(255) NOT NULL,
    account_email character varying(255) NOT NULL,
    access_token text NOT NULL,
    id_token text NOT NULL,
    token_expires_at timestamp without time zone NOT NULL,
    created_at timestamp without time zone DEFAULT now() NOT NULL,
    updated_at timestamp without time zone DEFAULT now() NOT NULL
);


--
-- Name: users; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.users (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    username character varying(255),
    email character varying(255) NOT NULL,
    password_hash character varying(255),
    is_email_verified boolean DEFAULT false NOT NULL,
    created_at timestamp without time zone DEFAULT now() NOT NULL,
    updated_at timestamp without time zone DEFAULT now() NOT NULL
);


--
-- Name: verifyemails; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.verifyemails (
    id bigint NOT NULL,
    email character varying(255) NOT NULL,
    secret_code character varying(255) NOT NULL,
    is_used boolean DEFAULT false NOT NULL,
    created_at timestamp without time zone DEFAULT now() NOT NULL,
    expired_at timestamp without time zone DEFAULT (now() + '00:15:00'::interval) NOT NULL
);


--
-- Name: verifyemails_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE public.verifyemails_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- Name: verifyemails_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE public.verifyemails_id_seq OWNED BY public.verifyemails.id;


--
-- Name: verifyemails id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.verifyemails ALTER COLUMN id SET DEFAULT nextval('public.verifyemails_id_seq'::regclass);


--
-- Name: posts posts_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.posts
    ADD CONSTRAINT posts_pkey PRIMARY KEY (id);


--
-- Name: postschedules postschedules_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.postschedules
    ADD CONSTRAINT postschedules_pkey PRIMARY KEY (id);


--
-- Name: postsuggestions postsuggestions_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.postsuggestions
    ADD CONSTRAINT postsuggestions_pkey PRIMARY KEY (id);


--
-- Name: prompts prompts_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.prompts
    ADD CONSTRAINT prompts_pkey PRIMARY KEY (id);


--
-- Name: scheduledposts scheduledposts_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.scheduledposts
    ADD CONSTRAINT scheduledposts_pkey PRIMARY KEY (id);


--
-- Name: schema_migrations schema_migrations_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.schema_migrations
    ADD CONSTRAINT schema_migrations_pkey PRIMARY KEY (version);


--
-- Name: socialaccounts socialaccounts_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.socialaccounts
    ADD CONSTRAINT socialaccounts_pkey PRIMARY KEY (id);


--
-- Name: postsuggestions unique_prompt_suggestion; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.postsuggestions
    ADD CONSTRAINT unique_prompt_suggestion UNIQUE (prompt_id, text);


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
-- Name: users users_username_key; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_username_key UNIQUE (username);


--
-- Name: verifyemails verifyemails_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.verifyemails
    ADD CONSTRAINT verifyemails_pkey PRIMARY KEY (id);


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
-- Name: idx_postschedules_post_id; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_postschedules_post_id ON public.postschedules USING btree (post_id);


--
-- Name: idx_postschedules_scheduled_time; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_postschedules_scheduled_time ON public.postschedules USING btree (scheduled_time);


--
-- Name: idx_postschedules_status; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_postschedules_status ON public.postschedules USING btree (status);


--
-- Name: idx_postschedules_user_id; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_postschedules_user_id ON public.postschedules USING btree (user_id);


--
-- Name: idx_postsuggestions_prompt_id; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_postsuggestions_prompt_id ON public.postsuggestions USING btree (prompt_id);


--
-- Name: idx_postsuggestions_text; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_postsuggestions_text ON public.postsuggestions USING gin (text public.gin_trgm_ops);


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
-- Name: idx_scheduledposts_draft_id; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_scheduledposts_draft_id ON public.scheduledposts USING btree (draft_id);


--
-- Name: idx_scheduledposts_scheduled_time; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_scheduledposts_scheduled_time ON public.scheduledposts USING btree (scheduled_time);


--
-- Name: idx_scheduledposts_status; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_scheduledposts_status ON public.scheduledposts USING btree (status);


--
-- Name: idx_scheduledposts_user_id; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_scheduledposts_user_id ON public.scheduledposts USING btree (user_id);


--
-- Name: idx_socialaccounts_platform; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_socialaccounts_platform ON public.socialaccounts USING btree (platform);


--
-- Name: idx_socialaccounts_user_id; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_socialaccounts_user_id ON public.socialaccounts USING btree (user_id);


--
-- Name: idx_users_email; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_users_email ON public.users USING btree (email);


--
-- Name: idx_users_username; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_users_username ON public.users USING btree (username);


--
-- Name: idx_verifyemails_email; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_verifyemails_email ON public.verifyemails USING btree (email);


--
-- Name: posts posts_suggestion_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.posts
    ADD CONSTRAINT posts_suggestion_id_fkey FOREIGN KEY (suggestion_id) REFERENCES public.postsuggestions(id) ON DELETE SET NULL;


--
-- Name: posts posts_user_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.posts
    ADD CONSTRAINT posts_user_id_fkey FOREIGN KEY (user_id) REFERENCES public.users(id) ON DELETE CASCADE;


--
-- Name: postschedules postschedules_post_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.postschedules
    ADD CONSTRAINT postschedules_post_id_fkey FOREIGN KEY (post_id) REFERENCES public.posts(id) ON DELETE CASCADE;


--
-- Name: postschedules postschedules_user_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.postschedules
    ADD CONSTRAINT postschedules_user_id_fkey FOREIGN KEY (user_id) REFERENCES public.users(id) ON DELETE CASCADE;


--
-- Name: postsuggestions postsuggestions_prompt_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.postsuggestions
    ADD CONSTRAINT postsuggestions_prompt_id_fkey FOREIGN KEY (prompt_id) REFERENCES public.prompts(id) ON DELETE CASCADE;


--
-- Name: prompts prompts_user_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.prompts
    ADD CONSTRAINT prompts_user_id_fkey FOREIGN KEY (user_id) REFERENCES public.users(id) ON DELETE CASCADE;


--
-- Name: scheduledposts scheduledposts_draft_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.scheduledposts
    ADD CONSTRAINT scheduledposts_draft_id_fkey FOREIGN KEY (draft_id) REFERENCES public.posts(id) ON DELETE CASCADE;


--
-- Name: scheduledposts scheduledposts_user_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.scheduledposts
    ADD CONSTRAINT scheduledposts_user_id_fkey FOREIGN KEY (user_id) REFERENCES public.users(id) ON DELETE CASCADE;


--
-- Name: socialaccounts socialaccounts_user_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.socialaccounts
    ADD CONSTRAINT socialaccounts_user_id_fkey FOREIGN KEY (user_id) REFERENCES public.users(id) ON DELETE CASCADE;


--
-- Name: verifyemails verifyemails_email_fkey; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.verifyemails
    ADD CONSTRAINT verifyemails_email_fkey FOREIGN KEY (email) REFERENCES public.users(email) ON DELETE CASCADE;


--
-- PostgreSQL database dump complete
--

