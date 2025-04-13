package database

import (
	"errors"

	"github.com/redis/rueidis"

	"go.db.restapi/config"
)

type RedisDataBase struct {
	redis  *RedisInstance
	config config.ConfigLoader
}

// RedisInstance contains the Redis client object
type RedisInstance struct {
	Client *rueidis.Client
}

func NewRedisDataBase(config config.ConfigLoader) Database[RedisInstance] {
	return &RedisDataBase{
		config: config,
	}
}

// Connect function establish a connection to database
func (m *RedisDataBase) Connect() error {
	if m.redis == nil {
		config := m.config.Get()
		if config.Cache.Type != "redis" {
			return errors.New("invalid connection type")
		}
		options := rueidis.ClientOption{
			InitAddress: []string{config.Cache.Server},
		}
		if config.Cache.Username != nil && config.Cache.Password != nil {
			options.Username = *config.Cache.Username
			options.Password = *config.Cache.Password
		}
		client, err := rueidis.NewClient(options)
		if err != nil {
			return err
		}
		m.redis = &RedisInstance{
			Client: &client,
		}
	}
	return nil
}

// Disconnect function closes the connection with database
func (m *RedisDataBase) Disconnect() error {
	if client := m.redis.Client; client != nil {
		(*client).Close()
	}
	m.redis = nil
	return nil
}

func (m *RedisDataBase) Get() *RedisInstance {
	m.Connect()
	return m.redis
}
