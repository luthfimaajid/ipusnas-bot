-- migrate:up
CREATE TABLE ipusnas_account (
	ipusnas_account_id INTEGER PRIMARY KEY,
	email VARCHAR(255) NOT NULL,
	password VARCHAR(255) NOT NULL,
	access_token VARCHAR(255),
	access_token_expiry DATE
);

CREATE TABLE book (
	book_id INTEGER PRIMARY KEY,
	title VARCHAR(255) NOT NULL,
	author VARCHAR(255) NOT NULL,
	cover_url VARCHAR(255),
	available_copy INTEGER NOT NULL
);

CREATE TABLE loan (
	load_id INTEGER PRIMARY KEY,
	book_id INTEGER NOT NULL,
	ipusnas_account_id INTEGER NOT NULL,
	loan_date DATE,
	remaining_days INTEGER,
	is_queueing BOOLEAN,
	FOREIGN KEY(book_id) REFERENCES book(book_id),
	FOREIGN KEY(ipusnas_account_id) REFERENCES ipusnas_account(ipusnas_account_id)
);

-- migrate:down
DROP TABLE ipusnas_account;
DROP TABLE book;
DROP TABLE loan;

