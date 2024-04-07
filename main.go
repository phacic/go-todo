package main

import (
	"errors"
	"fmt"
	"todo/ent"
	"todo/internal/config"
	"todo/internal/database"
	"todo/internal/router"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/template/mustache/v2"
	_ "github.com/lib/pq"
)

func errHandler(ctx *fiber.Ctx, err error) error {
	// default code
	code := fiber.StatusInternalServerError

	// fiber err
	var e *fiber.Error
	if errors.As(err, &e) {
		code = e.Code
	}

	var entErr *ent.NotFoundError
	if errors.As(err, &entErr) {
		code = fiber.StatusNotFound
	}

	// Set Content-Type: text/plain; charset=utf-8
	ctx.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSONCharsetUTF8)

	// Return status code with error message
	return ctx.Status(code).JSON(&fiber.Map{"message": err.Error()})
}

func main() {
	database.Connect(true)
	defer database.DBClient.Close()

	engine := mustache.New("./views", ".mustache")
	engine.Reload(true)
	//engine.Layout("layout/main")

	app := fiber.New(fiber.Config{
		ErrorHandler:      errHandler,
		Views:             engine,
		EnablePrintRoutes: false,
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
