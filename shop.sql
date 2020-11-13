--
-- PostgreSQL database dump
--

-- Dumped from database version 13.0 (Debian 13.0-1.pgdg100+1)
-- Dumped by pg_dump version 13.0 (Debian 13.0-1.pgdg100+1)

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
-- Name: tv; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.tv (
    id integer NOT NULL,
    brand character varying(10),
    year date NOT NULL,
    manufacturer character varying(30) NOT NULL,
    model character varying(10) NOT NULL,
    CONSTRAINT tv_manufacturer_check CHECK ((char_length((manufacturer)::text) > 2)),
    CONSTRAINT tv_model_check CHECK ((char_length((model)::text) > 1)),
    CONSTRAINT tv_year_check CHECK ((year >= '2010-01-01'::date))
);


ALTER TABLE public.tv OWNER TO postgres;

--
-- Name: tv_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.tv_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.tv_id_seq OWNER TO postgres;

--
-- Name: tv_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.tv_id_seq OWNED BY public.tv.id;


--
-- Name: tv id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.tv ALTER COLUMN id SET DEFAULT nextval('public.tv_id_seq'::regclass);


--
-- Data for Name: tv; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.tv (id, brand, year, manufacturer, model) FROM stdin;
\.


--
-- Name: tv_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.tv_id_seq', 1, false);


--
-- Name: tv tv_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.tv
    ADD CONSTRAINT tv_pkey PRIMARY KEY (id);


--
-- PostgreSQL database dump complete
--

