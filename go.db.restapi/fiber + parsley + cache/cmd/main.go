package main

import (
	"context"

	"github.com/matzefriedrich/parsley/pkg/bootstrap"

	"go.db.restapi/app"
	"go.db.restapi/modules"
)

func main() {
	context := context.Background()

	// Runs a Parsley-enabled app
	bootstrap.RunParsleyApplication(context, app.NewApp,
		modules.ConfigureServer,
		modules.ConfigureControllers,
		modules.ConfigureConfigloader,
		modules.ConfigureServices,
		modules.ConfigureRepositories,
		modules.ConfigureDataBase)
}
