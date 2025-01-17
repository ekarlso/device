--
-- PostgreSQL database dump
--

-- Dumped from database version 12.9
-- Dumped by pg_dump version 13.5 (Ubuntu 13.5-0ubuntu0.21.10.1)

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
-- Data for Name: device; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.device VALUES (1, 'mock', 'mock', 'mock', 'linux', true, 'le0jjBdynTukcfWbotjyQ1mf9IGpTmh0TkfIh1czmXA=', '10.255.240.2', NULL, NULL);


--
-- Data for Name: gateway; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.gateway VALUES (1, 'test01', 'group1', '127.0.0.1:52180', 'Hbu3CG+VjmMKpNdO6vMgUspVdz8CbG+NioJCZtTPhTQ=', '10.255.240.3', '8.8.8.8', false, NULL);
INSERT INTO public.gateway VALUES (2, 'test02', 'group1', '127.0.0.1:52181', 'jDms3Qb6gyis6g+NiBs9N222EpKp7wn+5hPPP0C6Xh8=', '10.255.240.4', '1.1.1.1', false, NULL);
INSERT INTO public.gateway VALUES (3, 'privileged01', 'group1', '127.0.0.1:52182', 'Na6jAutQ8XJRGRd+/jGmR5lLCXatzBBhISi768Qn+VE=', '10.255.240.5', '2.2.2.2', true, NULL);
INSERT INTO public.gateway VALUES (4, 'privileged02', 'group1', '127.0.0.1:52183', 'hIpbKPoX3J7zuwWNSn9ZWrrYIDwXM5NBp0J+C2ksik8=', '10.255.240.6', '4.4.4.4', true, NULL);


--
-- Data for Name: migrations; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.migrations VALUES (1, '2021-09-20 11:39:57.356329+00');
INSERT INTO public.migrations VALUES (2, '2021-09-20 11:39:57.457851+00');
INSERT INTO public.migrations VALUES (3, '2021-09-20 11:39:57.485695+00');
INSERT INTO public.migrations VALUES (4, '2021-09-21 09:25:05.324003+00');
INSERT INTO public.migrations VALUES (5, '2022-01-25 14:53:01.399203+00');


--
-- Data for Name: session; Type: TABLE DATA; Schema: public; Owner: postgres
--

INSERT INTO public.session VALUES ('8PlTPLgA491wlaY7wE75', 1, 'group1,group2', 'objectId123', '2021-09-21 18:46:05+00');
INSERT INTO public.session VALUES ('yX5X4TirbocbjLud9FHq', 1, 'group1,group2', 'objectId123', '2021-09-21 18:46:44+00');
INSERT INTO public.session VALUES ('rBWvRs0pz1Hl72s98zUs', 1, 'group1,group2', 'objectId123', '2021-09-21 18:47:09+00');


--
-- Name: device_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.device_id_seq', 2, false);


--
-- Name: gateway_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.gateway_id_seq', 5, true);


--
-- PostgreSQL database dump complete
--

