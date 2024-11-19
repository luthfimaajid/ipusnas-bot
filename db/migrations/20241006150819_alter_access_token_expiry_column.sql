-- migrate:up
CREATE TABLE ipusnas_account_temp (
	ipusnas_account_id INTEGER PRIMARY KEY,
	email VARCHAR(255) NOT NULL,
	password VARCHAR(255) NOT NULL,
	access_token VARCHAR(255),
	access_token_expiry_hour INTEGER
);

INSERT INTO ipusnas_account_temp(ipusnas_account_id, email, password) SELECT ipusnas_account_id, email, password FROM ipusnas_account;
DROP TABLE ipusnas_account;
ALTER TABLE ipusnas_account_temp RENAME TO ipusnas_account;

-- migrate:down

