package bootstrap

import (
	"github.com/ahbanavi/go-blog/app/exceptions"
	"github.com/ahbanavi/go-blog/routes"
	"github.com/gofiber/fiber/v2"
)

func NewApp() *fiber.App {

	app := fiber.New(fiber.Config{
		ErrorHandler: exceptions.JSONErrorHandler,
	})
	// app.Use(recover.New())
	routes.Setup(app)

	return app
}
