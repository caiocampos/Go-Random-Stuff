package controller

import (
	"github.com/gofiber/fiber/v2"

	"go.db.restapi/model"
	serv "go.db.restapi/service"
)

type UserController struct {
	serv    serv.UserService[model.User]
	started bool
}

func NewUserController(serv serv.UserService[model.User]) Controller {
	return &UserController{serv, false}
}

func (u *UserController) Init(app *fiber.App) {
	if !u.started {
		users := app.Group("/users")
		users.Get("/", u.findAll)
		users.Get("/id/:id", u.findByID)
		users.Get("/name/:name", u.findByName)
		users.Post("/", u.insert)
		users.Delete("/", u.delete)
		users.Delete("/id/:id", u.deleteByID)
		users.Put("/", u.update)
		u.started = true
	}
}

func (u *UserController) findAll(c *fiber.Ctx) error {
	users, err := u.serv.FindAll(c.Context())
	if err != nil {
		return c.Status(400).SendString(err.Error())
	}
	if users == nil {
		users = []model.User{}
	} else {
		parsedUsers := []model.User{}
		for _, user := range users {
			parsedUsers = append(parsedUsers, user.GetWithoutPass())
		}
		users = parsedUsers
	}
	return c.Status(fiber.StatusOK).JSON(users)
}

func (u *UserController) findByName(c *fiber.Ctx) error {
	user, err := u.serv.FindByName(c.Context(), c.Params("name"))
	if err != nil {
		return c.Status(400).SendString("Name not found")
	}
	return c.Status(fiber.StatusOK).JSON(user.GetWithoutPass())
}

func (u *UserController) findByID(c *fiber.Ctx) error {
	user, err := u.serv.FindByID(c.Context(), c.Params("id"))
	if err != nil {
		return c.Status(400).SendString("ID not found")
	}
	return c.Status(fiber.StatusOK).JSON(user.GetWithoutPass())
}

func (u *UserController) insert(c *fiber.Ctx) error {
	user := model.User{}
	if err := c.BodyParser(user); err != nil {
		return c.Status(400).SendString("Invalid request")
	}
	userResult, err := u.serv.Insert(c.Context(), user)
	if err != nil {
		return c.Status(400).SendString(err.Error())
	}
	return c.Status(fiber.StatusCreated).JSON(userResult.GetWithoutPass())
}

func (u *UserController) delete(c *fiber.Ctx) error {
	user := model.User{}
	if err := c.BodyParser(user); err != nil {
		return c.Status(400).SendString("Invalid request")
	}
	if err := u.serv.Delete(c.Context(), user); err != nil {
		return c.Status(400).SendString(err.Error())
	}
	return c.SendStatus(fiber.StatusNoContent)
}

func (u *UserController) deleteByID(c *fiber.Ctx) error {
	err := u.serv.DeleteByID(c.Context(), c.Params("id"))
	if err != nil {
		return c.Status(400).SendString("ID not found")
	}
	return c.SendStatus(fiber.StatusNoContent)
}

func (u *UserController) update(c *fiber.Ctx) error {
	user := model.User{}
	if err := c.BodyParser(user); err != nil {
		return c.Status(400).SendString("Invalid request")
	}
	if err := u.serv.Update(c.Context(), user); err != nil {
		return c.Status(400).SendString(err.Error())
	}
	return c.SendStatus(fiber.StatusOK)
}
