package service

import (
	"context"

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

func (u *UserService) repository() *repo.UserRepository {
	if u.repo == nil {
		u.init()
	}
	return u.repo
}

// FindAll method returns all users in database
func (u *UserService) FindAll(ctx context.Context) ([]model.User, error) {
	return u.repository().FindAll(ctx)
}

// FindByName method returns a user in database
func (u *UserService) FindByName(ctx context.Context, name string) (model.User, error) {
	return u.repository().FindByName(ctx, name)
}

// FindByID method returns a user in database
func (u *UserService) FindByID(ctx context.Context, id string) (model.User, error) {
	return u.repository().FindByID(ctx, id)
}

// Insert method inserts a user in database
func (u *UserService) Insert(ctx context.Context, user model.User) (model.User, error) {
	return u.repository().Insert(ctx, user)
}

// Delete method deletes a user in database
func (u *UserService) Delete(ctx context.Context, user model.User) error {
	return u.repository().Delete(ctx, user)
}

// Delete method deletes a user in database
func (u *UserService) DeleteByID(ctx context.Context, id string) error {
	return u.repository().DeleteByID(ctx, id)
}

// Update method updates a user in database
func (u *UserService) Update(ctx context.Context, user model.User) error {
	return u.repository().Update(ctx, user)
}
