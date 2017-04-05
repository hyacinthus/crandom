package main

import (
	"net/http"
	"strconv"
	"time"

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
	Created time.Time     `bson:"_created,omitempty" json:"created,omitempty"`
	Updated time.Time     `bson:"_updated,omitempty" json:"updated,omitempty"`
	ETag    string        `bson:"_etag,omitempty" json:"-"`
}

func getJoke(c echo.Context) error {
	var j joke
	col := db.C("joke")
	id := c.Param("id")
	if !bson.IsObjectIdHex(id) {
		return newHTTPError(http.StatusBadRequest, "InvalidID", "invalid request joke id")
	}
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
	// init page size
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

func createJoke(c echo.Context) error {
	j := &joke{
		ID:      bson.NewObjectId(),
		Created: bson.Now(),
	}
	if err := c.Bind(j); err != nil {
		return newHTTPError(http.StatusBadRequest, "InvalidRequestData", err.Error())
	}
	// save to db
	col := db.C("joke")
	err := col.Insert(j)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusCreated, j)
}

func updateJoke(c echo.Context) error {
	// get id
	id := c.Param("id")
	if !bson.IsObjectIdHex(id) {
		return newHTTPError(http.StatusBadRequest, "InvalidID", "invalid request joke id")
	}
	// get data
	var j = new(joke)
	if err := c.Bind(j); err != nil {
		return newHTTPError(http.StatusBadRequest, "InvalidRequestData", err.Error())
	}
	var data = make(bson.M)
	data["_updated"] = bson.Now()
	if j.Content != "" {
		data["content"] = j.Content
	}
	if j.Answer != "" {
		data["answer"] = j.Answer
	}
	if j.VIA != "" {
		data["via"] = j.VIA
	}
	if j.URL != "" {
		data["url"] = j.URL
	}
	// update db
	col := db.C("joke")
	err := col.UpdateId(id, data)
	if err == mgo.ErrNotFound {
		return newHTTPError(http.StatusBadRequest, "InvalidID", err.Error())
	}
	if err != nil {
		return err
	}
	// return new joke
	var newj = new(joke)
	err = col.FindId(bson.ObjectIdHex(id)).One(&newj)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, newj)
}
