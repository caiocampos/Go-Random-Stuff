package modules

import (
	"github.com/matzefriedrich/parsley/pkg/types"
	"go.db.restapi/database"
)

func ConfigureDataBase(registry types.ServiceRegistry) error {
	registry.Register(database.NewMongoDataBase, types.LifetimeTransient)

	return nil
}
