package todo

import "github.com/gofiber/fiber/v2"

func IndexPage(c *fiber.Ctx) error {
	return c.Render("todo/index", &fiber.Map{})
}
