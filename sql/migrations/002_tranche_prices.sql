CREATE TABLE IF NOT EXISTS currencies (
	currency_id VARCHAR PRIMARY KEY, -- the ISO code
	name        VARCHAR(60),
	type        int -- crypto vs fiat... again how we want to store enums?
);

INSERT INTO currencies VALUES
	('USD', 'US Dollar', 0),
	('cETH', 'Ether', 1),
	('cBTC', 'Bitcoin', 1);

CREATE TABLE tranche_prices (
	tranche_id    BIGINT REFERENCES tranches,
	currency_id   VARCHAR REFERENCES currencies,
	price         DOUBLE PRECISION CHECK (price > 0)
);

INSERT INTO tranche_prices
	SELECT tranche_id, 'USD', price_brg_usd
	FROM tranches;

ALTER TABLE tranches DROP COLUMN price_brg_usd;
