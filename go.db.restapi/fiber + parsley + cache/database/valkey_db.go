package database

import (
	"errors"

	"github.com/valkey-io/valkey-go"
	"go.db.restapi/config"
)

type ValkeyDataBase struct {
	valkey *ValkeyInstance
	config config.ConfigLoader
}

// ValkeyInstance contains the Valkey client object
type ValkeyInstance struct {
	Client *valkey.Client
}

func NewValkeyDataBase(config config.ConfigLoader) Database[ValkeyInstance] {
	return &ValkeyDataBase{
		config: config,
	}
}

// Connect function establish a connection to database
func (m *ValkeyDataBase) Connect() error {
	if m.valkey == nil {
		config := m.config.Get()
		if config.Cache.Type != "valkey" {
			return errors.New("invalid connection type")
		}
		options := valkey.ClientOption{
			InitAddress: []string{config.Cache.Server},
		}
		if config.Cache.Username != nil && config.Cache.Password != nil {
			options.Username = *config.Cache.Username
			options.Password = *config.Cache.Password
		}
		client, err := valkey.NewClient(options)
		if err != nil {
			return err
		}
		m.valkey = &ValkeyInstance{
			Client: &client,
		}
	}
	return nil
}

// Disconnect function closes the connection with database
func (m *ValkeyDataBase) Disconnect() error {
	if client := m.valkey.Client; client != nil {
		(*client).Close()
	}
	m.valkey = nil
	return nil
}

func (m *ValkeyDataBase) Get() *ValkeyInstance {
	m.Connect()
	return m.valkey
}
