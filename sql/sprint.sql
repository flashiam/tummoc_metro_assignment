--
-- PostgreSQL database dump
--

-- Dumped from database version 12.7 (Ubuntu 12.7-0ubuntu0.20.04.1)
-- Dumped by pg_dump version 13.0

-- Started on 2021-07-25 09:35:17

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
-- TOC entry 207 (class 1259 OID 83745)
-- Name: sprint; Type: TABLE; Schema: public; Owner: tummoc
--

CREATE TABLE public.sprint (
    sprint_id bigint NOT NULL,
    start_time timestamp with time zone NOT NULL,
    end_time timestamp with time zone NOT NULL,
    route_id bigint NOT NULL
);


ALTER TABLE public.sprint OWNER TO tummoc;

--
-- TOC entry 2947 (class 0 OID 83745)
-- Dependencies: 207
-- Data for Name: sprint; Type: TABLE DATA; Schema: public; Owner: tummoc
--

COPY public.sprint (sprint_id, start_time, end_time, route_id) FROM stdin;
1	2016-06-22 07:00:00+00	2016-06-22 07:53:00+00	1
2	2016-06-22 07:12:00+00	2016-06-22 08:03:00+00	1
3	2016-06-22 07:24:00+00	2016-06-22 08:13:00+00	1
4	2016-06-22 07:36:00+00	2016-06-22 08:29:00+00	1
5	2016-06-22 07:48:00+00	2016-06-22 08:41:00+00	1
6	2016-06-22 08:00:00+00	2016-06-22 08:53:00+00	1
7	2016-06-22 07:00:00+00	2016-06-22 07:52:00+00	2
8	2016-06-22 07:12:00+00	2016-06-22 08:12:00+00	2
9	2016-06-22 07:24:00+00	2016-06-22 08:16:00+00	2
10	2016-06-22 07:36:00+00	2016-06-22 08:28:00+00	2
11	2016-06-22 07:48:00+00	2016-06-22 08:40:00+00	2
12	2016-06-22 08:00:00+00	2016-06-22 08:52:00+00	2
\.


--
-- TOC entry 2820 (class 2606 OID 83749)
-- Name: sprint sprint_pkey; Type: CONSTRAINT; Schema: public; Owner: tummoc
--

ALTER TABLE ONLY public.sprint
    ADD CONSTRAINT sprint_pkey PRIMARY KEY (sprint_id);


-- Completed on 2021-07-25 09:35:18

--
-- PostgreSQL database dump complete
--

