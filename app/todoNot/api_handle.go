package todoNot

import "github.com/gofiber/fiber/v2"

func TodoNot(c *fiber.Ctx) error {
	return c.JSON(&fiber.Map{
		"loc": "todo-not api",
	})
}
