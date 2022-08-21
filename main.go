package main

import (
	"github.com/ahbanavi/go-blog/app/controllers"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "Hello World!",
		})
	})

	// Register routes here, TODO: move routes to separate package
	app.Get("/user", controllers.GetAllUsers)

	app.Listen(":3000")
}
