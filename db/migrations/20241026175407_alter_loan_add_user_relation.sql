-- migrate:up
DROP TABLE loan;
CREATE TABLE loan (
	loan_id INTEGER PRIMARY KEY,
	book_id INTEGER NOT NULL,
	ipusnas_account_id INTEGER NOT NULL,
	user_id INTEGER NOT NULL,
	loan_date DATE,
	remaining_days INTEGER,
	is_queueing BOOLEAN,
	FOREIGN KEY(book_id) REFERENCES book(book_id),
	FOREIGN KEY(ipusnas_account_id) REFERENCES ipusnas_account(ipusnas_account_id),
	FOREIGN KEY(user_id) REFERENCES "user"(user_id)
);



-- migrate:down

