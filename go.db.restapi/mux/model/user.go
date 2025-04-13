package model

import "gopkg.in/mgo.v2/bson"

// User struct represents a user with name, password and profession
type User struct {
	ID         bson.ObjectId `bson:"_id" json:"id"`
	Name       string        `bson:"name" json:"name"`
	Password   string        `bson:"password" json:"password"`
	Profession string        `bson:"profession" json:"profession"`
}
