package service

import (
	"context"
	"log"

	"go.db.restapi/config"
	"go.db.restapi/model"
	repo "go.db.restapi/repository"
)

// UserServiceImpl defines the services for the User entity
type UserServiceImpl struct {
	repo  repo.UserRepository[model.User]
	cache repo.UserCacheRepository[model.User]
}

func NewUserServiceImpl(config config.ConfigLoader, repo repo.UserRepository[model.User], cacheRepos []repo.UserCacheRepository[model.User]) UserService[model.User] {
	cacheType := config.Get().Cache.Type
	for _, r := range cacheRepos {
		if r.GetDBType() == cacheType {
			return &UserServiceImpl{
				repo:  repo,
				cache: r,
			}
		}
	}
	log.Fatal("cannot find " + cacheType + " repository for user model")
	return nil
}

// FindAll method returns all users in database
func (u *UserServiceImpl) FindAll(ctx context.Context) ([]model.User, error) {
	return u.repo.FindAll(ctx)
}

// FindByName method returns a user in database
func (u *UserServiceImpl) FindByName(ctx context.Context, name string) (model.User, error) {
	result, err := u.repo.FindByName(ctx, name)
	if err != nil {
		return model.User{}, err
	}
	u.cache.Insert(ctx, result)
	return result, err
}

// FindByID method returns a user in database
func (u *UserServiceImpl) FindByID(ctx context.Context, id string) (model.User, error) {
	result, err := u.cache.FindByID(ctx, id)
	if err != nil {
		result, err = u.repo.FindByID(ctx, id)
		if err != nil {
			return model.User{}, err
		}
		u.cache.Insert(ctx, result)
	}
	return result, err
}

// Insert method inserts a user in database
func (u *UserServiceImpl) Insert(ctx context.Context, user model.User) (model.User, error) {
	result, err := u.repo.Insert(ctx, user)
	if err != nil {
		return model.User{}, err
	}
	u.cache.Insert(ctx, result)
	return result, err
}

// Delete method deletes a user in database
func (u *UserServiceImpl) Delete(ctx context.Context, user model.User) error {
	u.cache.Delete(ctx, user)
	err := u.repo.Delete(ctx, user)
	return err
}

// Delete method deletes a user in database
func (u *UserServiceImpl) DeleteByID(ctx context.Context, id string) error {
	u.cache.DeleteByID(ctx, id)
	err := u.repo.DeleteByID(ctx, id)
	return err
}

// Update method updates a user in database
func (u *UserServiceImpl) Update(ctx context.Context, user model.User) error {
	err := u.repo.Update(ctx, user)
	if err != nil {
		return err
	}
	u.cache.Update(ctx, user)
	return err
}
