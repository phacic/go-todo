package todo

import "github.com/gofiber/fiber/v2"

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
