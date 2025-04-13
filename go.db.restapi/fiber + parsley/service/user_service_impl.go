package service

import (
	"context"

	"go.db.restapi/model"
	repo "go.db.restapi/repository"
)

// UserServiceImpl defines the services for the User entity
type UserServiceImpl struct {
	repo repo.UserRepository[model.User]
}

func NewUserServiceImpl(repo repo.UserRepository[model.User]) UserService[model.User] {
	return &UserServiceImpl{
		repo: repo,
	}
}

// FindAll method returns all users in database
func (u *UserServiceImpl) FindAll(ctx context.Context) ([]model.User, error) {
	return u.repo.FindAll(ctx)
}

// FindByName method returns a user in database
func (u *UserServiceImpl) FindByName(ctx context.Context, name string) (model.User, error) {
	return u.repo.FindByName(ctx, name)
}

// FindByID method returns a user in database
func (u *UserServiceImpl) FindByID(ctx context.Context, id string) (model.User, error) {
	return u.repo.FindByID(ctx, id)
}

// Insert method inserts a user in database
func (u *UserServiceImpl) Insert(ctx context.Context, user model.User) (model.User, error) {
	return u.repo.Insert(ctx, user)
}

// Delete method deletes a user in database
func (u *UserServiceImpl) Delete(ctx context.Context, user model.User) error {
	return u.repo.Delete(ctx, user)
}

// Delete method deletes a user in database
func (u *UserServiceImpl) DeleteByID(ctx context.Context, id string) error {
	return u.repo.DeleteByID(ctx, id)
}

// Update method updates a user in database
func (u *UserServiceImpl) Update(ctx context.Context, user model.User) error {
	return u.repo.Update(ctx, user)
}
