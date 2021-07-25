--
-- PostgreSQL database dump
--

-- Dumped from database version 12.7 (Ubuntu 12.7-0ubuntu0.20.04.1)
-- Dumped by pg_dump version 13.0

-- Started on 2021-07-25 09:35:37

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
-- TOC entry 210 (class 1259 OID 83759)
-- Name: sprint_time_keys; Type: TABLE; Schema: public; Owner: tummoc
--

CREATE TABLE public.sprint_time_keys (
    id integer NOT NULL,
    sprint_id bigint NOT NULL,
    time_key_id bigint NOT NULL
);


ALTER TABLE public.sprint_time_keys OWNER TO tummoc;

--
-- TOC entry 209 (class 1259 OID 83757)
-- Name: sprint_time_keys_id_seq; Type: SEQUENCE; Schema: public; Owner: tummoc
--

CREATE SEQUENCE public.sprint_time_keys_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.sprint_time_keys_id_seq OWNER TO tummoc;

--
-- TOC entry 2955 (class 0 OID 0)
-- Dependencies: 209
-- Name: sprint_time_keys_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: tummoc
--

ALTER SEQUENCE public.sprint_time_keys_id_seq OWNED BY public.sprint_time_keys.id;


--
-- TOC entry 2819 (class 2604 OID 83762)
-- Name: sprint_time_keys id; Type: DEFAULT; Schema: public; Owner: tummoc
--

ALTER TABLE ONLY public.sprint_time_keys ALTER COLUMN id SET DEFAULT nextval('public.sprint_time_keys_id_seq'::regclass);


--
-- TOC entry 2949 (class 0 OID 83759)
-- Dependencies: 210
-- Data for Name: sprint_time_keys; Type: TABLE DATA; Schema: public; Owner: tummoc
--

COPY public.sprint_time_keys (id, sprint_id, time_key_id) FROM stdin;
1	1	1
2	1	2
3	1	3
4	1	4
5	1	5
6	1	6
7	1	7
8	1	8
9	1	9
10	1	10
11	1	11
12	1	12
13	1	13
14	1	14
15	1	15
16	1	16
17	1	17
18	7	18
19	7	19
20	7	20
21	7	21
22	7	22
23	7	23
24	7	24
25	7	25
26	7	26
27	7	27
28	7	28
29	7	29
30	7	30
31	7	31
32	7	32
33	7	33
34	7	34
35	7	35
\.


--
-- TOC entry 2956 (class 0 OID 0)
-- Dependencies: 209
-- Name: sprint_time_keys_id_seq; Type: SEQUENCE SET; Schema: public; Owner: tummoc
--

SELECT pg_catalog.setval('public.sprint_time_keys_id_seq', 1, false);


--
-- TOC entry 2821 (class 2606 OID 83764)
-- Name: sprint_time_keys sprint_time_keys_pkey; Type: CONSTRAINT; Schema: public; Owner: tummoc
--

ALTER TABLE ONLY public.sprint_time_keys
    ADD CONSTRAINT sprint_time_keys_pkey PRIMARY KEY (id);


-- Completed on 2021-07-25 09:35:37

--
-- PostgreSQL database dump complete
--

