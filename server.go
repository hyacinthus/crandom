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

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Logger.SetLevel(log.DEBUG)

	// Routes
	e.GET("/jokes/:id", getJoke)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}
