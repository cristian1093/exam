
CREATE TABLE public.chargebacks (
	chargeback_id serial NOT NULL,
	"date" date NOT NULL,
	afilliation varchar(30) NOT NULL DEFAULT ''::character varying,
	business_name varchar(100) NOT NULL DEFAULT ''::character varying,
	date_tx date NOT NULL,
	"import" numeric(10,2) NOT NULL,
	account varchar(4) NOT NULL DEFAULT ''::character varying,
	"ref" varchar(50) NOT NULL DEFAULT ''::character varying,
	acl_type varchar(12) NOT NULL DEFAULT ''::character varying,
	clarification_type varchar(50) NOT NULL DEFAULT ''::character varying,
	autorization varchar(6) NOT NULL DEFAULT ''::character varying,
	clarification_status varchar(20) NOT NULL DEFAULT ''::character varying,
	commitment_date date NOT NULL,
	request_type varchar(20) NOT NULL DEFAULT ''::character varying,
	description varchar(200) NULL DEFAULT NULL::character varying,
	chargeback_code varchar(20) NULL DEFAULT NULL::character varying,
	filename varchar(100) NOT NULL,
	"key" varchar(3000) NULL DEFAULT NULL::character varying,
	created_at timestamp NOT NULL,
	updated_at timestamp NULL,
	CONSTRAINT chargebacks_test_pkey PRIMARY KEY (chargeback_id)
);



CREATE TABLE public.chargebacks_duplicates (
	chargeback_id serial NOT NULL,
	"date" date NOT NULL,
	afilliation varchar(30) NOT NULL DEFAULT ''::character varying,
	business_name varchar(100) NOT NULL DEFAULT ''::character varying,
	date_tx varchar(20) NOT NULL,
	"import" numeric(10,2) NOT NULL,
	account varchar(4) NOT NULL DEFAULT ''::character varying,
	"ref" varchar(50) NOT NULL DEFAULT ''::character varying,
	acl_type varchar(12) NOT NULL DEFAULT ''::character varying,
	clarification_type varchar(50) NOT NULL DEFAULT ''::character varying,
	autorization varchar(6) NOT NULL DEFAULT ''::character varying,
	clarification_status varchar(20) NOT NULL DEFAULT ''::character varying,
	commitment_date varchar(20) NOT NULL,
	request_type varchar(20) NOT NULL DEFAULT ''::character varying,
	description varchar(200) NULL DEFAULT NULL::character varying,
	chargeback_code varchar(20) NULL DEFAULT NULL::character varying,
	filename varchar(100) NOT NULL,
	CONSTRAINT chargebacks_duplicates_pkey PRIMARY KEY (chargeback_id)
);


CREATE TABLE public.chargebacks_temporal (
	chargeback_id serial NOT NULL,
	"date" date NOT NULL,
	afilliation varchar(30) NOT NULL DEFAULT ''::character varying,
	business_name varchar(100) NOT NULL DEFAULT ''::character varying,
	date_tx date NOT NULL,
	"import" numeric(10,2) NOT NULL,
	account varchar(4) NOT NULL DEFAULT ''::character varying,
	"ref" varchar(50) NOT NULL DEFAULT ''::character varying,
	acl_type varchar(12) NOT NULL DEFAULT ''::character varying,
	clarification_type varchar(50) NOT NULL DEFAULT ''::character varying,
	autorization varchar(6) NOT NULL DEFAULT ''::character varying,
	clarification_status varchar(20) NOT NULL DEFAULT ''::character varying,
	commitment_date date NOT NULL,
	request_type varchar(20) NOT NULL DEFAULT ''::character varying,
	description varchar(200) NULL DEFAULT NULL::character varying,
	chargeback_code varchar(20) NULL DEFAULT NULL::character varying,
	filename varchar(100) NOT NULL,
	CONSTRAINT chargebacks_temporal_pkey PRIMARY KEY (chargeback_id)
);



CREATE TABLE public.chargebacks_without_afilliation (
	chargeback_id serial NOT NULL,
	"date" date NOT NULL,
	afilliation varchar(30) NOT NULL DEFAULT ''::character varying,
	business_name varchar(100) NOT NULL DEFAULT ''::character varying,
	date_tx date NOT NULL,
	"import" numeric(10,2) NOT NULL,
	account varchar(4) NOT NULL DEFAULT ''::character varying,
	"ref" varchar(50) NOT NULL DEFAULT ''::character varying,
	acl_type varchar(12) NOT NULL DEFAULT ''::character varying,
	clarification_type varchar(50) NOT NULL DEFAULT ''::character varying,
	autorization varchar(6) NOT NULL DEFAULT ''::character varying,
	clarification_status varchar(20) NOT NULL DEFAULT ''::character varying,
	commitment_date date NOT NULL,
	request_type varchar(20) NOT NULL DEFAULT ''::character varying,
	description varchar(200) NULL DEFAULT NULL::character varying,
	chargeback_code varchar(20) NULL DEFAULT NULL::character varying,
	filename varchar(100) NOT NULL,
	date_ date NULL,
	CONSTRAINT chargebacks_without_afilliation_pkey PRIMARY KEY (chargeback_id)
);
