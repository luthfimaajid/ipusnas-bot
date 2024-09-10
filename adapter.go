package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type Adapter struct {
	cfg         *Config
	AccessToken string
}

func buffToStruct[T any](res *http.Response, dst T) {
	resByte, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(resByte, dst)
	if err != nil {
		log.Fatal(err)
	}
}

func structToBuff[T any](src T) (buf bytes.Buffer) {
	err := json.NewEncoder(&buf).Encode(src)
	if err != nil {
		log.Fatal(err)
	}

	return
}

func (a *Adapter) Login() {
	account := LoginRequest{
		Username:     a.cfg.Account.Email,
		Password:     a.cfg.Account.Password,
		ClientId:     a.cfg.Client.Id,
		ClientSecret: a.cfg.Client.Secret,
		DeviceId:     a.cfg.Client.DeviceId,
	}

	buf := structToBuff[LoginRequest](account)

	res, err := http.Post(a.cfg.Ipusnas.Url+a.cfg.Ipusnas.Api.Login, "application/json", &buf)
	if err != nil {
		log.Fatal(err)
	}

	var resBody LoginResponse
	buffToStruct[*LoginResponse](res, &resBody)

	a.AccessToken = resBody.Data.AccessToken
}

func (a *Adapter) BorrowBook(bookId int, confirm int) (resp BorrowBookResponse) {
	borrowBookReq := BorrowBookRequest{
		AccessToken: a.AccessToken,
		BookId:      bookId,
		LibraryId:   a.cfg.Ipusnas.LibraryId,
		Confirm:     confirm,
	}

	buf := structToBuff[BorrowBookRequest](borrowBookReq)

	res, err := http.Post(a.cfg.Ipusnas.Url+a.cfg.Ipusnas.Api.BorrowBook, "application/json", &buf)
	if err != nil {
		log.Fatal(err)
	}

	buffToStruct[*BorrowBookResponse](res, &resp)

	return
}

func (a *Adapter) BookDetail(bookId int) (resp BookDetailResponse) {
	url := fmt.Sprintf("%s%sclient_id=%s&book_id=%d&access_token=%s", a.cfg.Ipusnas.Url, a.cfg.Ipusnas.Api.BookDetail, a.cfg.Client.Id, bookId, a.AccessToken)
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	buffToStruct[*BookDetailResponse](res, &resp)

	return
}
