CREATE TABLE IF NOT EXISTS "schema_migrations" (version varchar(128) primary key);
CREATE TABLE book (
	book_id INTEGER PRIMARY KEY,
	title VARCHAR(255) NOT NULL,
	author VARCHAR(255) NOT NULL,
	cover_url VARCHAR(255),
	available_copy INTEGER NOT NULL
);
CREATE TABLE IF NOT EXISTS "ipusnas_account" (
	ipusnas_account_id INTEGER PRIMARY KEY,
	email VARCHAR(255) NOT NULL,
	password VARCHAR(255) NOT NULL,
	access_token VARCHAR(255),
	access_token_expiry INTEGER
);
CREATE TABLE IF NOT EXISTS "user" (
	user_id INTEGER PRIMARY KEY,
	email VARCHAR(255) NOT NULL,
	password VARCHAR(255) NOT NULL
);
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
-- Dbmate schema migrations
INSERT INTO "schema_migrations" (version) VALUES
  ('20241004081903'),
  ('20241004090723'),
  ('20241006150819'),
  ('20241008090949'),
  ('20241026175300'),
  ('20241026175407'),
  ('20241026180438');
