package main

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	Ipusnas struct {
		Url string `mapstructure:"URL"`
		Api struct {
			Login        string `mapstructure:"LOGIN"`
			RefreshToken string `mapstructure:"REFRESH_TOKEN"`
			BorrowBook   string `mapstructure:"BORROW_BOOK"`
			BookDetail   string `mapstructure:"BOOK_DETAIL"`
		} `mapstructure:"API"`
		LibraryId int `mapstructure:"LIBRARY_ID"`
	} `mapstructure:"IPUSNAS"`
	Account struct {
		Email    string `mapstructure:"EMAIL"`
		Password string `mapstructure:"PASSWORD"`
	} `mapstructure:"ACCOUNT"`
	Client struct {
		Id       string `mapstructure:"ID"`
		Secret   string `mapstructure:"SECRET"`
		DeviceId string `mapstructure:"DEVICE_ID"`
	} `mapstructure:"CLIENT"`
	Targets []int `mapstructure:"TARGETS"`
}

func LoadEnv() *Config {
	var conf Config
	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()

	if err != nil {
		log.Fatal(err.Error())
	}

	err = viper.Unmarshal(&conf)
	if err != nil {
		log.Fatal(err.Error())
	}

	return &conf
}
