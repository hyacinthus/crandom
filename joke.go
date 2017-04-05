package main

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type joke struct {
	ID      bson.ObjectId `bson:"_id,omitempty" json:"id"`
	Content string        `bson:"content" json:"content"`
	Answer  string        `bson:"answer,omitempty" json:"answer,omitempty"`
	VIA     string        `bson:"via,omitempty" json:"via,omitempty"`
	URL     string        `bson:"url,omitempty" json:"url,omitempty"`
	Created string        `bson:"_created,omitempty" json:"created,omitempty"`
	Updated string        `bson:"_updated,omitempty" json:"updated,omitempty"`
	ETag    string        `bson:"_etag,omitempty" json:"-"`
}

func getJoke(c echo.Context) error {
	var j joke
	col := db.C("joke")
	id := c.Param("id")
	if !bson.IsObjectIdHex(id) {
		return newHTTPError(http.StatusBadRequest, "InvalidID", "invalid request joke id")
	}
	oid := bson.ObjectIdHex(id)
	err := col.FindId(bson.ObjectIdHex(id)).One(&j)
	if err == mgo.ErrNotFound {
		return newHTTPError(http.StatusNotFound, "NotFound", err.Error())
	}
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, j)
}

func getRandomJokes(c echo.Context) error {
	// init size
	var size int
	var err error
	reqSize := c.QueryParam("size")
	if reqSize == "" {
		size = config.PageSize
	} else {
		size, err = strconv.Atoi(reqSize)
		if err != nil {
			return newHTTPError(http.StatusBadRequest, "InvalidParam", "size must be a number")
		}
	}

	// get random jokes
	var jokes = make([]joke, size)
	col := db.C("joke")
	err = col.Pipe([]bson.M{{"$sample": bson.M{"size": size}}}).All(&jokes)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, jokes)
}
