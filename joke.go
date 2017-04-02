package main

import (
	"net/http"

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
	n, _ := col.Count()
	c.Logger().Debugf("id:%s, count:%d", c.Param("id"), n)
	err := col.FindId(bson.ObjectIdHex(c.Param("id"))).One(&j)
	if err == mgo.ErrNotFound {
		return echo.ErrNotFound
	}
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, j)
}
