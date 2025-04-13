package modules

import (
	"github.com/matzefriedrich/parsley/pkg/types"
	"go.db.restapi/server"
)

func ConfigureServer(registry types.ServiceRegistry) error {
	registry.Register(server.NewFiberServer, types.LifetimeSingleton)

	return nil
}
