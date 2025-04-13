package repository

import (
	"context"
	"encoding/json"
	"log"

	"go.db.restapi/database"
	"go.db.restapi/model"

	"github.com/redis/rueidis"
)

// UserRedisRepository defines the repository for the User entity
type UserRedisRepository struct {
	started  bool
	database database.Database[database.RedisInstance]
}

func NewUserRedisRepository(database database.Database[database.RedisInstance]) UserCacheRepository[model.User] {
	return &UserRedisRepository{
		started:  false,
		database: database,
	}
}

func (u *UserRedisRepository) connect() {
	if !u.started {
		if err := u.database.Connect(); err != nil {
			log.Fatal(err)
		}
		u.started = true
	}
}

func (u *UserRedisRepository) disconnect() {
	if u.started {
		if err := u.database.Disconnect(); err != nil {
			log.Fatal(err)
		}
		u.started = false
	}
}

func (u *UserRedisRepository) client() rueidis.Client {
	if !u.started {
		u.connect()
	}
	return *u.database.Get().Client
}

func (u *UserRedisRepository) GetDBType() string {
	return "redis"
}

// FindByID method returns a user in database
func (u *UserRedisRepository) FindByID(ctx context.Context, id string) (model.User, error) {
	user := model.User{}
	client := u.client()
	err := client.Do(ctx, client.B().Get().Key(getUserKey(id)).Build()).DecodeJSON(&user)
	defer u.disconnect()
	return user, err
}

// Insert method inserts a user in database
func (u *UserRedisRepository) Insert(ctx context.Context, user model.User) (model.User, error) {
	client := u.client()
	jsonValue, err := json.Marshal(user)
	if err != nil {
		return user, err
	}
	cmd := client.B().Set().Key(getUserOIDKey(user.ID)).Value(rueidis.BinaryString(jsonValue)).ExSeconds(cacheDuration).Build()
	err = client.Do(ctx, cmd).Error()
	defer u.disconnect()
	return user, err
}

// DeleteByID method deletes a user in database
func (u *UserRedisRepository) DeleteByID(ctx context.Context, id string) error {
	client := u.client()
	err := client.Do(ctx, client.B().Del().Key(getUserKey(id)).Build()).Error()
	defer u.disconnect()
	return err
}

// Delete method deletes a user from database
func (u *UserRedisRepository) Delete(ctx context.Context, user model.User) error {
	return u.DeleteByID(ctx, getUserOIDKey(user.ID))
}

// Update method updates a user in database
func (u *UserRedisRepository) Update(ctx context.Context, user model.User) error {
	_, err := u.Insert(ctx, user)
	return err
}
