running recipe
recipe finished, closing ledger
ledger closed
--
-- PostgreSQL database dump
--

-- Dumped from database version 9.6.5
-- Dumped by pg_dump version 9.6.5

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SET check_function_bodies = false;
SET client_min_messages = warning;
SET row_security = off;

SET search_path = public, pg_catalog;

DROP INDEX IF EXISTS public.signersaccount;
DROP INDEX IF EXISTS public.sellingissuerindex;
DROP INDEX IF EXISTS public.scpquorumsbyseq;
DROP INDEX IF EXISTS public.scpenvsbyseq;
DROP INDEX IF EXISTS public.priceindex;
DROP INDEX IF EXISTS public.ledgersbyseq;
DROP INDEX IF EXISTS public.histfeebyseq;
DROP INDEX IF EXISTS public.histbyseq;
DROP INDEX IF EXISTS public.buyingissuerindex;
DROP INDEX IF EXISTS public.accountbalances;
ALTER TABLE IF EXISTS ONLY public.txhistory DROP CONSTRAINT IF EXISTS txhistory_pkey;
ALTER TABLE IF EXISTS ONLY public.txfeehistory DROP CONSTRAINT IF EXISTS txfeehistory_pkey;
ALTER TABLE IF EXISTS ONLY public.trustlines DROP CONSTRAINT IF EXISTS trustlines_pkey;
ALTER TABLE IF EXISTS ONLY public.storestate DROP CONSTRAINT IF EXISTS storestate_pkey;
ALTER TABLE IF EXISTS ONLY public.signers DROP CONSTRAINT IF EXISTS signers_pkey;
ALTER TABLE IF EXISTS ONLY public.scpquorums DROP CONSTRAINT IF EXISTS scpquorums_pkey;
ALTER TABLE IF EXISTS ONLY public.pubsub DROP CONSTRAINT IF EXISTS pubsub_pkey;
ALTER TABLE IF EXISTS ONLY public.publishqueue DROP CONSTRAINT IF EXISTS publishqueue_pkey;
ALTER TABLE IF EXISTS ONLY public.peers DROP CONSTRAINT IF EXISTS peers_pkey;
ALTER TABLE IF EXISTS ONLY public.offers DROP CONSTRAINT IF EXISTS offers_pkey;
ALTER TABLE IF EXISTS ONLY public.ledgerheaders DROP CONSTRAINT IF EXISTS ledgerheaders_pkey;
ALTER TABLE IF EXISTS ONLY public.ledgerheaders DROP CONSTRAINT IF EXISTS ledgerheaders_ledgerseq_key;
ALTER TABLE IF EXISTS ONLY public.ban DROP CONSTRAINT IF EXISTS ban_pkey;
ALTER TABLE IF EXISTS ONLY public.accounts DROP CONSTRAINT IF EXISTS accounts_pkey;
ALTER TABLE IF EXISTS ONLY public.accountdata DROP CONSTRAINT IF EXISTS accountdata_pkey;
DROP TABLE IF EXISTS public.txhistory;
DROP TABLE IF EXISTS public.txfeehistory;
DROP TABLE IF EXISTS public.trustlines;
DROP TABLE IF EXISTS public.storestate;
DROP TABLE IF EXISTS public.signers;
DROP TABLE IF EXISTS public.scpquorums;
DROP TABLE IF EXISTS public.scphistory;
DROP TABLE IF EXISTS public.pubsub;
DROP TABLE IF EXISTS public.publishqueue;
DROP TABLE IF EXISTS public.peers;
DROP TABLE IF EXISTS public.offers;
DROP TABLE IF EXISTS public.ledgerheaders;
DROP TABLE IF EXISTS public.ban;
DROP TABLE IF EXISTS public.accounts;
DROP TABLE IF EXISTS public.accountdata;
DROP EXTENSION IF EXISTS plpgsql;
DROP SCHEMA IF EXISTS public;
--
-- Name: public; Type: SCHEMA; Schema: -; Owner: -
--

CREATE SCHEMA public;


--
-- Name: SCHEMA public; Type: COMMENT; Schema: -; Owner: -
--

COMMENT ON SCHEMA public IS 'standard public schema';


--
-- Name: plpgsql; Type: EXTENSION; Schema: -; Owner: -
--

CREATE EXTENSION IF NOT EXISTS plpgsql WITH SCHEMA pg_catalog;


--
-- Name: EXTENSION plpgsql; Type: COMMENT; Schema: -; Owner: -
--

COMMENT ON EXTENSION plpgsql IS 'PL/pgSQL procedural language';


SET search_path = public, pg_catalog;

SET default_tablespace = '';

SET default_with_oids = false;

--
-- Name: accountdata; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE accountdata (
    accountid character varying(56) NOT NULL,
    dataname character varying(64) NOT NULL,
    datavalue character varying(112) NOT NULL,
    lastmodified integer DEFAULT 0 NOT NULL
);


--
-- Name: accounts; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE accounts (
    accountid character varying(56) NOT NULL,
    balance bigint NOT NULL,
    seqnum bigint NOT NULL,
    numsubentries integer NOT NULL,
    inflationdest character varying(56),
    homedomain character varying(32) NOT NULL,
    thresholds text NOT NULL,
    flags integer NOT NULL,
    lastmodified integer NOT NULL,
    CONSTRAINT accounts_balance_check CHECK ((balance >= 0)),
    CONSTRAINT accounts_numsubentries_check CHECK ((numsubentries >= 0))
);


--
-- Name: ban; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE ban (
    nodeid character(56) NOT NULL
);


--
-- Name: ledgerheaders; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE ledgerheaders (
    ledgerhash character(64) NOT NULL,
    prevhash character(64) NOT NULL,
    bucketlisthash character(64) NOT NULL,
    ledgerseq integer,
    closetime bigint NOT NULL,
    data text NOT NULL,
    CONSTRAINT ledgerheaders_closetime_check CHECK ((closetime >= 0)),
    CONSTRAINT ledgerheaders_ledgerseq_check CHECK ((ledgerseq >= 0))
);


--
-- Name: offers; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE offers (
    sellerid character varying(56) NOT NULL,
    offerid bigint NOT NULL,
    sellingassettype integer NOT NULL,
    sellingassetcode character varying(12),
    sellingissuer character varying(56),
    buyingassettype integer NOT NULL,
    buyingassetcode character varying(12),
    buyingissuer character varying(56),
    amount bigint NOT NULL,
    pricen integer NOT NULL,
    priced integer NOT NULL,
    price double precision NOT NULL,
    flags integer NOT NULL,
    lastmodified integer NOT NULL,
    CONSTRAINT offers_amount_check CHECK ((amount >= 0)),
    CONSTRAINT offers_offerid_check CHECK ((offerid >= 0))
);


--
-- Name: peers; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE peers (
    ip character varying(15) NOT NULL,
    port integer DEFAULT 0 NOT NULL,
    nextattempt timestamp without time zone NOT NULL,
    numfailures integer DEFAULT 0 NOT NULL,
    CONSTRAINT peers_numfailures_check CHECK ((numfailures >= 0)),
    CONSTRAINT peers_port_check CHECK (((port > 0) AND (port <= 65535)))
);


--
-- Name: publishqueue; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE publishqueue (
    ledger integer NOT NULL,
    state text
);


--
-- Name: pubsub; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE pubsub (
    resid character(32) NOT NULL,
    lastread integer
);


--
-- Name: scphistory; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE scphistory (
    nodeid character(56) NOT NULL,
    ledgerseq integer NOT NULL,
    envelope text NOT NULL,
    CONSTRAINT scphistory_ledgerseq_check CHECK ((ledgerseq >= 0))
);


--
-- Name: scpquorums; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE scpquorums (
    qsethash character(64) NOT NULL,
    lastledgerseq integer NOT NULL,
    qset text NOT NULL,
    CONSTRAINT scpquorums_lastledgerseq_check CHECK ((lastledgerseq >= 0))
);


--
-- Name: signers; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE signers (
    accountid character varying(56) NOT NULL,
    publickey character varying(56) NOT NULL,
    weight integer NOT NULL
);


--
-- Name: storestate; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE storestate (
    statename character(32) NOT NULL,
    state text
);


--
-- Name: trustlines; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE trustlines (
    accountid character varying(56) NOT NULL,
    assettype integer NOT NULL,
    issuer character varying(56) NOT NULL,
    assetcode character varying(12) NOT NULL,
    tlimit bigint NOT NULL,
    balance bigint NOT NULL,
    flags integer NOT NULL,
    lastmodified integer NOT NULL,
    CONSTRAINT trustlines_balance_check CHECK ((balance >= 0)),
    CONSTRAINT trustlines_tlimit_check CHECK ((tlimit > 0))
);


--
-- Name: txfeehistory; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE txfeehistory (
    txid character(64) NOT NULL,
    ledgerseq integer NOT NULL,
    txindex integer NOT NULL,
    txchanges text NOT NULL,
    CONSTRAINT txfeehistory_ledgerseq_check CHECK ((ledgerseq >= 0))
);


--
-- Name: txhistory; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE txhistory (
    txid character(64) NOT NULL,
    ledgerseq integer NOT NULL,
    txindex integer NOT NULL,
    txbody text NOT NULL,
    txresult text NOT NULL,
    txmeta text NOT NULL,
    CONSTRAINT txhistory_ledgerseq_check CHECK ((ledgerseq >= 0))
);


--
-- Data for Name: accountdata; Type: TABLE DATA; Schema: public; Owner: -
--



--
-- Data for Name: accounts; Type: TABLE DATA; Schema: public; Owner: -
--

INSERT INTO accounts VALUES ('GBRPYHIL2CI3FNQ4BXLFMNDLFJUNPU2HY3ZMFSHONUCEOASW7QC7OX2H', 999999996999999700, 3, 0, NULL, '', 'AQAAAA==', 0, 2);
INSERT INTO accounts VALUES ('GBXGQJWVLWOYHFLVTKWV5FGHA3LNYY2JQKM7OAJAUEQFU6LPCSEFVXON', 999999900, 8589934593, 1, NULL, '', 'AQAAAA==', 0, 3);
INSERT INTO accounts VALUES ('GC23QF2HUE52AMXUFUH3AYJAXXGXXV2VHXYYR6EYXETPKDXZSAW67XO4', 999999900, 8589934593, 0, NULL, '', 'AQAAAA==', 0, 4);
INSERT INTO accounts VALUES ('GCXKG6RN4ONIEPCMNFB732A436Z5PNDSRLGWK7GBLCMQLIFO4S7EYWVU', 999999800, 8589934594, 1, NULL, '', 'AQAAAA==', 0, 5);


--
-- Data for Name: ban; Type: TABLE DATA; Schema: public; Owner: -
--



--
-- Data for Name: ledgerheaders; Type: TABLE DATA; Schema: public; Owner: -
--

INSERT INTO ledgerheaders VALUES ('63d98f536ee68d1b27b5b89f23af5311b7569a24faf1403ad0b52b633b07be99', '0000000000000000000000000000000000000000000000000000000000000000', '572a2e32ff248a07b0e70fd1f6d318c1facd20b6cc08c33d5775259868125a16', 1, 0, 'AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAABXKi4y/ySKB7DnD9H20xjB+s0gtswIwz1XdSWYaBJaFgAAAAEN4Lazp2QAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAZAX14QAAAABkAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA');
INSERT INTO ledgerheaders VALUES ('dd40cea954728c26c26f826f4af9b8d0b12b6d41d7d228a427e270db1c4ce7c0', '63d98f536ee68d1b27b5b89f23af5311b7569a24faf1403ad0b52b633b07be99', '9ca4f8c78b0147f710569acb911c6e14146f2bc97b20057139fc980a645363b9', 2, 1508958237, 'AAAACGPZj1Nu5o0bJ7W4nyOvUxG3Vpok+vFAOtC1K2M7B76ZP3bET1BQWO7v25lKgR/u6tNkZT/xU5qQN30ohPD6CoAAAAAAWfDgHQAAAAIAAAAIAAAAAQAAAAgAAAAIAAAAAwAAJxAAAAAATmS1LmeTcag6pi81IhaOG+MYxa34gawjnsRB3rvSgEycpPjHiwFH9xBWmsuRHG4UFG8ryXsgBXE5/JgKZFNjuQAAAAIN4Lazp2QAAAAAAAAAAAEsAAAAAAAAAAAAAAAAAAAAZAX14QAAACcQAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA');
INSERT INTO ledgerheaders VALUES ('16ab43099304ab9b6bc4d19e4c7912702b76ab8b962d03cb2897b9ce9edc50cf', 'dd40cea954728c26c26f826f4af9b8d0b12b6d41d7d228a427e270db1c4ce7c0', '64b6dea3512bbb931889ba6104058b10cd971e43ecb60cd1b6fa8c6363ff53bb', 3, 1508958238, 'AAAACN1AzqlUcowmwm+Cb0r5uNCxK21B19IopCficNscTOfABpb8iNjnaNVEKYbVMPQAwdwcmt0c3xusA3Z+rB3pAjIAAAAAWfDgHgAAAAAAAAAAN13a+3Gscvm/C1jZHTPHN4HPw5H2WdEz2kGkUix7+Fxktt6jUSu7kxiJumEEBYsQzZceQ+y2DNG2+oxjY/9TuwAAAAMN4Lazp2QAAAAAAAAAAAH0AAAAAAAAAAAAAAAAAAAAZAX14QAAACcQAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA');
INSERT INTO ledgerheaders VALUES ('b0a6ed841394876441af3b1bd61e70d98dba0c434cf6168c86174b3fd699f496', '16ab43099304ab9b6bc4d19e4c7912702b76ab8b962d03cb2897b9ce9edc50cf', '58e65f7a7e9b35595b49390ba9070002c6358925e4fa9495ad53c0ad91df0cd8', 4, 1508958239, 'AAAACBarQwmTBKuba8TRnkx5EnArdquLli0DyyiXuc6e3FDPJaCqW8JqWdRyq81/q6m82qUC1IZDRAV65iEDtv5b9wIAAAAAWfDgHwAAAAAAAAAA89aPZigAg+IMx5cMdGKxwXobEhFVZAiHObRfY4fyizFY5l96fps1WVtJOQupBwACxjWJJeT6lJWtU8Ctkd8M2AAAAAQN4Lazp2QAAAAAAAAAAAJYAAAAAAAAAAAAAAAAAAAAZAX14QAAACcQAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA');
INSERT INTO ledgerheaders VALUES ('d3d7dbfda3aa32170e127cd45c4ca631c3be6ecab6c32cb61973e2c6a385c404', 'b0a6ed841394876441af3b1bd61e70d98dba0c434cf6168c86174b3fd699f496', '1c5ca70b4cd412b979a60536101ba9f58ebb1867c8af9b88aa000e8ad1d3d898', 5, 1508958240, 'AAAACLCm7YQTlIdkQa87G9YecNmNugxDTPYWjIYXSz/WmfSWfCfhJg1mW1paGSz2hPapaJHGWobcDXSOChoYqct14lkAAAAAWfDgIAAAAAAAAAAAnpHN6r0WB8KXbaGDbS3aPdFFMJZ+yrybEFpZjcDQA/0cXKcLTNQSuXmmBTYQG6n1jrsYZ8ivm4iqAA6K0dPYmAAAAAUN4Lazp2QAAAAAAAAAAAK8AAAAAAAAAAAAAAAAAAAAZAX14QAAACcQAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA');


--
-- Data for Name: offers; Type: TABLE DATA; Schema: public; Owner: -
--



--
-- Data for Name: peers; Type: TABLE DATA; Schema: public; Owner: -
--



--
-- Data for Name: publishqueue; Type: TABLE DATA; Schema: public; Owner: -
--



--
-- Data for Name: pubsub; Type: TABLE DATA; Schema: public; Owner: -
--



--
-- Data for Name: scphistory; Type: TABLE DATA; Schema: public; Owner: -
--

INSERT INTO scphistory VALUES ('GA66ZI4V3QSN5DYBUSY7DB6I3YVJADUGGYH7SUV5RXNQUZ2AJSHLU2GE', 2, 'AAAAAD3so5XcJN6PAaSx8YfI3iqQDoY2D/lSvY3bCmdATI66AAAAAAAAAAIAAAACAAAAAQAAAEg/dsRPUFBY7u/bmUqBH+7q02RlP/FTmpA3fSiE8PoKgAAAAABZ8OAdAAAAAgAAAAgAAAABAAAACAAAAAgAAAADAAAnEAAAAAAAAAABj1NAvc7Iyl/TdMWm5m+Sns8AI3vTu893iCTbjyQp1j8AAABATXWgy95ByNAI5qlfUEq9CO8wC/u7KZr2G8pvflodIEs5Emtou/8qrkG11QsxN7HWQHHBDa7WhrBEcUdcryfmDw==');
INSERT INTO scphistory VALUES ('GA66ZI4V3QSN5DYBUSY7DB6I3YVJADUGGYH7SUV5RXNQUZ2AJSHLU2GE', 3, 'AAAAAD3so5XcJN6PAaSx8YfI3iqQDoY2D/lSvY3bCmdATI66AAAAAAAAAAMAAAACAAAAAQAAADAGlvyI2Odo1UQphtUw9ADB3Bya3RzfG6wDdn6sHekCMgAAAABZ8OAeAAAAAAAAAAAAAAABj1NAvc7Iyl/TdMWm5m+Sns8AI3vTu893iCTbjyQp1j8AAABAfJK6CWklXr/iXc9wcimSiMLuiXSlnO2SlfAkdkR4OJsJIr/dTzSrt7l+TPKEq7eSZjK4Hx84OaMW9neA89S2Aw==');
INSERT INTO scphistory VALUES ('GA66ZI4V3QSN5DYBUSY7DB6I3YVJADUGGYH7SUV5RXNQUZ2AJSHLU2GE', 4, 'AAAAAD3so5XcJN6PAaSx8YfI3iqQDoY2D/lSvY3bCmdATI66AAAAAAAAAAQAAAACAAAAAQAAADAloKpbwmpZ1HKrzX+rqbzapQLUhkNEBXrmIQO2/lv3AgAAAABZ8OAfAAAAAAAAAAAAAAABj1NAvc7Iyl/TdMWm5m+Sns8AI3vTu893iCTbjyQp1j8AAABAVB1jNqlASvh9apK3iGscL4Ob5I+8ieQR9wAdBAaJ4z1wkfCgXsESbkKULtW7qlmUwRlA0Cy4quY0Bp8XuYZBCA==');
INSERT INTO scphistory VALUES ('GA66ZI4V3QSN5DYBUSY7DB6I3YVJADUGGYH7SUV5RXNQUZ2AJSHLU2GE', 5, 'AAAAAD3so5XcJN6PAaSx8YfI3iqQDoY2D/lSvY3bCmdATI66AAAAAAAAAAUAAAACAAAAAQAAADB8J+EmDWZbWloZLPaE9qlokcZahtwNdI4KGhipy3XiWQAAAABZ8OAgAAAAAAAAAAAAAAABj1NAvc7Iyl/TdMWm5m+Sns8AI3vTu893iCTbjyQp1j8AAABAba/cATtCjD5k5EPYNWPGSQYtrDoHvm3MRRL7BBbJ043spdKOL0MITIJ0pEhWUNh+8ijvXCSsrA+wk/cx2lR8CQ==');


--
-- Data for Name: scpquorums; Type: TABLE DATA; Schema: public; Owner: -
--

INSERT INTO scpquorums VALUES ('8f5340bdcec8ca5fd374c5a6e66f929ecf00237bd3bbcf778824db8f2429d63f', 5, 'AAAAAQAAAAEAAAAAPeyjldwk3o8BpLHxh8jeKpAOhjYP+VK9jdsKZ0BMjroAAAAA');


--
-- Data for Name: signers; Type: TABLE DATA; Schema: public; Owner: -
--



--
-- Data for Name: storestate; Type: TABLE DATA; Schema: public; Owner: -
--

INSERT INTO storestate VALUES ('databaseschema                  ', '5');
INSERT INTO storestate VALUES ('forcescponnextlaunch            ', 'false');
INSERT INTO storestate VALUES ('lastclosedledger                ', 'd3d7dbfda3aa32170e127cd45c4ca631c3be6ecab6c32cb61973e2c6a385c404');
INSERT INTO storestate VALUES ('historyarchivestate             ', '{
    "version": 1,
    "server": "v0.6.3-66-gd66575cf",
    "currentLedger": 5,
    "currentBuckets": [
        {
            "curr": "d547bae554014d3b02ea524bd743c9ea59f0202a48965eba4a5d303a8c114e43",
            "next": {
                "state": 0
            },
            "snap": "db9fa4d9c5bcc26caa2ec416c6bc868fcb9fdf35a03eeabd6a844ccde91de222"
        },
        {
            "curr": "ef31a20a398ee73ce22275ea8177786bac54656f33dcc4f3fec60d55ddf163d9",
            "next": {
                "state": 1,
                "output": "db9fa4d9c5bcc26caa2ec416c6bc868fcb9fdf35a03eeabd6a844ccde91de222"
            },
            "snap": "0000000000000000000000000000000000000000000000000000000000000000"
        },
        {
            "curr": "0000000000000000000000000000000000000000000000000000000000000000",
            "next": {
                "state": 0
            },
            "snap": "0000000000000000000000000000000000000000000000000000000000000000"
        },
        {
            "curr": "0000000000000000000000000000000000000000000000000000000000000000",
            "next": {
                "state": 0
            },
            "snap": "0000000000000000000000000000000000000000000000000000000000000000"
        },
        {
            "curr": "0000000000000000000000000000000000000000000000000000000000000000",
            "next": {
                "state": 0
            },
            "snap": "0000000000000000000000000000000000000000000000000000000000000000"
        },
        {
            "curr": "0000000000000000000000000000000000000000000000000000000000000000",
            "next": {
                "state": 0
            },
            "snap": "0000000000000000000000000000000000000000000000000000000000000000"
        },
        {
            "curr": "0000000000000000000000000000000000000000000000000000000000000000",
            "next": {
                "state": 0
            },
            "snap": "0000000000000000000000000000000000000000000000000000000000000000"
        },
        {
            "curr": "0000000000000000000000000000000000000000000000000000000000000000",
            "next": {
                "state": 0
            },
            "snap": "0000000000000000000000000000000000000000000000000000000000000000"
        },
        {
            "curr": "0000000000000000000000000000000000000000000000000000000000000000",
            "next": {
                "state": 0
            },
            "snap": "0000000000000000000000000000000000000000000000000000000000000000"
        },
        {
            "curr": "0000000000000000000000000000000000000000000000000000000000000000",
            "next": {
                "state": 0
            },
            "snap": "0000000000000000000000000000000000000000000000000000000000000000"
        },
        {
            "curr": "0000000000000000000000000000000000000000000000000000000000000000",
            "next": {
                "state": 0
            },
            "snap": "0000000000000000000000000000000000000000000000000000000000000000"
        }
    ]
}');
INSERT INTO storestate VALUES ('lastscpdata                     ', 'AAAAAgAAAAA97KOV3CTejwGksfGHyN4qkA6GNg/5Ur2N2wpnQEyOugAAAAAAAAAFAAAAA49TQL3OyMpf03TFpuZvkp7PACN707vPd4gk248kKdY/AAAAAQAAADB8J+EmDWZbWloZLPaE9qlokcZahtwNdI4KGhipy3XiWQAAAABZ8OAgAAAAAAAAAAAAAAABAAAAMHwn4SYNZltaWhks9oT2qWiRxlqG3A10jgoaGKnLdeJZAAAAAFnw4CAAAAAAAAAAAAAAAECDe0UvuhldNQ6sb6wEzqZ5DP9ElsgSqVUCh16Vo1mBl+z41eKVMOB05XB3L9cfAYQH97ZqKtFfpkWBKPFkAKwNAAAAAD3so5XcJN6PAaSx8YfI3iqQDoY2D/lSvY3bCmdATI66AAAAAAAAAAUAAAACAAAAAQAAADB8J+EmDWZbWloZLPaE9qlokcZahtwNdI4KGhipy3XiWQAAAABZ8OAgAAAAAAAAAAAAAAABj1NAvc7Iyl/TdMWm5m+Sns8AI3vTu893iCTbjyQp1j8AAABAba/cATtCjD5k5EPYNWPGSQYtrDoHvm3MRRL7BBbJ043spdKOL0MITIJ0pEhWUNh+8ijvXCSsrA+wk/cx2lR8CQAAAAGwpu2EE5SHZEGvOxvWHnDZjboMQ0z2FoyGF0s/1pn0lgAAAAEAAAAArqN6LeOagjxMaUP96Bzfs9e0corNZXzBWJkFoK7kvkwAAABkAAAAAgAAAAIAAAAAAAAAAAAAAAEAAAAAAAAAAQAAAABuaCbVXZ2DlXWarV6UxwbW3GNJgpn3ASChIFp5bxSIWgAAAAFVU0QAAAAAALW4F0ehO6Ay9C0PsGEgvc1711U98Yj4mLkm9Q75kC3vAAAAAB3NZQAAAAAAAAAAAa7kvkwAAABAVs7S213+8YayvnyQekqwGQ7VFoukDEQ6P9fkP3cQp9fnHEem4aJT0+Uk7LX9qUc32Rw7G0o12+/vMKEZE9H9DwAAAAEAAAABAAAAAQAAAAA97KOV3CTejwGksfGHyN4qkA6GNg/5Ur2N2wpnQEyOugAAAAA=');


--
-- Data for Name: trustlines; Type: TABLE DATA; Schema: public; Owner: -
--

INSERT INTO trustlines VALUES ('GBXGQJWVLWOYHFLVTKWV5FGHA3LNYY2JQKM7OAJAUEQFU6LPCSEFVXON', 1, 'GC23QF2HUE52AMXUFUH3AYJAXXGXXV2VHXYYR6EYXETPKDXZSAW67XO4', 'USD', 9223372036854775807, 500000000, 1, 5);
INSERT INTO trustlines VALUES ('GCXKG6RN4ONIEPCMNFB732A436Z5PNDSRLGWK7GBLCMQLIFO4S7EYWVU', 1, 'GC23QF2HUE52AMXUFUH3AYJAXXGXXV2VHXYYR6EYXETPKDXZSAW67XO4', 'USD', 9223372036854775807, 500000000, 1, 5);


--
-- Data for Name: txfeehistory; Type: TABLE DATA; Schema: public; Owner: -
--

INSERT INTO txfeehistory VALUES ('2374e99349b9ef7dba9a5db3339b78fda8f34777b1af33ba468ad5c0df946d4d', 2, 1, 'AAAAAgAAAAMAAAABAAAAAAAAAABi/B0L0JGythwN1lY0aypo19NHxvLCyO5tBEcCVvwF9w3gtrOnZAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAQAAAAAAAAAAAAAAAAAAAAAAAAEAAAACAAAAAAAAAABi/B0L0JGythwN1lY0aypo19NHxvLCyO5tBEcCVvwF9w3gtrOnY/+cAAAAAAAAAAEAAAAAAAAAAAAAAAAAAAAAAQAAAAAAAAAAAAAAAAAAAA==');
INSERT INTO txfeehistory VALUES ('f508f24622cb2658b8b46c84fd1cd0f13b68090baae85b41e8b67cf16edf578c', 2, 2, 'AAAAAQAAAAEAAAACAAAAAAAAAABi/B0L0JGythwN1lY0aypo19NHxvLCyO5tBEcCVvwF9w3gtrOnY/84AAAAAAAAAAIAAAAAAAAAAAAAAAAAAAAAAQAAAAAAAAAAAAAAAAAAAA==');
INSERT INTO txfeehistory VALUES ('2b2e82dbabb024b27a0c3140ca71d8ac9bc71831f9f5a3bd69eca3d88fb0ec5c', 2, 3, 'AAAAAQAAAAEAAAACAAAAAAAAAABi/B0L0JGythwN1lY0aypo19NHxvLCyO5tBEcCVvwF9w3gtrOnY/7UAAAAAAAAAAMAAAAAAAAAAAAAAAAAAAAAAQAAAAAAAAAAAAAAAAAAAA==');
INSERT INTO txfeehistory VALUES ('811192c38643df73c015a5a1d77b802dff05d4f50fc6d10816aa75c0a6109f9a', 3, 1, 'AAAAAgAAAAMAAAACAAAAAAAAAABuaCbVXZ2DlXWarV6UxwbW3GNJgpn3ASChIFp5bxSIWgAAAAA7msoAAAAAAgAAAAAAAAAAAAAAAAAAAAAAAAAAAQAAAAAAAAAAAAAAAAAAAAAAAAEAAAADAAAAAAAAAABuaCbVXZ2DlXWarV6UxwbW3GNJgpn3ASChIFp5bxSIWgAAAAA7msmcAAAAAgAAAAEAAAAAAAAAAAAAAAAAAAAAAQAAAAAAAAAAAAAAAAAAAA==');
INSERT INTO txfeehistory VALUES ('bd486dbdd02d460817671c4a5a7e9d6e865ca29cb41e62d7aaf70a2fee5b36de', 3, 2, 'AAAAAgAAAAMAAAACAAAAAAAAAACuo3ot45qCPExpQ/3oHN+z17Ryis1lfMFYmQWgruS+TAAAAAA7msoAAAAAAgAAAAAAAAAAAAAAAAAAAAAAAAAAAQAAAAAAAAAAAAAAAAAAAAAAAAEAAAADAAAAAAAAAACuo3ot45qCPExpQ/3oHN+z17Ryis1lfMFYmQWgruS+TAAAAAA7msmcAAAAAgAAAAEAAAAAAAAAAAAAAAAAAAAAAQAAAAAAAAAAAAAAAAAAAA==');
INSERT INTO txfeehistory VALUES ('7736f0b869de0f74a5ed7f8d6529949238eb0f0421f3fc2bbc438084f21c8055', 4, 1, 'AAAAAgAAAAMAAAACAAAAAAAAAAC1uBdHoTugMvQtD7BhIL3Ne9dVPfGI+Ji5JvUO+ZAt7wAAAAA7msoAAAAAAgAAAAAAAAAAAAAAAAAAAAAAAAAAAQAAAAAAAAAAAAAAAAAAAAAAAAEAAAAEAAAAAAAAAAC1uBdHoTugMvQtD7BhIL3Ne9dVPfGI+Ji5JvUO+ZAt7wAAAAA7msmcAAAAAgAAAAEAAAAAAAAAAAAAAAAAAAAAAQAAAAAAAAAAAAAAAAAAAA==');
INSERT INTO txfeehistory VALUES ('3b6dc5f46ed8140595cff4b3b8cba6b7fc0bc0801fd222e4c717b9497ee5ced0', 5, 1, 'AAAAAgAAAAMAAAADAAAAAAAAAACuo3ot45qCPExpQ/3oHN+z17Ryis1lfMFYmQWgruS+TAAAAAA7msmcAAAAAgAAAAEAAAABAAAAAAAAAAAAAAAAAQAAAAAAAAAAAAAAAAAAAAAAAAEAAAAFAAAAAAAAAACuo3ot45qCPExpQ/3oHN+z17Ryis1lfMFYmQWgruS+TAAAAAA7msk4AAAAAgAAAAIAAAABAAAAAAAAAAAAAAAAAQAAAAAAAAAAAAAAAAAAAA==');


--
-- Data for Name: txhistory; Type: TABLE DATA; Schema: public; Owner: -
--

INSERT INTO txhistory VALUES ('2374e99349b9ef7dba9a5db3339b78fda8f34777b1af33ba468ad5c0df946d4d', 2, 1, 'AAAAAGL8HQvQkbK2HA3WVjRrKmjX00fG8sLI7m0ERwJW/AX3AAAAZAAAAAAAAAABAAAAAAAAAAAAAAABAAAAAAAAAAAAAAAArqN6LeOagjxMaUP96Bzfs9e0corNZXzBWJkFoK7kvkwAAAAAO5rKAAAAAAAAAAABVvwF9wAAAECDzqvkQBQoNAJifPRXDoLhvtycT3lFPCQ51gkdsFHaBNWw05S/VhW0Xgkr0CBPE4NaFV2Kmcs3ZwLmib4TRrML', 'I3Tpk0m57326ml2zM5t4/ajzR3exrzO6RorVwN+UbU0AAAAAAAAAZAAAAAAAAAABAAAAAAAAAAAAAAAAAAAAAA==', 'AAAAAAAAAAEAAAACAAAAAAAAAAIAAAAAAAAAAK6jei3jmoI8TGlD/egc37PXtHKKzWV8wViZBaCu5L5MAAAAADuaygAAAAACAAAAAAAAAAAAAAAAAAAAAAAAAAABAAAAAAAAAAAAAAAAAAAAAAAAAQAAAAIAAAAAAAAAAGL8HQvQkbK2HA3WVjRrKmjX00fG8sLI7m0ERwJW/AX3DeC2s2vJNNQAAAAAAAAAAwAAAAAAAAAAAAAAAAAAAAABAAAAAAAAAAAAAAAAAAAA');
INSERT INTO txhistory VALUES ('f508f24622cb2658b8b46c84fd1cd0f13b68090baae85b41e8b67cf16edf578c', 2, 2, 'AAAAAGL8HQvQkbK2HA3WVjRrKmjX00fG8sLI7m0ERwJW/AX3AAAAZAAAAAAAAAACAAAAAAAAAAAAAAABAAAAAAAAAAAAAAAAtbgXR6E7oDL0LQ+wYSC9zXvXVT3xiPiYuSb1DvmQLe8AAAAAO5rKAAAAAAAAAAABVvwF9wAAAEArq72awj8I0bL2hCJ9DOo1fUwxpwUnK9R5NvR/ewXcofvO+1pRNJJeTNHMGlbdlFKeqkwm3fkPP91vgLWX47oD', '9QjyRiLLJli4tGyE/RzQ8TtoCQuq6FtB6LZ88W7fV4wAAAAAAAAAZAAAAAAAAAABAAAAAAAAAAAAAAAAAAAAAA==', 'AAAAAAAAAAEAAAACAAAAAAAAAAIAAAAAAAAAALW4F0ehO6Ay9C0PsGEgvc1711U98Yj4mLkm9Q75kC3vAAAAADuaygAAAAACAAAAAAAAAAAAAAAAAAAAAAAAAAABAAAAAAAAAAAAAAAAAAAAAAAAAQAAAAIAAAAAAAAAAGL8HQvQkbK2HA3WVjRrKmjX00fG8sLI7m0ERwJW/AX3DeC2szAuatQAAAAAAAAAAwAAAAAAAAAAAAAAAAAAAAABAAAAAAAAAAAAAAAAAAAA');
INSERT INTO txhistory VALUES ('2b2e82dbabb024b27a0c3140ca71d8ac9bc71831f9f5a3bd69eca3d88fb0ec5c', 2, 3, 'AAAAAGL8HQvQkbK2HA3WVjRrKmjX00fG8sLI7m0ERwJW/AX3AAAAZAAAAAAAAAADAAAAAAAAAAAAAAABAAAAAAAAAAAAAAAAbmgm1V2dg5V1mq1elMcG1txjSYKZ9wEgoSBaeW8UiFoAAAAAO5rKAAAAAAAAAAABVvwF9wAAAEDJul1tLGLF4Vxwt0dDCVEf6tb5l4byMrGgCp+lVZMmxct54iNf2mxtjx6Md5ZJ4E4Dlcsf46EAhBGSUPsn8fYD', 'Ky6C26uwJLJ6DDFAynHYrJvHGDH59aO9aeyj2I+w7FwAAAAAAAAAZAAAAAAAAAABAAAAAAAAAAAAAAAAAAAAAA==', 'AAAAAAAAAAEAAAACAAAAAAAAAAIAAAAAAAAAAG5oJtVdnYOVdZqtXpTHBtbcY0mCmfcBIKEgWnlvFIhaAAAAADuaygAAAAACAAAAAAAAAAAAAAAAAAAAAAAAAAABAAAAAAAAAAAAAAAAAAAAAAAAAQAAAAIAAAAAAAAAAGL8HQvQkbK2HA3WVjRrKmjX00fG8sLI7m0ERwJW/AX3DeC2svSToNQAAAAAAAAAAwAAAAAAAAAAAAAAAAAAAAABAAAAAAAAAAAAAAAAAAAA');
INSERT INTO txhistory VALUES ('811192c38643df73c015a5a1d77b802dff05d4f50fc6d10816aa75c0a6109f9a', 3, 1, 'AAAAAG5oJtVdnYOVdZqtXpTHBtbcY0mCmfcBIKEgWnlvFIhaAAAAZAAAAAIAAAABAAAAAAAAAAAAAAABAAAAAAAAAAYAAAABVVNEAAAAAAC1uBdHoTugMvQtD7BhIL3Ne9dVPfGI+Ji5JvUO+ZAt73//////////AAAAAAAAAAFvFIhaAAAAQPlg7GLhJg0x7jpAw1Ew6H2XF6yRImfJIwFfx09Nui5btOJAFewFANfOaAB8FQZl5p3A5g3k6DHDigfUNUD16gc=', 'gRGSw4ZD33PAFaWh13uALf8F1PUPxtEIFqp1wKYQn5oAAAAAAAAAZAAAAAAAAAABAAAAAAAAAAYAAAAAAAAAAA==', 'AAAAAAAAAAEAAAACAAAAAAAAAAMAAAABAAAAAG5oJtVdnYOVdZqtXpTHBtbcY0mCmfcBIKEgWnlvFIhaAAAAAVVTRAAAAAAAtbgXR6E7oDL0LQ+wYSC9zXvXVT3xiPiYuSb1DvmQLe8AAAAAAAAAAH//////////AAAAAQAAAAAAAAAAAAAAAQAAAAMAAAAAAAAAAG5oJtVdnYOVdZqtXpTHBtbcY0mCmfcBIKEgWnlvFIhaAAAAADuayZwAAAACAAAAAQAAAAEAAAAAAAAAAAAAAAABAAAAAAAAAAAAAAAAAAAA');
INSERT INTO txhistory VALUES ('bd486dbdd02d460817671c4a5a7e9d6e865ca29cb41e62d7aaf70a2fee5b36de', 3, 2, 'AAAAAK6jei3jmoI8TGlD/egc37PXtHKKzWV8wViZBaCu5L5MAAAAZAAAAAIAAAABAAAAAAAAAAAAAAABAAAAAAAAAAYAAAABVVNEAAAAAAC1uBdHoTugMvQtD7BhIL3Ne9dVPfGI+Ji5JvUO+ZAt73//////////AAAAAAAAAAGu5L5MAAAAQB9kmKW2q3v7Qfy8PMekEb1TTI5ixqkI0BogXrOt7gO162Qbkh2dSTUfeDovc0PAafhDXxthVAlsLujlBmyjBAY=', 'vUhtvdAtRggXZxxKWn6dboZcopy0HmLXqvcKL+5bNt4AAAAAAAAAZAAAAAAAAAABAAAAAAAAAAYAAAAAAAAAAA==', 'AAAAAAAAAAEAAAACAAAAAAAAAAMAAAABAAAAAK6jei3jmoI8TGlD/egc37PXtHKKzWV8wViZBaCu5L5MAAAAAVVTRAAAAAAAtbgXR6E7oDL0LQ+wYSC9zXvXVT3xiPiYuSb1DvmQLe8AAAAAAAAAAH//////////AAAAAQAAAAAAAAAAAAAAAQAAAAMAAAAAAAAAAK6jei3jmoI8TGlD/egc37PXtHKKzWV8wViZBaCu5L5MAAAAADuayZwAAAACAAAAAQAAAAEAAAAAAAAAAAAAAAABAAAAAAAAAAAAAAAAAAAA');
INSERT INTO txhistory VALUES ('7736f0b869de0f74a5ed7f8d6529949238eb0f0421f3fc2bbc438084f21c8055', 4, 1, 'AAAAALW4F0ehO6Ay9C0PsGEgvc1711U98Yj4mLkm9Q75kC3vAAAAZAAAAAIAAAABAAAAAAAAAAAAAAABAAAAAAAAAAEAAAAArqN6LeOagjxMaUP96Bzfs9e0corNZXzBWJkFoK7kvkwAAAABVVNEAAAAAAC1uBdHoTugMvQtD7BhIL3Ne9dVPfGI+Ji5JvUO+ZAt7wAAAAA7msoAAAAAAAAAAAH5kC3vAAAAQDjBSAulKc9tRqGg+OkVbKPz4olRQYUevyCfv0LAlqbXG6yPbpR0BR6o7mrimRm8O4VoRBGIATQB42NOWcFzdQw=', 'dzbwuGneD3Sl7X+NZSmUkjjrDwQh8/wrvEOAhPIcgFUAAAAAAAAAZAAAAAAAAAABAAAAAAAAAAEAAAAAAAAAAA==', 'AAAAAAAAAAEAAAACAAAAAwAAAAMAAAABAAAAAK6jei3jmoI8TGlD/egc37PXtHKKzWV8wViZBaCu5L5MAAAAAVVTRAAAAAAAtbgXR6E7oDL0LQ+wYSC9zXvXVT3xiPiYuSb1DvmQLe8AAAAAAAAAAH//////////AAAAAQAAAAAAAAAAAAAAAQAAAAQAAAABAAAAAK6jei3jmoI8TGlD/egc37PXtHKKzWV8wViZBaCu5L5MAAAAAVVTRAAAAAAAtbgXR6E7oDL0LQ+wYSC9zXvXVT3xiPiYuSb1DvmQLe8AAAAAO5rKAH//////////AAAAAQAAAAAAAAAA');
INSERT INTO txhistory VALUES ('3b6dc5f46ed8140595cff4b3b8cba6b7fc0bc0801fd222e4c717b9497ee5ced0', 5, 1, 'AAAAAK6jei3jmoI8TGlD/egc37PXtHKKzWV8wViZBaCu5L5MAAAAZAAAAAIAAAACAAAAAAAAAAAAAAABAAAAAAAAAAEAAAAAbmgm1V2dg5V1mq1elMcG1txjSYKZ9wEgoSBaeW8UiFoAAAABVVNEAAAAAAC1uBdHoTugMvQtD7BhIL3Ne9dVPfGI+Ji5JvUO+ZAt7wAAAAAdzWUAAAAAAAAAAAGu5L5MAAAAQFbO0ttd/vGGsr58kHpKsBkO1RaLpAxEOj/X5D93EKfX5xxHpuGiU9PlJOy1/alHN9kcOxtKNdvv7zChGRPR/Q8=', 'O23F9G7YFAWVz/SzuMumt/wLwIAf0iLkxxe5SX7lztAAAAAAAAAAZAAAAAAAAAABAAAAAAAAAAEAAAAAAAAAAA==', 'AAAAAAAAAAEAAAAEAAAAAwAAAAMAAAABAAAAAG5oJtVdnYOVdZqtXpTHBtbcY0mCmfcBIKEgWnlvFIhaAAAAAVVTRAAAAAAAtbgXR6E7oDL0LQ+wYSC9zXvXVT3xiPiYuSb1DvmQLe8AAAAAAAAAAH//////////AAAAAQAAAAAAAAAAAAAAAQAAAAUAAAABAAAAAG5oJtVdnYOVdZqtXpTHBtbcY0mCmfcBIKEgWnlvFIhaAAAAAVVTRAAAAAAAtbgXR6E7oDL0LQ+wYSC9zXvXVT3xiPiYuSb1DvmQLe8AAAAAHc1lAH//////////AAAAAQAAAAAAAAAAAAAAAwAAAAQAAAABAAAAAK6jei3jmoI8TGlD/egc37PXtHKKzWV8wViZBaCu5L5MAAAAAVVTRAAAAAAAtbgXR6E7oDL0LQ+wYSC9zXvXVT3xiPiYuSb1DvmQLe8AAAAAO5rKAH//////////AAAAAQAAAAAAAAAAAAAAAQAAAAUAAAABAAAAAK6jei3jmoI8TGlD/egc37PXtHKKzWV8wViZBaCu5L5MAAAAAVVTRAAAAAAAtbgXR6E7oDL0LQ+wYSC9zXvXVT3xiPiYuSb1DvmQLe8AAAAAHc1lAH//////////AAAAAQAAAAAAAAAA');


--
-- Name: accountdata accountdata_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY accountdata
    ADD CONSTRAINT accountdata_pkey PRIMARY KEY (accountid, dataname);


--
-- Name: accounts accounts_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY accounts
    ADD CONSTRAINT accounts_pkey PRIMARY KEY (accountid);


--
-- Name: ban ban_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY ban
    ADD CONSTRAINT ban_pkey PRIMARY KEY (nodeid);


--
-- Name: ledgerheaders ledgerheaders_ledgerseq_key; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY ledgerheaders
    ADD CONSTRAINT ledgerheaders_ledgerseq_key UNIQUE (ledgerseq);


--
-- Name: ledgerheaders ledgerheaders_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY ledgerheaders
    ADD CONSTRAINT ledgerheaders_pkey PRIMARY KEY (ledgerhash);


--
-- Name: offers offers_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY offers
    ADD CONSTRAINT offers_pkey PRIMARY KEY (offerid);


--
-- Name: peers peers_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY peers
    ADD CONSTRAINT peers_pkey PRIMARY KEY (ip, port);


--
-- Name: publishqueue publishqueue_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY publishqueue
    ADD CONSTRAINT publishqueue_pkey PRIMARY KEY (ledger);


--
-- Name: pubsub pubsub_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY pubsub
    ADD CONSTRAINT pubsub_pkey PRIMARY KEY (resid);


--
-- Name: scpquorums scpquorums_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY scpquorums
    ADD CONSTRAINT scpquorums_pkey PRIMARY KEY (qsethash);


--
-- Name: signers signers_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY signers
    ADD CONSTRAINT signers_pkey PRIMARY KEY (accountid, publickey);


--
-- Name: storestate storestate_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY storestate
    ADD CONSTRAINT storestate_pkey PRIMARY KEY (statename);


--
-- Name: trustlines trustlines_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY trustlines
    ADD CONSTRAINT trustlines_pkey PRIMARY KEY (accountid, issuer, assetcode);


--
-- Name: txfeehistory txfeehistory_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY txfeehistory
    ADD CONSTRAINT txfeehistory_pkey PRIMARY KEY (ledgerseq, txindex);


--
-- Name: txhistory txhistory_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY txhistory
    ADD CONSTRAINT txhistory_pkey PRIMARY KEY (ledgerseq, txindex);


--
-- Name: accountbalances; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX accountbalances ON accounts USING btree (balance) WHERE (balance >= 1000000000);


--
-- Name: buyingissuerindex; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX buyingissuerindex ON offers USING btree (buyingissuer);


--
-- Name: histbyseq; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX histbyseq ON txhistory USING btree (ledgerseq);


--
-- Name: histfeebyseq; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX histfeebyseq ON txfeehistory USING btree (ledgerseq);


--
-- Name: ledgersbyseq; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX ledgersbyseq ON ledgerheaders USING btree (ledgerseq);


--
-- Name: priceindex; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX priceindex ON offers USING btree (price);


--
-- Name: scpenvsbyseq; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX scpenvsbyseq ON scphistory USING btree (ledgerseq);


--
-- Name: scpquorumsbyseq; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX scpquorumsbyseq ON scpquorums USING btree (lastledgerseq);


--
-- Name: sellingissuerindex; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX sellingissuerindex ON offers USING btree (sellingissuer);


--
-- Name: signersaccount; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX signersaccount ON signers USING btree (accountid);


--
-- PostgreSQL database dump complete
--

