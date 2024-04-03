package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/template/mustache/v2"
	"todo/internal/config"
	"todo/internal/router"
)

func main() {
	engine := mustache.New("./views", ".mustache")
	engine.Reload(true)
	//engine.Layout("layout/main")

	app := fiber.New(fiber.Config{
		Views:             engine,
		EnablePrintRoutes: true,
		ViewsLayout:       "layout/main",
	})

	app.Static("static", "./static", fiber.Static{
		Browse: true,
	})

	app.Use(logger.New(logger.Config{}))
	app.Use(helmet.New(helmet.Config{}))

	router.SetupRoutes(app)

	host := fmt.Sprintf("%s:%s", config.Settings.App.Host, config.Settings.App.Port)
	if err := app.Listen(host); err != nil {
		panic(err)
	}
}
