--
-- PostgreSQL database dump
--

-- Dumped from database version 12.7 (Ubuntu 12.7-0ubuntu0.20.04.1)
-- Dumped by pg_dump version 13.0

-- Started on 2021-07-25 09:36:28

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
-- TOC entry 206 (class 1259 OID 83736)
-- Name: time_key; Type: TABLE; Schema: public; Owner: tummoc
--

CREATE TABLE public.time_key (
    key_id bigint NOT NULL,
    center_id bigint NOT NULL,
    up_id bigint NOT NULL,
    down_id bigint NOT NULL,
    time_value_up smallint DEFAULT 0 NOT NULL,
    time_value_down smallint DEFAULT 0 NOT NULL,
    CONSTRAINT time_key_time_value_down_check CHECK (((time_value_down >= '-127'::integer) AND (time_value_down <= 128))),
    CONSTRAINT time_key_time_value_up_check CHECK (((time_value_up >= '-127'::integer) AND (time_value_up <= 128)))
);


ALTER TABLE public.time_key OWNER TO tummoc;

--
-- TOC entry 2951 (class 0 OID 83736)
-- Dependencies: 206
-- Data for Name: time_key; Type: TABLE DATA; Schema: public; Owner: tummoc
--

COPY public.time_key (key_id, center_id, up_id, down_id, time_value_up, time_value_down) FROM stdin;
1	1	0	2	0	3
2	2	1	3	0	4
3	3	2	4	0	4
4	4	3	5	0	3
5	5	4	6	0	4
6	6	5	7	0	3
7	7	6	8	0	3
8	8	7	9	0	3
9	9	8	10	0	3
10	10	9	11	0	3
11	11	10	9	0	4
12	12	11	13	0	3
13	13	12	14	0	3
14	14	13	15	0	4
15	15	14	16	0	3
16	16	15	17	0	3
17	17	16	0	0	0
18	18	0	19	0	3
19	19	18	20	0	3
20	20	19	21	0	3
21	21	20	22	0	3
22	22	21	23	0	3
23	23	22	24	0	3
24	24	23	25	0	3
25	25	24	26	0	3
26	26	25	27	0	3
27	27	26	28	0	3
28	28	27	29	0	3
29	29	28	30	0	3
30	30	29	31	0	3
31	31	30	32	0	3
32	32	31	33	0	3
33	33	32	34	0	3
34	34	33	35	0	3
35	35	34	0	0	0
\.


--
-- TOC entry 2824 (class 2606 OID 83744)
-- Name: time_key time_key_pkey; Type: CONSTRAINT; Schema: public; Owner: tummoc
--

ALTER TABLE ONLY public.time_key
    ADD CONSTRAINT time_key_pkey PRIMARY KEY (key_id);


-- Completed on 2021-07-25 09:36:29

--
-- PostgreSQL database dump complete
--

