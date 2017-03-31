package main

import (
	"fmt"

	"github.com/jinzhu/configor"
	"github.com/joho/godotenv"
)

var config = struct {
	Debug bool `default:"false"`

	DB struct {
		Name     string
		User     string `default:"root"`
		Password string `required:"true" env:"DBPassword"`
		Port     uint   `default:"3306"`
	}
}{}

func init() {
	godotenv.Load()
	configor.Load(&config)
	fmt.Printf("config: %#v", config)
}
