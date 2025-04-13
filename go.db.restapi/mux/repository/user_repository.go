package repository

import (
	"gopkg.in/mgo.v2/bson"

	"go.db.restapi/model"
)

// UserRepository defines the repository for the User entity
type UserRepository struct {
	collection string
}

func (u *UserRepository) connect() {
	u.collection = "user"
	connect()
}

// FindAll method returns all users in database
func (u *UserRepository) FindAll() ([]model.User, error) {
	u.connect()
	var users []model.User
	err := db.C(u.collection).Find(bson.M{}).All(&users)
	return users, err
}

// FindByName method returns a user in database
func (u *UserRepository) FindByName(name string) (model.User, error) {
	u.connect()
	var user model.User
	err := db.C(u.collection).Find(bson.M{"name": name}).One(&user)
	return user, err
}

// FindByID method returns a user in database
func (u *UserRepository) FindByID(id string) (model.User, error) {
	u.connect()
	var user model.User
	err := db.C(u.collection).FindId(bson.ObjectIdHex(id)).One(&user)
	return user, err
}

// Insert method inserts a user in database
func (u *UserRepository) Insert(user model.User) error {
	u.connect()
	return db.C(u.collection).Insert(&user)
}

// Delete method deletes a user in database
func (u *UserRepository) Delete(user model.User) error {
	u.connect()
	return db.C(u.collection).Remove(&user)
}

// Update method updates a user in database
func (u *UserRepository) Update(user model.User) error {
	u.connect()
	return db.C(u.collection).UpdateId(user.ID, &user)
}
