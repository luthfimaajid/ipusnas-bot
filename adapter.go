package main

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"time"
)

type adapter struct {
	cfg *Config
}

func encode(src interface{}) (dst *bytes.Buffer, err error) {
	dst = &bytes.Buffer{}
	err = json.NewEncoder(dst).Encode(src)
	return
}

func decode[T any](src io.Reader) (dst *T, err error) {
	dst = new(T)
	err = json.NewDecoder(src).Decode(dst)
	return
}

func NewIpusnasAPI(c *Config) *adapter {
	return &adapter{
		cfg: c,
	}
}

func (a *adapter) Login(email string, password string) (account *IpusnasAccount, err error) {
	account = new(IpusnasAccount)
	login := LoginRequest{
		Username:     email,
		Password:     password,
		ClientId:     a.cfg.Client.Id,
		ClientSecret: a.cfg.Client.Secret,
		DeviceId:     a.cfg.Client.DeviceId,
	}

	reqBody, err := encode(login)
	if err != nil {
		return
	}

	res, err := http.Post(a.cfg.Ipusnas.Url+a.cfg.Ipusnas.Api.Login, ContentTypeJSON, reqBody)
	if err != nil {
		return
	}

	resBody, err := decode[LoginResponse](res.Body)
	if err != nil {
		return
	}

	account.AccessToken = &resBody.Data.AccessToken
	expDate, err := time.Parse("2006-01-02", resBody.Data.Expired)
	exp := expDate.Unix()

	if err != nil {
		return
	}

	account.AccessTokenExpiry = &exp

	return
}

func (a *adapter) BorrowBook(bookId int, confirm int) (resBody *BorrowBookResponse, err error) {
	borrowBookReq := BorrowBookRequest{
		// AccessToken: a.AccessToken,
		BookId:    bookId,
		LibraryId: a.cfg.Ipusnas.LibraryId,
		Confirm:   confirm,
	}

	reqBody, err := encode(borrowBookReq)
	if err != nil {
		return
	}

	res, err := http.Post(a.cfg.Ipusnas.Url+a.cfg.Ipusnas.Api.BorrowBook, ContentTypeJSON, reqBody)
	if err != nil {
		return
	}

	resBody, err = decode[BorrowBookResponse](res.Body)

	return
}

// func (a *adapter) BookDetail(bookId int) (resBody *BookDetailResponse, err error) {
// 	url := fmt.Sprintf("%s%sclient_id=%s&book_id=%d&access_token=%s", a.cfg.Ipusnas.Url, a.cfg.Ipusnas.Api.BookDetail, a.cfg.Client.Id, bookId, a.AccessToken)
// 	res, err := http.Get(url)
// 	if err != nil {
// 		return
// 	}

// 	return decode[BookDetailResponse](res.Body)
// }
