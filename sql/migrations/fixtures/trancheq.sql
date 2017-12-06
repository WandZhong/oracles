INSERT INTO tokens VALUES
  ('SWC', '2017-11-01', '', NULL),
  ('MET', now(), '', NULL);

INSERT INTO tranches VALUES
	('1', 'tge1', 'SWC', now(), '2017-12-01', '2017-12-14', now() + '20 days', 100000000000000000000000, 'Scott', 10000000000000000000000);

INSERT INTO tranche_prices VALUES
	('1', 'USD', 2.4);
