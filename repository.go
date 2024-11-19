package main

import (
	"context"
	"time"

	"github.com/jmoiron/sqlx"
)

type repository struct {
	db *sqlx.DB
}

func NewRepository(d *sqlx.DB) *repository {
	return &repository{
		db: d,
	}
}

func (r *repository) Insert(ctx context.Context, account IpusnasAccount) error {
	queryStr := `INSERT INTO ipusnas_account VALUES(NULL, :email, :password, :access_token, :access_token_expiry);`

	_, err := r.db.NamedExecContext(ctx, queryStr, account)

	return err
}

// fetch account with the remaining token expiration < 24 hour
func (r *repository) GetAlmostExpired(ctx context.Context) ([]IpusnasAccount, error) {
	queryStr := `SELECT * FROM ipusnas_account WHERE access_token_expiry < ?1 OR access_token IS NULL;`
	accounts := []IpusnasAccount{}
	now := time.Now().AddDate(0, 0, +1).Unix()

	err := r.db.SelectContext(ctx, &accounts, queryStr, now)
	if err != nil {
		return nil, err
	}

	return accounts, nil
}

func (r *repository) UpdateAccessToken(ctx context.Context, accounts []IpusnasAccount) error {
	// CANNOT WORK IN SQLITE3
	/*
		queryStr := `UPDATE ipusnas_account
		SET
			access_token = new_val.token,
			access_token_expiry = new_val.expiry
		FROM (VALUES
				(:ipusnas_account_id, :access_token, :access_token_expiry)
		) new_val (id, token, expiry)
		WHERE ipusnas_account.ipusnas_account_id = new_val.id;`
	*/

	queryStr := `UPDATE ipusnas_account
	SET
		access_token = :access_token,
		access_token_expiry = :access_token_expiry
	WHERE ipusnas_account_id = :ipusnas_account_id;
	`

	for _, a := range accounts {
		_, err := r.db.NamedExecContext(ctx, queryStr, a)
		if err != nil {
			return err
		}
	}

	return nil
}
