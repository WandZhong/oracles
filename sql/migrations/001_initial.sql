CREATE TABLE tokens (
	token_id            VARCHAR PRIMARY KEY,  -- name
	created_on          TIMESTAMP WITH TIME ZONE NOT NULL,
	comment             TEXT NOT NULL,

	-- restrictions (might be extracted to separate table)
	max_total_contrib   DECIMAL CHECK (max_total_contrib > 0)  -- in wad
);

CREATE TABLE tranches (
	tranche_id       BIGSERIAL PRIMARY KEY,
	name             TEXT DEFAULT '' NOT NULL,  -- release name
	token_id         VARCHAR NOT NULL REFERENCES tokens,
	created_on       TIMESTAMP WITH TIME ZONE NOT NULL,
	starts_on        TIMESTAMP WITH TIME ZONE NOT NULL,
	executes_on      TIMESTAMP WITH TIME ZONE NOT NULL,
	ends_on          TIMESTAMP WITH TIME ZONE CHECK (ends_on > starts_on),
	supply           DECIMAL NOT NULL CHECK (supply > 0),  -- in wad
	owner            TEXT,

	max_contrib      DECIMAL CHECK (max_contrib > 0),  -- in wad
	-- TODO: find a way to handle different flavors (brg*)
	price_brg_usd    DOUBLE PRECISION CHECK (price_brg_usd > 0)
);

CREATE TABLE pledge_queue (
	pledge_queue_id    VARCHAR PRIMARY KEY,  -- transaction hash
	created_on         TIMESTAMP WITH TIME ZONE NOT NULL,
	wallet_addr        VARCHAR(50) NOT NULL,
	ctr_addr           VARCHAR(50) NOT NULL,
	wad                DECIMAL CHECK (wad > 0) NOT NULL,  -- brg*
	wad_withdrawed     DECIMAL DEFAULT 0 CHECK (wad_withdrawed >= 0) NOT NULL, -- brg*
	wad_executed       DECIMAL DEFAULT 0 CHECK (wad_executed >= 0) NOT NULL, -- brg* used for tranche distribution. This denormalize the model a bit, so maybe it will be removed.
	currency           CHAR(3) NOT NULL CHECK (currency != ''),
	direct             BOOL NOT NULL
);

-- TODO link it with pledge_queue
CREATE TABLE direct_pledges (
	pledge_queue_id    VARCHAR PRIMARY KEY,  -- same as pledge_queue
	from_curr          VARCHAR(10) NOT NULL,
	to_curr            VARCHAR(10) NOT NULL,
	rate               real
);

-- TODO finish this
CREATE TABLE withdraws (
	withdraw_id        UUID PRIMARY KEY,
	created_on         TIMESTAMP WITH TIME ZONE NOT NULL,
	wallet_addr        VARCHAR(50) NOT NULL,
	ctr_addr           VARCHAR(50) NOT NULL,
	wad                DECIMAL CHECK (wad > 0) NOT NULL,  -- brg*
)

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
