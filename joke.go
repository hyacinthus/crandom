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
	oid := bson.ObjectIdHex(c.Param("id"))
	if !oid.Valid() {
		return c.JSON(http.StatusBadRequest, newErrMsg("InvalidID", "invalid request joke id"))
	}
	err := col.FindId(bson.ObjectIdHex(c.Param("id"))).One(&j)
	if err == mgo.ErrNotFound {
		return c.JSON(http.StatusNotFound, newErrMsg("NotFound", err.Error()))
	}
	if err != nil {
		return c.JSON(http.StatusInternalServerError, newErrMsg("ServerError", err.Error()))
	}
	return c.JSON(http.StatusOK, j)
}
