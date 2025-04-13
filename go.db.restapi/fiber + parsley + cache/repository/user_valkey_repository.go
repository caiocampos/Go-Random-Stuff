package repository

import (
	"context"
	"encoding/json"
	"log"

	"github.com/valkey-io/valkey-go"

	"go.db.restapi/database"
	"go.db.restapi/model"
)

// UserValkeyRepository defines the repository for the User entity
type UserValkeyRepository struct {
	started  bool
	database database.Database[database.ValkeyInstance]
}

func NewUserValkeyRepository(database database.Database[database.ValkeyInstance]) UserCacheRepository[model.User] {
	return &UserValkeyRepository{
		started:  false,
		database: database,
	}
}

func (u *UserValkeyRepository) connect() {
	if !u.started {
		if err := u.database.Connect(); err != nil {
			log.Fatal(err)
		}
		u.started = true
	}
}

func (u *UserValkeyRepository) disconnect() {
	if u.started {
		if err := u.database.Disconnect(); err != nil {
			log.Fatal(err)
		}
		u.started = false
	}
}

func (u *UserValkeyRepository) client() valkey.Client {
	if !u.started {
		u.connect()
	}
	return *u.database.Get().Client
}

func (u *UserValkeyRepository) GetDBType() string {
	return "valkey"
}

// FindByID method returns a user in database
func (u *UserValkeyRepository) FindByID(ctx context.Context, id string) (model.User, error) {
	user := model.User{}
	client := u.client()
	err := client.Do(ctx, client.B().Get().Key(getUserKey(id)).Build()).DecodeJSON(&user)
	defer u.disconnect()
	return user, err
}

// Insert method inserts a user in database
func (u *UserValkeyRepository) Insert(ctx context.Context, user model.User) (model.User, error) {
	client := u.client()
	jsonValue, err := json.Marshal(user)
	if err != nil {
		return user, err
	}
	cmd := client.B().Set().Key(getUserOIDKey(user.ID)).Value(valkey.BinaryString(jsonValue)).ExSeconds(cacheDuration).Build()
	err = client.Do(ctx, cmd).Error()
	defer u.disconnect()
	return user, err
}

// DeleteByID method deletes a user in database
func (u *UserValkeyRepository) DeleteByID(ctx context.Context, id string) error {
	client := u.client()
	err := client.Do(ctx, client.B().Del().Key(getUserKey(id)).Build()).Error()
	defer u.disconnect()
	return err
}

// Delete method deletes a user from database
func (u *UserValkeyRepository) Delete(ctx context.Context, user model.User) error {
	return u.DeleteByID(ctx, getUserOIDKey(user.ID))
}

// Update method updates a user in database
func (u *UserValkeyRepository) Update(ctx context.Context, user model.User) error {
	_, err := u.Insert(ctx, user)
	return err
}
