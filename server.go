package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/labstack/gommon/log"
	"gopkg.in/mgo.v2"
)

var (
	db *mgo.Database
)

func main() {
	// mongo conn
	session, err := mgo.Dial(config.DB.Host + ":" + config.DB.Port)
	if err != nil {
		panic(err)
	}
	defer session.Close()
	db = session.DB(config.DB.Name)

	// echo
	e := echo.New()
	e.HTTPErrorHandler = httpErrorHandler

	// debug
	if config.Debug {
		e.Debug = true
		e.Logger.SetLevel(log.DEBUG)
		e.Logger.Debug("start debug mode")
	}

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	// Routes
	e.GET("/joke/random", getRandomJokes)
	e.GET("/jokes/:id", getJoke)
	e.PUT("/jokes/:id", updateJoke)
	e.DELETE("/jokes/:id", deleteJoke)
	e.POST("/jokes", createJoke)

	// Start server
	e.Logger.Fatal(e.Start(":" + config.Port))
}
