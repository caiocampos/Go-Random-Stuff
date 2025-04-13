package model

import "go.mongodb.org/mongo-driver/bson/primitive"

// User struct represents a user with name, password and profession
type User struct {
	ID         primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Name       string             `bson:"name" json:"name"`
	Password   string             `bson:"password" json:"password"`
	Profession string             `bson:"profession" json:"profession"`
}
