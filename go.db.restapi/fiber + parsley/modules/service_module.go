package modules

import (
	"github.com/matzefriedrich/parsley/pkg/types"
	"go.db.restapi/service"
)

func ConfigureServices(registry types.ServiceRegistry) error {
	registry.Register(service.NewUserServiceImpl, types.LifetimeTransient)

	return nil
}
