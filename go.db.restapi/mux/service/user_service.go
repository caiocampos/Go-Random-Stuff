package service

import (
	"go.db.restapi/model"
	repo "go.db.restapi/repository"
)

// UserService defines the services for the User entity
type UserService struct{ repo *repo.UserRepository }

func (u *UserService) init() {
	if u.repo == nil {
		u.repo = &repo.UserRepository{}
	}
}

// FindAll method returns all users in database
func (u *UserService) FindAll() ([]model.User, error) {
	u.init()
	return u.repo.FindAll()
}

// FindByName method returns a user in database
func (u *UserService) FindByName(name string) (model.User, error) {
	u.init()
	return u.repo.FindByName(name)
}

// FindByID method returns a user in database
func (u *UserService) FindByID(id string) (model.User, error) {
	u.init()
	return u.repo.FindByID(id)
}

// Insert method inserts a user in database
func (u *UserService) Insert(user model.User) error {
	u.init()
	return u.repo.Insert(user)
}

// Delete method deletes a user in database
func (u *UserService) Delete(user model.User) error {
	u.init()
	return u.repo.Delete(user)
}

// Update method updates a user in database
func (u *UserService) Update(user model.User) error {
	u.init()
	return u.repo.Update(user)
}
