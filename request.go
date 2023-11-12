package main

type LoginRequest struct {
	Username     string `json:"username"`
	Password     string `json:"password"`
	ClientId     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	DeviceId     string `json:"device_id"`
}

type RefreshTokenRequet struct {
	AccessToken string `json:"access_token"`
}

type BorrowBookRequest struct {
	AccessToken string `json:"access_token"`
	BookId      int    `json:"book_id"`
	LibraryId   int    `json:"library_id"`
	Confirm     int    `json:"confirm"`
}
