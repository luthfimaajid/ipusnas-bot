package main

type Meta struct {
	Code         int    `json:"code"`
	Confirm      string `json:"confirm"`
	Error        string `json:"error"`
	ErrorMessage string `json:"error_message"`
}

type LoginResponse struct {
	Meta Meta `json:"meta"`
	Data struct {
		AccessToken string `json:"access_token"`
		Expired     string `json:"expired"`
	} `json:"data"`
}

type BorrowBookResponse struct {
	Meta Meta `json:"meta"`
}

type BookDetailResponse struct {
	Meta Meta `json:"meta"`
	Data struct {
		Book struct {
			Title string `json:"title"`
		} `json:"Book"`
	} `json:"data"`
}
