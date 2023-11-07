--
-- PostgreSQL database dump
--

-- Dumped from database version 12.3 (Debian 12.3-1.pgdg100+1)
-- Dumped by pg_dump version 12.16

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

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: tfa; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.tfa (
    user_id character varying(254) NOT NULL,
    created_at timestamp(0) with time zone,
    deleted_at timestamp(0) with time zone,
    phone character varying(254),
    mail character varying(254),
    phone_updated_at timestamp(0) with time zone,
    mail_updated_at timestamp(0) with time zone,
    token_data text
);


ALTER TABLE public.tfa OWNER TO postgres;

--
-- Name: tfa tfa_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.tfa
    ADD CONSTRAINT tfa_pkey PRIMARY KEY (user_id);


--
-- Name: tfamail; Type: INDEX; Schema: public; Owner: postgres
--

CREATE UNIQUE INDEX tfamail ON public.tfa USING btree (mail);


--
-- Name: tfaphone; Type: INDEX; Schema: public; Owner: postgres
--

CREATE UNIQUE INDEX tfaphone ON public.tfa USING btree (phone);


--
-- PostgreSQL database dump complete
--

