package todo

import (
	"todo/internal/database"

	"github.com/gofiber/fiber/v2"
)

func APIV1(c *fiber.Ctx) error {
	return c.JSON(&fiber.Map{
		"loc": "api v1",
	})
}

func APIV2(c *fiber.Ctx) error {
	return c.JSON(&fiber.Map{
		"loc": "api v2",
	})
}

func List(c *fiber.Ctx) error {
	todoSlice, err := database.DBClient.Todo.Query().All(database.DBCtx)
	if err != nil {
		return err
	}

	return c.JSON(todoSlice)
}

func Retrieve(c *fiber.Ctx) error {
	todoId, err := c.ParamsInt("id")
	if err != nil {
		return fiber.NewError(fiber.ErrBadRequest.Code, "path missing")
	}
	todo, err := database.DBClient.Todo.Get(database.DBCtx, todoId)
	if err != nil {
		return err
	}
	return c.JSON(todo)
}

func Create(c *fiber.Ctx) error {
	t := new(Todo)
	if err := c.BodyParser(t); err != nil {
		return err
	}

	ct, err := CreateTodo(*t)
	if err != nil {
		return err
	}
	return c.JSON(ct)
}
