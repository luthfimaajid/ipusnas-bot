package main

type IpusnasAccount struct {
	IpusnasAccountId  int     `db:"ipusnas_account_id"`
	Email             string  `db:"email"`
	Password          string  `db:"password"`
	AccessToken       *string `db:"access_token"`
	AccessTokenExpiry *int64  `db:"access_token_expiry"`
}
