package modules

import (
	"github.com/matzefriedrich/parsley/pkg/types"
	"go.db.restapi/repository"
)

func ConfigureRepositories(registry types.ServiceRegistry) error {
	registry.Register(repository.NewUserMongoRepository, types.LifetimeTransient)

	return nil
}
