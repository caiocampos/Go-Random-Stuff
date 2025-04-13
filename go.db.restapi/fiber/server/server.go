package server

import (
	"log"
	"strconv"

	"github.com/gofiber/fiber/v2"

	"go.db.restapi/config"
	ctrl "go.db.restapi/controller"
)

type server struct {
	app    *fiber.App
	user   ctrl.UserController
}

var serv *server

// Init method boots the end-points for the server
func Init() {
	if serv == nil {
		serv := new(server)
		serv.app = fiber.New()
		serv.user = ctrl.UserController{}
		serv.user.Init(serv.app)

		config.ReadTOML()
		port := strconv.Itoa(config.TOMLConfig.App.Port)
		if err := serv.app.Listen(":"+port); err != nil {
			log.Fatal(err)
		}
	}
}
