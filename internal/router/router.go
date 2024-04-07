package router

import (
	"todo/app/todo"
	"todo/app/todoNot"

	"github.com/gofiber/fiber/v2"
)

func setupPagesRoutes(app *fiber.App) {
	pages := app.Group("")

	todoPages := pages.Group("todo")
	todoPages.Get("", todo.IndexPage)

	todoNotPages := pages.Group("todo-not")
	todoNotPages.Get("", todoNot.IndexPage)
}

func setupAPIRoutes(app *fiber.App) {
	api := app.Group("api")

	v1 := api.Group("v1")
	todoV1 := v1.Group("todo")
	todoV1.Get("", todo.List)
	todoV1.Post("", todo.Create)
	todoV1.Get("/:id", todo.Retrieve)

	toNoDoV1 := v1.Group("todo-not")
	toNoDoV1.Get("", todoNot.TodoNot)

	v2 := api.Group("v2")
	v2.Get("", todo.APIV2)
}

func SetupRoutes(app *fiber.App) {
	setupPagesRoutes(app)
	setupAPIRoutes(app)
}
