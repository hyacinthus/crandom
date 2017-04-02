package main

import (
	"github.com/jinzhu/configor"
	"github.com/joho/godotenv"
)

var config = struct {
	Debug bool `default:"false"`

	DB struct {
		Host string `default:"mongo"`
		Port string `default:"27017"`
		Name string `default:"cold"`
	}
}{}

func init() {
	godotenv.Load()
	configor.Load(&config)
}
