-- migrate:up
DROP TABLE loan;
CREATE TABLE loan (
	loan_id INTEGER PRIMARY KEY,
	book_id INTEGER NOT NULL,
	ipusnas_account_id INTEGER NOT NULL,
	user_id INTEGER NOT NULL,
	available_at DATE,
	status INTEGER NOT NULL,
	FOREIGN KEY(book_id) REFERENCES book(book_id),
	FOREIGN KEY(ipusnas_account_id) REFERENCES ipusnas_account(ipusnas_account_id),
	FOREIGN KEY(user_id) REFERENCES "user"(user_id)
);

-- migrate:down

