package main

import (
	"log"
	"time"
)

type Book struct {
	id         int
	isBorrowed bool
	title      string
}

func main() {
	cfg := LoadEnv()

	adapter := Adapter{
		cfg: cfg,
	}

	books := []Book{}

	for i := 0; i < len(cfg.Targets); i++ {
		resp := adapter.BookDetail(cfg.Targets[i])

		books = append(books, Book{
			id:         cfg.Targets[i],
			isBorrowed: false,
			title:      resp.Data.Book.Title,
		})
	}

	adapter.Login()

	count := 0
	for range time.Tick(time.Second * 2) {
		for i := 0; i < len(books); i++ {
			if books[i].isBorrowed {
				continue
			}

			res := adapter.BorrowBook(books[i].id, 0)
			if res.Meta.ErrorMessage == BorrowBookNeedConfirm {
				res2 := adapter.BorrowBook(books[i].id, 1)
				if res2.Meta.Confirm == BorrowBookSuccess {
					books[i].isBorrowed = true
					log.Printf("%s berhasil dipinjam!", books[i].title)
					count++
				}
			} else {
				log.Printf("%s tidak bisa dipinjam karena: %s", books[i].title, res.Meta.ErrorMessage)
			}
		}

		if count == len(books) {
			log.Printf("Semua buku berhasil dipinjam!")
			break
		}
	}
}
