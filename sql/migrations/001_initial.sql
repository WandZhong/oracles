CREATE TABLE swc_queue (
	swc_queue_id  UUID NOT NULL PRIMARY KEY,
	user_id       UUID NOT NULL,
	wad           DECIMAL CHECK (wad > 0),
	currency      CHAR[3] NOT NULL,
	created_on    TIMESTAMP NOT NULL
);

CREATE TABLE swc_queue_tranche (
	swc_queue_tranche_id  SERIAL PRIMARY KEY,
	created_on            TIMESTAMP NOT NULL,
	-- TODO: how to handle different flavours (brg*)
	price_brg_usd         DECIMAL CHECK (price_brg_usd > 0)
);

CREATE TABLE swc_queue_distribution_list (
	swc_queue_tranche_id  SERIAL PRIMARY KEY,
	user_id               UUID NOT NULL,
	wad                   DECIMAL CHECK (wad > 0)
);
