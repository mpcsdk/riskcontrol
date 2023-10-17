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

--
-- Name: notify_update(); Type: FUNCTION; Schema: public; Owner: postgres
--

CREATE FUNCTION public.notify_update() RETURNS trigger
    LANGUAGE plpgsql
    AS $$
DECLARE
    changed_column_names text;
		updated_rows RECORD;
BEGIN
   -- 获取所有更新过的字段名称
		
		if TG_OP = 'DELETE' THEN
				RAISE NOTICE '这是一条输出消息:%', OLD;

			if TG_NARGS = 0 THEN
				PERFORM pg_notify('rule_ch', OLD.rule_id||',rm');
			else
				for updated_rows in select * from OLD LOOP
					PERFORM pg_notify('rule_ch', rdata.rule_id||',rm');
				END LOOP;
			end if;
		else 
				RAISE NOTICE '这是一条输出消息:%', NEW;

			if TG_NARGS = 0 THEN
				PERFORM pg_notify('rule_ch', NEW.rule_id||',up');
			else
				for updated_rows in select * from NEW LOOP
					PERFORM pg_notify('rule_ch', rdata.rule_id||',up');
				END LOOP;
			end if;
		END if;
    RETURN NEW;
END;
$$;


ALTER FUNCTION public.notify_update() OWNER TO postgres;

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: agg_ft_24h; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.agg_ft_24h (
    "from" character varying(255),
    "to" character varying(255),
    value numeric,
    contract character varying(255),
    updated_at timestamp(6) without time zone,
    method_name character varying(255),
    "fromBlock" bigint,
    "toBlock" bigint,
    method_sig character varying(255),
    ft_name character varying(255)
);


ALTER TABLE public.agg_ft_24h OWNER TO postgres;

--
-- Name: agg_nft_24h; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.agg_nft_24h (
    "from" character varying(255),
    "to" character varying(255),
    value bigint,
    contract character varying(255),
    updated_at timestamp(6) without time zone,
    method_name character varying(255),
    "fromBlock" bigint,
    "toBlock" bigint,
    method_sig character varying(255),
    nft_name character varying(255)
);


ALTER TABLE public.agg_nft_24h OWNER TO postgres;

--
-- Name: casbin_rule; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.casbin_rule (
    id bigint NOT NULL,
    ptype character varying(100),
    v0 character varying(100),
    v1 character varying(100),
    v2 character varying(100),
    v3 character varying(100),
    v4 character varying(100),
    v5 character varying(100)
);


ALTER TABLE public.casbin_rule OWNER TO postgres;

--
-- Name: casbin_rule_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.casbin_rule_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.casbin_rule_id_seq OWNER TO postgres;

--
-- Name: casbin_rule_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.casbin_rule_id_seq OWNED BY public.casbin_rule.id;


--
-- Name: contract_abi; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.contract_abi (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    addr text,
    abi text
);


ALTER TABLE public.contract_abi OWNER TO postgres;

--
-- Name: contract_abi_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.contract_abi_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.contract_abi_id_seq OWNER TO postgres;

--
-- Name: contract_abi_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.contract_abi_id_seq OWNED BY public.contract_abi.id;


--
-- Name: eth_tx_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.eth_tx_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.eth_tx_id_seq OWNER TO postgres;

--
-- Name: eth_tx; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.eth_tx (
    id bigint DEFAULT nextval('public.eth_tx_id_seq'::regclass) NOT NULL,
    created_at timestamp(6) without time zone,
    updated_at timestamp(6) without time zone,
    deleted_at timestamp(6) without time zone,
    address text,
    contract text,
    method_name text,
    method_sig text,
    event_name text,
    event_sig text,
    topics json,
    "from" text,
    "to" text,
    value text,
    kind character varying,
    block_number bigint,
    block_hash character varying,
    tx_hash character varying,
    tx_index bigint,
    log_index bigint,
    data character varying,
    name character varying(255)
);


ALTER TABLE public.eth_tx OWNER TO postgres;

--
-- Name: exa_customers; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.exa_customers (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    customer_name text,
    customer_phone_data text,
    sys_user_id bigint,
    sys_user_authority_id bigint
);


ALTER TABLE public.exa_customers OWNER TO postgres;

--
-- Name: COLUMN exa_customers.customer_name; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.exa_customers.customer_name IS '客户名';


--
-- Name: COLUMN exa_customers.customer_phone_data; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.exa_customers.customer_phone_data IS '客户手机号';


--
-- Name: COLUMN exa_customers.sys_user_id; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.exa_customers.sys_user_id IS '管理ID';


--
-- Name: COLUMN exa_customers.sys_user_authority_id; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.exa_customers.sys_user_authority_id IS '管理角色ID';


--
-- Name: exa_customers_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.exa_customers_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.exa_customers_id_seq OWNER TO postgres;

--
-- Name: exa_customers_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.exa_customers_id_seq OWNED BY public.exa_customers.id;


--
-- Name: exa_file_chunks; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.exa_file_chunks (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    exa_file_id bigint,
    file_chunk_number bigint,
    file_chunk_path text
);


ALTER TABLE public.exa_file_chunks OWNER TO postgres;

--
-- Name: exa_file_chunks_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.exa_file_chunks_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.exa_file_chunks_id_seq OWNER TO postgres;

--
-- Name: exa_file_chunks_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.exa_file_chunks_id_seq OWNED BY public.exa_file_chunks.id;


--
-- Name: exa_file_upload_and_downloads; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.exa_file_upload_and_downloads (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    name text,
    url text,
    tag text,
    key text
);


ALTER TABLE public.exa_file_upload_and_downloads OWNER TO postgres;

--
-- Name: COLUMN exa_file_upload_and_downloads.name; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.exa_file_upload_and_downloads.name IS '文件名';


--
-- Name: COLUMN exa_file_upload_and_downloads.url; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.exa_file_upload_and_downloads.url IS '文件地址';


--
-- Name: COLUMN exa_file_upload_and_downloads.tag; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.exa_file_upload_and_downloads.tag IS '文件标签';


--
-- Name: COLUMN exa_file_upload_and_downloads.key; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.exa_file_upload_and_downloads.key IS '编号';


--
-- Name: exa_file_upload_and_downloads_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.exa_file_upload_and_downloads_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.exa_file_upload_and_downloads_id_seq OWNER TO postgres;

--
-- Name: exa_file_upload_and_downloads_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.exa_file_upload_and_downloads_id_seq OWNED BY public.exa_file_upload_and_downloads.id;


--
-- Name: exa_files; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.exa_files (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    file_name text,
    file_md5 text,
    file_path text,
    chunk_total bigint,
    is_finish boolean
);


ALTER TABLE public.exa_files OWNER TO postgres;

--
-- Name: exa_files_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.exa_files_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.exa_files_id_seq OWNER TO postgres;

--
-- Name: exa_files_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.exa_files_id_seq OWNED BY public.exa_files.id;


--
-- Name: jwt_blacklists; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.jwt_blacklists (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    jwt text
);


ALTER TABLE public.jwt_blacklists OWNER TO postgres;

--
-- Name: COLUMN jwt_blacklists.jwt; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.jwt_blacklists.jwt IS 'jwt';


--
-- Name: jwt_blacklists_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.jwt_blacklists_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.jwt_blacklists_id_seq OWNER TO postgres;

--
-- Name: jwt_blacklists_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.jwt_blacklists_id_seq OWNED BY public.jwt_blacklists.id;


--
-- Name: mpc_context; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.mpc_context (
    user_id character varying NOT NULL,
    context character varying,
    updated_at timestamp(6) without time zone,
    request character varying,
    token character varying,
    created_at timestamp(6) without time zone,
    deleted_at timestamp(6) without time zone,
    pub_key character varying
);


ALTER TABLE public.mpc_context OWNER TO postgres;

--
-- Name: rule; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.rule (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    rule_id text,
    rule_desc text,
    rules text
);


ALTER TABLE public.rule OWNER TO postgres;

--
-- Name: rule_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.rule_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.rule_id_seq OWNER TO postgres;

--
-- Name: rule_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.rule_id_seq OWNED BY public.rule.id;


--
-- Name: scrape_logs_stat; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.scrape_logs_stat (
    chain_id character varying NOT NULL,
    last_block bigint,
    update_at timestamp without time zone
);


ALTER TABLE public.scrape_logs_stat OWNER TO postgres;

--
-- Name: sys_apis; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.sys_apis (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    path text,
    description text,
    api_group text,
    method text DEFAULT 'POST'::text
);


ALTER TABLE public.sys_apis OWNER TO postgres;

--
-- Name: COLUMN sys_apis.path; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.sys_apis.path IS 'api路径';


--
-- Name: COLUMN sys_apis.description; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.sys_apis.description IS 'api中文描述';


--
-- Name: COLUMN sys_apis.api_group; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.sys_apis.api_group IS 'api组';


--
-- Name: COLUMN sys_apis.method; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.sys_apis.method IS '方法';


--
-- Name: sys_apis_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.sys_apis_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.sys_apis_id_seq OWNER TO postgres;

--
-- Name: sys_apis_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.sys_apis_id_seq OWNED BY public.sys_apis.id;


--
-- Name: sys_authorities; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.sys_authorities (
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    authority_id bigint NOT NULL,
    authority_name text,
    parent_id bigint,
    default_router text DEFAULT 'dashboard'::text
);


ALTER TABLE public.sys_authorities OWNER TO postgres;

--
-- Name: COLUMN sys_authorities.authority_id; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.sys_authorities.authority_id IS '角色ID';


--
-- Name: COLUMN sys_authorities.authority_name; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.sys_authorities.authority_name IS '角色名';


--
-- Name: COLUMN sys_authorities.parent_id; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.sys_authorities.parent_id IS '父角色ID';


--
-- Name: COLUMN sys_authorities.default_router; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.sys_authorities.default_router IS '默认菜单';


--
-- Name: sys_authorities_authority_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.sys_authorities_authority_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.sys_authorities_authority_id_seq OWNER TO postgres;

--
-- Name: sys_authorities_authority_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.sys_authorities_authority_id_seq OWNED BY public.sys_authorities.authority_id;


--
-- Name: sys_authority_btns; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.sys_authority_btns (
    authority_id bigint,
    sys_menu_id bigint,
    sys_base_menu_btn_id bigint
);


ALTER TABLE public.sys_authority_btns OWNER TO postgres;

--
-- Name: COLUMN sys_authority_btns.authority_id; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.sys_authority_btns.authority_id IS '角色ID';


--
-- Name: COLUMN sys_authority_btns.sys_menu_id; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.sys_authority_btns.sys_menu_id IS '菜单ID';


--
-- Name: COLUMN sys_authority_btns.sys_base_menu_btn_id; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.sys_authority_btns.sys_base_menu_btn_id IS '菜单按钮ID';


--
-- Name: sys_authority_menus; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.sys_authority_menus (
    sys_base_menu_id bigint NOT NULL,
    sys_authority_authority_id bigint NOT NULL
);


ALTER TABLE public.sys_authority_menus OWNER TO postgres;

--
-- Name: COLUMN sys_authority_menus.sys_authority_authority_id; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.sys_authority_menus.sys_authority_authority_id IS '角色ID';


--
-- Name: sys_auto_code_histories; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.sys_auto_code_histories (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    package text,
    business_db text,
    table_name text,
    request_meta text,
    auto_code_path text,
    injection_meta text,
    struct_name text,
    struct_cn_name text,
    api_ids text,
    flag bigint
);


ALTER TABLE public.sys_auto_code_histories OWNER TO postgres;

--
-- Name: sys_auto_code_histories_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.sys_auto_code_histories_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.sys_auto_code_histories_id_seq OWNER TO postgres;

--
-- Name: sys_auto_code_histories_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.sys_auto_code_histories_id_seq OWNED BY public.sys_auto_code_histories.id;


--
-- Name: sys_auto_codes; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.sys_auto_codes (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    package_name text,
    label text,
    "desc" text
);


ALTER TABLE public.sys_auto_codes OWNER TO postgres;

--
-- Name: COLUMN sys_auto_codes.package_name; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.sys_auto_codes.package_name IS '包名';


--
-- Name: COLUMN sys_auto_codes.label; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.sys_auto_codes.label IS '展示名';


--
-- Name: COLUMN sys_auto_codes."desc"; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.sys_auto_codes."desc" IS '描述';


--
-- Name: sys_auto_codes_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.sys_auto_codes_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.sys_auto_codes_id_seq OWNER TO postgres;

--
-- Name: sys_auto_codes_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.sys_auto_codes_id_seq OWNED BY public.sys_auto_codes.id;


--
-- Name: sys_base_menu_btns; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.sys_base_menu_btns (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    name text,
    "desc" text,
    sys_base_menu_id bigint
);


ALTER TABLE public.sys_base_menu_btns OWNER TO postgres;

--
-- Name: COLUMN sys_base_menu_btns.name; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.sys_base_menu_btns.name IS '按钮关键key';


--
-- Name: COLUMN sys_base_menu_btns.sys_base_menu_id; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.sys_base_menu_btns.sys_base_menu_id IS '菜单ID';


--
-- Name: sys_base_menu_btns_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.sys_base_menu_btns_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.sys_base_menu_btns_id_seq OWNER TO postgres;

--
-- Name: sys_base_menu_btns_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.sys_base_menu_btns_id_seq OWNED BY public.sys_base_menu_btns.id;


--
-- Name: sys_base_menu_parameters; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.sys_base_menu_parameters (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    sys_base_menu_id bigint,
    type text,
    key text,
    value text
);


ALTER TABLE public.sys_base_menu_parameters OWNER TO postgres;

--
-- Name: COLUMN sys_base_menu_parameters.type; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.sys_base_menu_parameters.type IS '地址栏携带参数为params还是query';


--
-- Name: COLUMN sys_base_menu_parameters.key; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.sys_base_menu_parameters.key IS '地址栏携带参数的key';


--
-- Name: COLUMN sys_base_menu_parameters.value; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.sys_base_menu_parameters.value IS '地址栏携带参数的值';


--
-- Name: sys_base_menu_parameters_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.sys_base_menu_parameters_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.sys_base_menu_parameters_id_seq OWNER TO postgres;

--
-- Name: sys_base_menu_parameters_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.sys_base_menu_parameters_id_seq OWNED BY public.sys_base_menu_parameters.id;


--
-- Name: sys_base_menus; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.sys_base_menus (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    menu_level bigint,
    parent_id text,
    path text,
    name text,
    hidden boolean,
    component text,
    sort bigint,
    active_name text,
    keep_alive boolean,
    default_menu boolean,
    title text,
    icon text,
    close_tab boolean
);


ALTER TABLE public.sys_base_menus OWNER TO postgres;

--
-- Name: COLUMN sys_base_menus.parent_id; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.sys_base_menus.parent_id IS '父菜单ID';


--
-- Name: COLUMN sys_base_menus.path; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.sys_base_menus.path IS '路由path';


--
-- Name: COLUMN sys_base_menus.name; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.sys_base_menus.name IS '路由name';


--
-- Name: COLUMN sys_base_menus.hidden; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.sys_base_menus.hidden IS '是否在列表隐藏';


--
-- Name: COLUMN sys_base_menus.component; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.sys_base_menus.component IS '对应前端文件路径';


--
-- Name: COLUMN sys_base_menus.sort; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.sys_base_menus.sort IS '排序标记';


--
-- Name: COLUMN sys_base_menus.active_name; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.sys_base_menus.active_name IS '高亮菜单';


--
-- Name: COLUMN sys_base_menus.keep_alive; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.sys_base_menus.keep_alive IS '是否缓存';


--
-- Name: COLUMN sys_base_menus.default_menu; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.sys_base_menus.default_menu IS '是否是基础路由（开发中）';


--
-- Name: COLUMN sys_base_menus.title; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.sys_base_menus.title IS '菜单名';


--
-- Name: COLUMN sys_base_menus.icon; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.sys_base_menus.icon IS '菜单图标';


--
-- Name: COLUMN sys_base_menus.close_tab; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.sys_base_menus.close_tab IS '自动关闭tab';


--
-- Name: sys_base_menus_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.sys_base_menus_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.sys_base_menus_id_seq OWNER TO postgres;

--
-- Name: sys_base_menus_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.sys_base_menus_id_seq OWNED BY public.sys_base_menus.id;


--
-- Name: sys_chat_gpt_options; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.sys_chat_gpt_options (
    sk text
);


ALTER TABLE public.sys_chat_gpt_options OWNER TO postgres;

--
-- Name: sys_data_authority_id; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.sys_data_authority_id (
    sys_authority_authority_id bigint NOT NULL,
    data_authority_id_authority_id bigint NOT NULL
);


ALTER TABLE public.sys_data_authority_id OWNER TO postgres;

--
-- Name: COLUMN sys_data_authority_id.sys_authority_authority_id; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.sys_data_authority_id.sys_authority_authority_id IS '角色ID';


--
-- Name: COLUMN sys_data_authority_id.data_authority_id_authority_id; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.sys_data_authority_id.data_authority_id_authority_id IS '角色ID';


--
-- Name: sys_dictionaries; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.sys_dictionaries (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    name text,
    type text,
    status boolean,
    "desc" text
);


ALTER TABLE public.sys_dictionaries OWNER TO postgres;

--
-- Name: COLUMN sys_dictionaries.name; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.sys_dictionaries.name IS '字典名（中）';


--
-- Name: COLUMN sys_dictionaries.type; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.sys_dictionaries.type IS '字典名（英）';


--
-- Name: COLUMN sys_dictionaries.status; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.sys_dictionaries.status IS '状态';


--
-- Name: COLUMN sys_dictionaries."desc"; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.sys_dictionaries."desc" IS '描述';


--
-- Name: sys_dictionaries_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.sys_dictionaries_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.sys_dictionaries_id_seq OWNER TO postgres;

--
-- Name: sys_dictionaries_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.sys_dictionaries_id_seq OWNED BY public.sys_dictionaries.id;


--
-- Name: sys_dictionary_details; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.sys_dictionary_details (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    label text,
    value bigint,
    status boolean,
    sort bigint,
    sys_dictionary_id bigint
);


ALTER TABLE public.sys_dictionary_details OWNER TO postgres;

--
-- Name: COLUMN sys_dictionary_details.label; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.sys_dictionary_details.label IS '展示值';


--
-- Name: COLUMN sys_dictionary_details.value; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.sys_dictionary_details.value IS '字典值';


--
-- Name: COLUMN sys_dictionary_details.status; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.sys_dictionary_details.status IS '启用状态';


--
-- Name: COLUMN sys_dictionary_details.sort; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.sys_dictionary_details.sort IS '排序标记';


--
-- Name: COLUMN sys_dictionary_details.sys_dictionary_id; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.sys_dictionary_details.sys_dictionary_id IS '关联标记';


--
-- Name: sys_dictionary_details_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.sys_dictionary_details_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.sys_dictionary_details_id_seq OWNER TO postgres;

--
-- Name: sys_dictionary_details_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.sys_dictionary_details_id_seq OWNED BY public.sys_dictionary_details.id;


--
-- Name: sys_operation_records; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.sys_operation_records (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    ip text,
    method text,
    path text,
    status bigint,
    latency bigint,
    agent text,
    error_message text,
    body text,
    resp text,
    user_id bigint
);


ALTER TABLE public.sys_operation_records OWNER TO postgres;

--
-- Name: COLUMN sys_operation_records.ip; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.sys_operation_records.ip IS '请求ip';


--
-- Name: COLUMN sys_operation_records.method; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.sys_operation_records.method IS '请求方法';


--
-- Name: COLUMN sys_operation_records.path; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.sys_operation_records.path IS '请求路径';


--
-- Name: COLUMN sys_operation_records.status; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.sys_operation_records.status IS '请求状态';


--
-- Name: COLUMN sys_operation_records.latency; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.sys_operation_records.latency IS '延迟';


--
-- Name: COLUMN sys_operation_records.agent; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.sys_operation_records.agent IS '代理';


--
-- Name: COLUMN sys_operation_records.error_message; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.sys_operation_records.error_message IS '错误信息';


--
-- Name: COLUMN sys_operation_records.body; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.sys_operation_records.body IS '请求Body';


--
-- Name: COLUMN sys_operation_records.resp; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.sys_operation_records.resp IS '响应Body';


--
-- Name: COLUMN sys_operation_records.user_id; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.sys_operation_records.user_id IS '用户id';


--
-- Name: sys_operation_records_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.sys_operation_records_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.sys_operation_records_id_seq OWNER TO postgres;

--
-- Name: sys_operation_records_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.sys_operation_records_id_seq OWNED BY public.sys_operation_records.id;


--
-- Name: sys_user_authority; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.sys_user_authority (
    sys_user_id bigint NOT NULL,
    sys_authority_authority_id bigint NOT NULL
);


ALTER TABLE public.sys_user_authority OWNER TO postgres;

--
-- Name: COLUMN sys_user_authority.sys_authority_authority_id; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.sys_user_authority.sys_authority_authority_id IS '角色ID';


--
-- Name: sys_users; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.sys_users (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    uuid text,
    username text,
    password text,
    nick_name text DEFAULT '系统用户'::text,
    side_mode text DEFAULT 'dark'::text,
    header_img text DEFAULT 'https://qmplusimg.henrongyi.top/gva_header.jpg'::text,
    base_color text DEFAULT '#fff'::text,
    active_color text DEFAULT '#1890ff'::text,
    authority_id bigint DEFAULT 888,
    phone text,
    email text,
    enable bigint DEFAULT 1
);


ALTER TABLE public.sys_users OWNER TO postgres;

--
-- Name: COLUMN sys_users.uuid; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.sys_users.uuid IS '用户UUID';


--
-- Name: COLUMN sys_users.username; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.sys_users.username IS '用户登录名';


--
-- Name: COLUMN sys_users.password; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.sys_users.password IS '用户登录密码';


--
-- Name: COLUMN sys_users.nick_name; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.sys_users.nick_name IS '用户昵称';


--
-- Name: COLUMN sys_users.side_mode; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.sys_users.side_mode IS '用户侧边主题';


--
-- Name: COLUMN sys_users.header_img; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.sys_users.header_img IS '用户头像';


--
-- Name: COLUMN sys_users.base_color; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.sys_users.base_color IS '基础颜色';


--
-- Name: COLUMN sys_users.active_color; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.sys_users.active_color IS '活跃颜色';


--
-- Name: COLUMN sys_users.authority_id; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.sys_users.authority_id IS '用户角色ID';


--
-- Name: COLUMN sys_users.phone; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.sys_users.phone IS '用户手机号';


--
-- Name: COLUMN sys_users.email; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.sys_users.email IS '用户邮箱';


--
-- Name: COLUMN sys_users.enable; Type: COMMENT; Schema: public; Owner: postgres
--

COMMENT ON COLUMN public.sys_users.enable IS '用户是否被冻结 1正常 2冻结';


--
-- Name: sys_users_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.sys_users_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.sys_users_id_seq OWNER TO postgres;

--
-- Name: sys_users_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.sys_users_id_seq OWNED BY public.sys_users.id;


--
-- Name: tfa; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.tfa (
    user_id character varying(255) NOT NULL,
    created_at timestamp(6) without time zone,
    deleted_at timestamp(6) without time zone,
    phone character varying(255),
    mail character varying(255),
    phone_updated_at timestamp(6) without time zone,
    mail_updated_at timestamp(6) without time zone
);


ALTER TABLE public.tfa OWNER TO postgres;

--
-- Name: casbin_rule id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.casbin_rule ALTER COLUMN id SET DEFAULT nextval('public.casbin_rule_id_seq'::regclass);


--
-- Name: contract_abi id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.contract_abi ALTER COLUMN id SET DEFAULT nextval('public.contract_abi_id_seq'::regclass);


--
-- Name: exa_customers id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.exa_customers ALTER COLUMN id SET DEFAULT nextval('public.exa_customers_id_seq'::regclass);


--
-- Name: exa_file_chunks id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.exa_file_chunks ALTER COLUMN id SET DEFAULT nextval('public.exa_file_chunks_id_seq'::regclass);


--
-- Name: exa_file_upload_and_downloads id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.exa_file_upload_and_downloads ALTER COLUMN id SET DEFAULT nextval('public.exa_file_upload_and_downloads_id_seq'::regclass);


--
-- Name: exa_files id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.exa_files ALTER COLUMN id SET DEFAULT nextval('public.exa_files_id_seq'::regclass);


--
-- Name: jwt_blacklists id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.jwt_blacklists ALTER COLUMN id SET DEFAULT nextval('public.jwt_blacklists_id_seq'::regclass);


--
-- Name: rule id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.rule ALTER COLUMN id SET DEFAULT nextval('public.rule_id_seq'::regclass);


--
-- Name: sys_apis id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.sys_apis ALTER COLUMN id SET DEFAULT nextval('public.sys_apis_id_seq'::regclass);


--
-- Name: sys_authorities authority_id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.sys_authorities ALTER COLUMN authority_id SET DEFAULT nextval('public.sys_authorities_authority_id_seq'::regclass);


--
-- Name: sys_auto_code_histories id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.sys_auto_code_histories ALTER COLUMN id SET DEFAULT nextval('public.sys_auto_code_histories_id_seq'::regclass);


--
-- Name: sys_auto_codes id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.sys_auto_codes ALTER COLUMN id SET DEFAULT nextval('public.sys_auto_codes_id_seq'::regclass);


--
-- Name: sys_base_menu_btns id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.sys_base_menu_btns ALTER COLUMN id SET DEFAULT nextval('public.sys_base_menu_btns_id_seq'::regclass);


--
-- Name: sys_base_menu_parameters id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.sys_base_menu_parameters ALTER COLUMN id SET DEFAULT nextval('public.sys_base_menu_parameters_id_seq'::regclass);


--
-- Name: sys_base_menus id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.sys_base_menus ALTER COLUMN id SET DEFAULT nextval('public.sys_base_menus_id_seq'::regclass);


--
-- Name: sys_dictionaries id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.sys_dictionaries ALTER COLUMN id SET DEFAULT nextval('public.sys_dictionaries_id_seq'::regclass);


--
-- Name: sys_dictionary_details id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.sys_dictionary_details ALTER COLUMN id SET DEFAULT nextval('public.sys_dictionary_details_id_seq'::regclass);


--
-- Name: sys_operation_records id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.sys_operation_records ALTER COLUMN id SET DEFAULT nextval('public.sys_operation_records_id_seq'::regclass);


--
-- Name: sys_users id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.sys_users ALTER COLUMN id SET DEFAULT nextval('public.sys_users_id_seq'::regclass);


--
-- Name: casbin_rule casbin_rule_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.casbin_rule
    ADD CONSTRAINT casbin_rule_pkey PRIMARY KEY (id);


--
-- Name: contract_abi contract_abi_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.contract_abi
    ADD CONSTRAINT contract_abi_pkey PRIMARY KEY (id);


--
-- Name: eth_tx eth_tx_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.eth_tx
    ADD CONSTRAINT eth_tx_pkey PRIMARY KEY (id);


--
-- Name: exa_customers exa_customers_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.exa_customers
    ADD CONSTRAINT exa_customers_pkey PRIMARY KEY (id);


--
-- Name: exa_file_chunks exa_file_chunks_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.exa_file_chunks
    ADD CONSTRAINT exa_file_chunks_pkey PRIMARY KEY (id);


--
-- Name: exa_file_upload_and_downloads exa_file_upload_and_downloads_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.exa_file_upload_and_downloads
    ADD CONSTRAINT exa_file_upload_and_downloads_pkey PRIMARY KEY (id);


--
-- Name: exa_files exa_files_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.exa_files
    ADD CONSTRAINT exa_files_pkey PRIMARY KEY (id);


--
-- Name: jwt_blacklists jwt_blacklists_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.jwt_blacklists
    ADD CONSTRAINT jwt_blacklists_pkey PRIMARY KEY (id);


--
-- Name: mpc_context mpc_context_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.mpc_context
    ADD CONSTRAINT mpc_context_pkey PRIMARY KEY (user_id);


--
-- Name: rule rule_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.rule
    ADD CONSTRAINT rule_pkey PRIMARY KEY (id);


--
-- Name: scrape_logs_stat scrape_logs_stat_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.scrape_logs_stat
    ADD CONSTRAINT scrape_logs_stat_pkey PRIMARY KEY (chain_id);


--
-- Name: sys_apis sys_apis_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.sys_apis
    ADD CONSTRAINT sys_apis_pkey PRIMARY KEY (id);


--
-- Name: sys_authorities sys_authorities_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.sys_authorities
    ADD CONSTRAINT sys_authorities_pkey PRIMARY KEY (authority_id);


--
-- Name: sys_authority_menus sys_authority_menus_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.sys_authority_menus
    ADD CONSTRAINT sys_authority_menus_pkey PRIMARY KEY (sys_base_menu_id, sys_authority_authority_id);


--
-- Name: sys_auto_code_histories sys_auto_code_histories_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.sys_auto_code_histories
    ADD CONSTRAINT sys_auto_code_histories_pkey PRIMARY KEY (id);


--
-- Name: sys_auto_codes sys_auto_codes_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.sys_auto_codes
    ADD CONSTRAINT sys_auto_codes_pkey PRIMARY KEY (id);


--
-- Name: sys_base_menu_btns sys_base_menu_btns_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.sys_base_menu_btns
    ADD CONSTRAINT sys_base_menu_btns_pkey PRIMARY KEY (id);


--
-- Name: sys_base_menu_parameters sys_base_menu_parameters_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.sys_base_menu_parameters
    ADD CONSTRAINT sys_base_menu_parameters_pkey PRIMARY KEY (id);


--
-- Name: sys_base_menus sys_base_menus_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.sys_base_menus
    ADD CONSTRAINT sys_base_menus_pkey PRIMARY KEY (id);


--
-- Name: sys_data_authority_id sys_data_authority_id_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.sys_data_authority_id
    ADD CONSTRAINT sys_data_authority_id_pkey PRIMARY KEY (sys_authority_authority_id, data_authority_id_authority_id);


--
-- Name: sys_dictionaries sys_dictionaries_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.sys_dictionaries
    ADD CONSTRAINT sys_dictionaries_pkey PRIMARY KEY (id);


--
-- Name: sys_dictionary_details sys_dictionary_details_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.sys_dictionary_details
    ADD CONSTRAINT sys_dictionary_details_pkey PRIMARY KEY (id);


--
-- Name: sys_operation_records sys_operation_records_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.sys_operation_records
    ADD CONSTRAINT sys_operation_records_pkey PRIMARY KEY (id);


--
-- Name: sys_user_authority sys_user_authority_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.sys_user_authority
    ADD CONSTRAINT sys_user_authority_pkey PRIMARY KEY (sys_user_id, sys_authority_authority_id);


--
-- Name: sys_users sys_users_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.sys_users
    ADD CONSTRAINT sys_users_pkey PRIMARY KEY (id);


--
-- Name: tfa tfa_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.tfa
    ADD CONSTRAINT tfa_pkey PRIMARY KEY (user_id);


--
-- Name: block_number; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX block_number ON public.eth_tx USING btree (block_number);


--
-- Name: fromBlock; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX "fromBlock" ON public.agg_ft_24h USING btree ("fromBlock");


--
-- Name: fromcontracttmethod; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX fromcontracttmethod ON public.agg_ft_24h USING btree ("from", contract, method_name);


--
-- Name: idx_casbin_rule; Type: INDEX; Schema: public; Owner: postgres
--

CREATE UNIQUE INDEX idx_casbin_rule ON public.casbin_rule USING btree (ptype, v0, v1, v2, v3, v4, v5);


--
-- Name: idx_contract_abi_deleted_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_contract_abi_deleted_at ON public.contract_abi USING btree (deleted_at);


--
-- Name: idx_eth_tx_deleted_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_eth_tx_deleted_at ON public.eth_tx USING btree (deleted_at);


--
-- Name: idx_exa_customers_deleted_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_exa_customers_deleted_at ON public.exa_customers USING btree (deleted_at);


--
-- Name: idx_exa_file_chunks_deleted_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_exa_file_chunks_deleted_at ON public.exa_file_chunks USING btree (deleted_at);


--
-- Name: idx_exa_file_upload_and_downloads_deleted_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_exa_file_upload_and_downloads_deleted_at ON public.exa_file_upload_and_downloads USING btree (deleted_at);


--
-- Name: idx_exa_files_deleted_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_exa_files_deleted_at ON public.exa_files USING btree (deleted_at);


--
-- Name: idx_jwt_blacklists_deleted_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_jwt_blacklists_deleted_at ON public.jwt_blacklists USING btree (deleted_at);


--
-- Name: idx_rule_deleted_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_rule_deleted_at ON public.rule USING btree (deleted_at);


--
-- Name: idx_sys_apis_deleted_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_sys_apis_deleted_at ON public.sys_apis USING btree (deleted_at);


--
-- Name: idx_sys_auto_code_histories_deleted_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_sys_auto_code_histories_deleted_at ON public.sys_auto_code_histories USING btree (deleted_at);


--
-- Name: idx_sys_auto_codes_deleted_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_sys_auto_codes_deleted_at ON public.sys_auto_codes USING btree (deleted_at);


--
-- Name: idx_sys_base_menu_btns_deleted_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_sys_base_menu_btns_deleted_at ON public.sys_base_menu_btns USING btree (deleted_at);


--
-- Name: idx_sys_base_menu_parameters_deleted_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_sys_base_menu_parameters_deleted_at ON public.sys_base_menu_parameters USING btree (deleted_at);


--
-- Name: idx_sys_base_menus_deleted_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_sys_base_menus_deleted_at ON public.sys_base_menus USING btree (deleted_at);


--
-- Name: idx_sys_dictionaries_deleted_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_sys_dictionaries_deleted_at ON public.sys_dictionaries USING btree (deleted_at);


--
-- Name: idx_sys_dictionary_details_deleted_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_sys_dictionary_details_deleted_at ON public.sys_dictionary_details USING btree (deleted_at);


--
-- Name: idx_sys_operation_records_deleted_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_sys_operation_records_deleted_at ON public.sys_operation_records USING btree (deleted_at);


--
-- Name: idx_sys_users_deleted_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_sys_users_deleted_at ON public.sys_users USING btree (deleted_at);


--
-- Name: idx_sys_users_username; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_sys_users_username ON public.sys_users USING btree (username);


--
-- Name: idx_sys_users_uuid; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_sys_users_uuid ON public.sys_users USING btree (uuid);


--
-- Name: nft_fromBlock; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX "nft_fromBlock" ON public.agg_nft_24h USING btree ("fromBlock");


--
-- Name: nft_fromcontractmethod; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX nft_fromcontractmethod ON public.agg_nft_24h USING btree ("from", contract, method_name);


--
-- Name: nft_toBlock; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX "nft_toBlock" ON public.agg_nft_24h USING btree ("toBlock");


--
-- Name: nrfromcontractmethod; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX nrfromcontractmethod ON public.eth_tx USING btree (method_name, "from", block_number, contract);


--
-- Name: toBlock; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX "toBlock" ON public.agg_ft_24h USING btree ("toBlock");


--
-- Name: rule notify_update_trigger_rule; Type: TRIGGER; Schema: public; Owner: postgres
--

CREATE TRIGGER notify_update_trigger_rule AFTER INSERT OR DELETE OR UPDATE OF rule_id, rule_desc, rules ON public.rule FOR EACH ROW EXECUTE FUNCTION public.notify_update();


--
-- PostgreSQL database dump complete
--

