-- migrate:up
CREATE TABLE "user" (
	user_id INTEGER PRIMARY KEY,
	email VARCHAR(255) NOT NULL,
	password VARCHAR(255) NOT NULL
);



-- migrate:down

