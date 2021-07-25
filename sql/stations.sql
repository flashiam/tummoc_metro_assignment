--
-- PostgreSQL database dump
--

-- Dumped from database version 12.7 (Ubuntu 12.7-0ubuntu0.20.04.1)
-- Dumped by pg_dump version 13.0

-- Started on 2021-07-25 09:36:00

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
-- TOC entry 205 (class 1259 OID 83725)
-- Name: station; Type: TABLE; Schema: public; Owner: tummoc
--

CREATE TABLE public.station (
    station_id bigint NOT NULL,
    name text DEFAULT ''::text NOT NULL,
    location_id bigint NOT NULL
);


ALTER TABLE public.station OWNER TO tummoc;

--
-- TOC entry 2950 (class 0 OID 83725)
-- Dependencies: 205
-- Data for Name: station; Type: TABLE DATA; Schema: public; Owner: tummoc
--

COPY public.station (station_id, name, location_id) FROM stdin;
1	Baiyappanhalli Metro Station	1
2	Swamy Vivekananda Road Metro Station	2
3	Indiranagar Metro Station	3
4	Halasuru Metro Station	4
5	Trinity Metro Station	5
6	MG Road Metro Station	6
7	Cubbon Park Metro	7
8	Vidhana Soudha Metro Station	8
9	Sir M. Visveshwarya Station	9
10	Majestic Metro Station	10
11	City Railway Metro Station	11
12	Magadi Road Metro Station	12
13	Hosahalli Metro Station	13
14	Vijaynagar Metro Station	14
15	Attiguppe Metro Station	15
16	Deepanjali Nagar Metro Station	16
17	Mysore Road Metro Station	17
18	Silk Institute Metro Station	18
19	Thalagattapura Metro Station\n	19
20	Vajarahalli Metro Station	20
21	Doddakallasandra Metro Station	21
22	Konanakunte Cross Metro Station	22
23	Yelachenahalli Metro Station	23
24	Jayaprakash Nagar Metro Station	24
25	Banashankari Metro Station	25
26	Rashtreeya Vidyalaya Road Metro Station	26
27	Jayanagar Metro Station	27
28	Southend Circle Metro Station	28
29	Lalbagh Metro Station	29
30	National College Metro Station	30
31	Krishna Rajendra Market Metro Station	31
32	Chickpete Metro Station	32
33	Majestic Metro Station	33
34	Sampige Road Metro Station	34
35	Srirampura Metro Station	35
\.


--
-- TOC entry 2821 (class 2606 OID 83735)
-- Name: station station_location_id_key; Type: CONSTRAINT; Schema: public; Owner: tummoc
--

ALTER TABLE ONLY public.station
    ADD CONSTRAINT station_location_id_key UNIQUE (location_id);


--
-- TOC entry 2823 (class 2606 OID 83733)
-- Name: station station_pkey; Type: CONSTRAINT; Schema: public; Owner: tummoc
--

ALTER TABLE ONLY public.station
    ADD CONSTRAINT station_pkey PRIMARY KEY (station_id);


-- Completed on 2021-07-25 09:36:00

--
-- PostgreSQL database dump complete
--

