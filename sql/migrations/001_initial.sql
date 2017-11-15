CREATE TABLE tranches (
	tranche_id       BIGSERIAL PRIMARY KEY,
	token            VARCHAR NOT NULL,
	created_on       TIMESTAMP WITH TIME ZONE NOT NULL,
	-- TODO: how to handle different flavours (brg*)
	price_brg_usd    DECIMAL CHECK (price_brg_usd > 0)
);

CREATE TABLE pledge_queue (
	pledge_queue_id    VARCHAR PRIMARY KEY,  -- transaction hash
	created_on         TIMESTAMP WITH TIME ZONE NOT NULL,
	wallet_addr        VARCHAR(50) NOT NULL,
	ctr_addr           VARCHAR(50) NOT NULL,
	wad                DECIMAL CHECK (wad > 0),  -- brg*
	wad_withdrawed     DECIMAL CHECK (wad >= 0), -- brg*
	wad_executed       DECIMAL CHECK (wad >= 0), -- brg* used for tranche distribution. This denormalize the model a bit, so maybe it will be removed.
	currency           CHAR(3) NOT NULL CHECK (currency != ''),
	direct             BOOL NOT NULL
);

CREATE TABLE tranche_subscriptions (
	tranche_subscription_id  UUID PRIMARY KEY,
	tranche_id               BIGINT NOT NULL REFERENCES tranches,
	pledge_queue_id          VARCHAR NOT NULL REFERENCES pledge_queue,
	created_on               TIMESTAMP WITH TIME ZONE NOT NULL,
	wad                      DECIMAL CHECK (wad > 0),
	executed                 BOOL NOT NULL
);

CREATE TABLE tranche_distributions (
	tranche_distribution_id  BIGSERIAL PRIMARY KEY,
	tranche_id               BIGINT NOT NULL REFERENCES tranches,
	wad                      DECIMAL CHECK (wad > 0),  -- new token amount
	wallet_addr              VARCHAR(50) NOT NULL,
);
