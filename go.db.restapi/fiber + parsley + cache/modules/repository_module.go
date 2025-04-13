package modules

import (
	"github.com/matzefriedrich/parsley/pkg/features"
	"github.com/matzefriedrich/parsley/pkg/types"
	"go.db.restapi/model"
	"go.db.restapi/repository"
)

func ConfigureRepositories(registry types.ServiceRegistry) error {
	registry.Register(repository.NewUserMongoRepository, types.LifetimeTransient)

	registry.Register(repository.NewUserRedisRepository, types.LifetimeTransient)
	registry.Register(repository.NewUserValkeyRepository, types.LifetimeTransient)
	features.RegisterList[repository.UserCacheRepository[model.User]](registry)

	return nil
}
