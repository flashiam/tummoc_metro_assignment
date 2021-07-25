--
-- PostgreSQL database dump
--

-- Dumped from database version 12.7 (Ubuntu 12.7-0ubuntu0.20.04.1)
-- Dumped by pg_dump version 13.0

-- Started on 2021-07-25 09:33:24

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
-- TOC entry 208 (class 1259 OID 83750)
-- Name: location; Type: TABLE; Schema: public; Owner: tummoc
--

CREATE TABLE public.location (
    loc_id bigint NOT NULL,
    latitude numeric(8,6) DEFAULT 0 NOT NULL,
    longitude numeric(8,6) DEFAULT 0 NOT NULL
);


ALTER TABLE public.location OWNER TO tummoc;

--
-- TOC entry 2949 (class 0 OID 83750)
-- Dependencies: 208
-- Data for Name: location; Type: TABLE DATA; Schema: public; Owner: tummoc
--

COPY public.location (loc_id, latitude, longitude) FROM stdin;
1	12.989667	77.653634
2	12.985947	77.644932
3	12.978285	77.638756
4	12.976309	77.626708
5	12.972966	77.616838
6	12.975466	77.606575
7	12.980976	77.594716
8	12.979715	77.592726
9	12.974337	77.583981
10	12.976205	77.570852
11	12.975198	77.565214
12	12.975587	77.555484
13	12.974055	77.544890
14	12.970690	77.537363
15	12.961960	77.533632
16	12.952116	77.537096
17	12.946593	77.530195
18	12.861403	77.529506
19	12.871379	77.538588
20	12.877050	77.544411
21	12.884603	77.552204
22	12.889393	77.563272
23	12.895798	77.570016
24	12.907480	77.573046
25	12.915473	77.573628
26	12.921481	77.580295
27	12.929626	77.580259
28	12.938231	77.580084
29	12.946215	77.580012
30	12.950497	77.573668
31	12.960080	77.574629
32	12.967700	77.574785
33	12.976205	77.570852
34	12.990565	77.571668
35	12.976205	77.570852
\.


--
-- TOC entry 2822 (class 2606 OID 83756)
-- Name: location location_pkey; Type: CONSTRAINT; Schema: public; Owner: tummoc
--

ALTER TABLE ONLY public.location
    ADD CONSTRAINT location_pkey PRIMARY KEY (loc_id);


-- Completed on 2021-07-25 09:33:24

--
-- PostgreSQL database dump complete
--

