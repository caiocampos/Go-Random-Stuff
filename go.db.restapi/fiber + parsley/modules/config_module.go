package modules

import (
	"github.com/matzefriedrich/parsley/pkg/types"
	"go.db.restapi/config"
)

func ConfigureConfigloader(registry types.ServiceRegistry) error {
	registry.Register(config.NewTOMLConfigloader, types.LifetimeSingleton)

	return nil
}
