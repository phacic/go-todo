package todoNot

import "github.com/gofiber/fiber/v2"

func IndexPage(c *fiber.Ctx) error {
	return c.Render("todo-not/index", &fiber.Map{})
}
