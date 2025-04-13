package app

import (
	"context"

	"github.com/matzefriedrich/parsley/pkg/bootstrap"
	"go.db.restapi/server"
)

type parsleyApplication struct {
	server server.Server
}

func NewApp(server server.Server) bootstrap.Application {
	return &parsleyApplication{
		server: server,
	}
}

// Run The entrypoint for the Parsley application.
func (a *parsleyApplication) Run(_ context.Context) error {
	return a.server.Init()
}
