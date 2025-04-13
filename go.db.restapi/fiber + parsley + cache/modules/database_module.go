package modules

import (
	"github.com/matzefriedrich/parsley/pkg/types"
	"go.db.restapi/database"
)

func ConfigureDataBase(registry types.ServiceRegistry) error {
	registry.Register(database.NewMongoDataBase, types.LifetimeTransient)
	registry.Register(database.NewRedisDataBase, types.LifetimeTransient)
	registry.Register(database.NewValkeyDataBase, types.LifetimeTransient)

	return nil
}
