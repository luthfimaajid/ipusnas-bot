-- migrate:up
CREATE UNIQUE INDEX unique_ipusnas_account_email ON ipusnas_account(email);

-- migrate:down


