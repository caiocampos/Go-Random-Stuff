package modules

import (
	"github.com/matzefriedrich/parsley/pkg/features"
	"github.com/matzefriedrich/parsley/pkg/types"
	"go.db.restapi/controller"
)

func ConfigureControllers(registry types.ServiceRegistry) error {
	registry.Register(controller.NewUserController, types.LifetimeTransient)
	features.RegisterList[controller.Controller](registry)

	return nil
}
