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
-- Name: drafts; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.drafts (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    user_id uuid NOT NULL,
    suggestion_id uuid NOT NULL,
    draft_text text NOT NULL,
    created_at timestamp without time zone DEFAULT now() NOT NULL,
    updated_at timestamp without time zone DEFAULT now() NOT NULL
);


--
-- Name: postsuggestions; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.postsuggestions (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    prompt_id uuid NOT NULL,
    suggestion_text text NOT NULL,
    created_at timestamp without time zone DEFAULT now() NOT NULL,
    updated_at timestamp without time zone DEFAULT now() NOT NULL
);


--
-- Name: prompts; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.prompts (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    user_id uuid NOT NULL,
    prompt_text text NOT NULL,
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
    created_at timestamp without time zone DEFAULT now() NOT NULL,
    updated_at timestamp without time zone DEFAULT now() NOT NULL
);


--
-- Name: drafts drafts_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.drafts
    ADD CONSTRAINT drafts_pkey PRIMARY KEY (id);


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
    ADD CONSTRAINT unique_prompt_suggestion UNIQUE (prompt_id, suggestion_text);


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
-- Name: idx_drafts_suggestion_id; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_drafts_suggestion_id ON public.drafts USING btree (suggestion_id);


--
-- Name: idx_drafts_user_id; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_drafts_user_id ON public.drafts USING btree (user_id);


--
-- Name: idx_postsuggestions_prompt_id; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_postsuggestions_prompt_id ON public.postsuggestions USING btree (prompt_id);


--
-- Name: idx_postsuggestions_suggestion_text; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_postsuggestions_suggestion_text ON public.postsuggestions USING gin (suggestion_text public.gin_trgm_ops);


--
-- Name: idx_prompts_prompt_text; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_prompts_prompt_text ON public.prompts USING gin (prompt_text public.gin_trgm_ops);


--
-- Name: idx_prompts_user_id; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_prompts_user_id ON public.prompts USING btree (user_id);


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
-- Name: drafts drafts_suggestion_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.drafts
    ADD CONSTRAINT drafts_suggestion_id_fkey FOREIGN KEY (suggestion_id) REFERENCES public.postsuggestions(id) ON DELETE CASCADE;


--
-- Name: drafts drafts_user_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.drafts
    ADD CONSTRAINT drafts_user_id_fkey FOREIGN KEY (user_id) REFERENCES public.users(id) ON DELETE CASCADE;


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
    ADD CONSTRAINT scheduledposts_draft_id_fkey FOREIGN KEY (draft_id) REFERENCES public.drafts(id) ON DELETE CASCADE;


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
-- PostgreSQL database dump complete
--

