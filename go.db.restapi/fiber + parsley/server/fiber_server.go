package server

import (
	"log"
	"strconv"
	"strings"

	"github.com/bytedance/sonic"
	"github.com/gofiber/fiber/v2"

	"go.db.restapi/config"
	ctrl "go.db.restapi/controller"
)

type FiberServer struct {
	app         *fiber.App
	started     bool
	config      config.ConfigLoader
	controllers []ctrl.Controller
}

func NewFiberServer(config config.ConfigLoader, controllers []ctrl.Controller) Server {
	return &FiberServer{config: config, controllers: controllers, started: false}
}

// Init method boots the end-points for the server
func (fs *FiberServer) Init() error {
	if !fs.started {
		err := fs.config.Load()
		if err != nil {
			log.Fatal(err)
			return err
		}
		if strings.ToLower(fs.config.Get().App.JsonProcessor) == "sonic" {
			fs.app = fiber.New(fiber.Config{
				JSONEncoder: sonic.Marshal,
				JSONDecoder: sonic.Unmarshal,
			})
		} else {
			fs.app = fiber.New()
		}
		for _, controller := range fs.controllers {
			controller.Init(fs.app)
		}
		port := strconv.Itoa(fs.config.Get().App.Port)
		err = fs.app.Listen(":" + port)
		if err != nil {
			log.Fatal(err)
			return err
		}
		fs.started = true
	}
	return nil
}
