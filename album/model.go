package album

import "gopkg.in/mgo.v2/bson"

//Album represents a music album
type Album struct {
	ID     bson.ObjectId `bson:"_id"`
	Title  string        `json:"title"`
	Artist string        `json:"artist"`
	Year   int32         `json:"year"`
}

//Albums is an array of Album
type Albums []Album
