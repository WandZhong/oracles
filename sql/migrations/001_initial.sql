CREATE TABLE swc_queue (
	swc_queue_id  VARCHAR NOT NULL PRIMARY KEY,  -- transaction hash
	user_addr     VARCHAR(50) NOT NULL,
	ctr_addr      VARCHAR(50) NOT NULL,
	wad           DECIMAL CHECK (wad > 0), -- brg
	currency      CHAR(3) NOT NULL CHECK (currency != ''),
	created_on    TIMESTAMP WITH TIME ZONE NOT NULL,
	direct        BOOL NOT NULL
);

CREATE TABLE swc_queue_tranche (
	swc_queue_tranche_id  SERIAL PRIMARY KEY,
	created_on            TIMESTAMP WITH TIME ZONE NOT NULL,
	-- TODO: how to handle different flavours (brg*)
	price_brg_usd         DECIMAL CHECK (price_brg_usd > 0)
);

CREATE TABLE swc_queue_distribution_list (
	swc_queue_tranche_id  SERIAL PRIMARY KEY,
	user_id               UUID NOT NULL,
	wad                   DECIMAL CHECK (wad > 0)
);
