package main

import (
	"context"
)

type repo interface {
	Insert(ctx context.Context, account IpusnasAccount) error
	GetAlmostExpired(ctx context.Context) ([]IpusnasAccount, error)
	UpdateAccessToken(ctx context.Context, accounts []IpusnasAccount) error
}

type usecase struct {
	r repo
	a *adapter
}

func NewUsecase(repo repo, ad *adapter) *usecase {
	return &usecase{
		r: repo,
		a: ad,
	}
}

func (uc *usecase) CreateNewAccount(ctx context.Context, account IpusnasAccount) error {
	return uc.r.Insert(ctx, account)
}

func (uc *usecase) RefreshAllToken(ctx context.Context) error {
	accounts, err := uc.r.GetAlmostExpired(ctx)
	if err != nil {
		return err
	}

	updates := []IpusnasAccount{}

	for _, account := range accounts {
		a, err := uc.a.Login(account.Email, account.Password)
		if err != nil {
			return err
		}

		updates = append(updates, IpusnasAccount{
			IpusnasAccountId:  account.IpusnasAccountId,
			AccessToken:       a.AccessToken,
			AccessTokenExpiry: a.AccessTokenExpiry,
		})
	}

	return uc.r.UpdateAccessToken(ctx, updates)
}

func (uc *usecase) AddLoanQueue(ctx context.Context) error {
	// fetch one throwaway account that has <5 active loan

	// create a new loan with
	// user id and book id from request
	// and throwaway id from last fetch
	return nil
}

func (uc *usecase) TryToBorrow(ctx context.Context) error {
	// fetch all active loan, join it with throwaway to get access_token
	// loop all loan and borrow the book
	return nil
}
