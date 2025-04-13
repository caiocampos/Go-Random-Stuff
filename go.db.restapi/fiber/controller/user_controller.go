package controller

import (
	"github.com/gofiber/fiber/v2"

	"go.db.restapi/model"
	serv "go.db.restapi/service"
)

type UserController struct{ serv *serv.UserService }

func (u *UserController) Init(app *fiber.App) {
	if u.serv == nil {
		u.serv = &serv.UserService{}
		app.Get("/user", u.findAll)
		app.Get("/user/id/{id}", u.findByID)
		app.Get("/user/name/{name}", u.findByName)
		app.Post("/user", u.insert)
		app.Delete("/user", u.delete)
		app.Delete("/user/id/{id}", u.deleteByID)
		app.Put("/user", u.update)
	}
}

func (u *UserController) findAll(c *fiber.Ctx) error {
	users, err := u.serv.FindAll(c.Context())
	if err != nil {
		return c.Status(400).SendString(err.Error())
	}
	if users == nil {
		users = []model.User{}
	}
	return c.Status(fiber.StatusOK).JSON(users)
}

func (u *UserController) findByName(c *fiber.Ctx) error {
	user, err := u.serv.FindByName(c.Context(), c.Params("name"))
	if err != nil {
		return c.Status(400).SendString("Name not found")
	}
	return c.Status(fiber.StatusOK).JSON(user)
}

func (u *UserController) findByID(c *fiber.Ctx) error {
	user, err := u.serv.FindByID(c.Context(), c.Params("id"))
	if err != nil {
		return c.Status(400).SendString("ID not found")
	}
	return c.Status(fiber.StatusOK).JSON(user)
}

func (u *UserController) insert(c *fiber.Ctx) error {
	user := new(model.User)
	if err := c.BodyParser(user); err != nil {
		return c.Status(400).SendString("Invalid request")
	}
	userResult, err := u.serv.Insert(c.Context(), *user)
	if err != nil {
		return c.Status(400).SendString(err.Error())
	}
	return c.Status(fiber.StatusCreated).JSON(userResult)
}

func (u *UserController) delete(c *fiber.Ctx) error {
	user := new(model.User)
	if err := c.BodyParser(user); err != nil {
		return c.Status(400).SendString("Invalid request")
	}
	if err := u.serv.Delete(c.Context(), *user); err != nil {
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
	user := new(model.User)
	if err := c.BodyParser(user); err != nil {
		return c.Status(400).SendString("Invalid request")
	}
	if err := u.serv.Update(c.Context(), *user); err != nil {
		return c.Status(400).SendString(err.Error())
	}
	return c.SendStatus(fiber.StatusOK)
}
