--
-- PostgreSQL database dump
--

-- Dumped from database version 12.7 (Ubuntu 12.7-0ubuntu0.20.04.1)
-- Dumped by pg_dump version 13.0

-- Started on 2021-07-25 09:34:50

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
-- TOC entry 204 (class 1259 OID 83716)
-- Name: route; Type: TABLE; Schema: public; Owner: tummoc
--

CREATE TABLE public.route (
    route_id bigint NOT NULL,
    start_station_id bigint NOT NULL,
    end_station_id bigint NOT NULL
);


ALTER TABLE public.route OWNER TO tummoc;

--
-- TOC entry 2951 (class 0 OID 83716)
-- Dependencies: 204
-- Data for Name: route; Type: TABLE DATA; Schema: public; Owner: tummoc
--

COPY public.route (route_id, start_station_id, end_station_id) FROM stdin;
1	1	17
2	18	35
\.


--
-- TOC entry 2820 (class 2606 OID 83724)
-- Name: route route_end_station_id_key; Type: CONSTRAINT; Schema: public; Owner: tummoc
--

ALTER TABLE ONLY public.route
    ADD CONSTRAINT route_end_station_id_key UNIQUE (end_station_id);


--
-- TOC entry 2822 (class 2606 OID 83720)
-- Name: route route_pkey; Type: CONSTRAINT; Schema: public; Owner: tummoc
--

ALTER TABLE ONLY public.route
    ADD CONSTRAINT route_pkey PRIMARY KEY (route_id);


--
-- TOC entry 2824 (class 2606 OID 83722)
-- Name: route route_start_station_id_key; Type: CONSTRAINT; Schema: public; Owner: tummoc
--

ALTER TABLE ONLY public.route
    ADD CONSTRAINT route_start_station_id_key UNIQUE (start_station_id);


-- Completed on 2021-07-25 09:34:51

--
-- PostgreSQL database dump complete
--

