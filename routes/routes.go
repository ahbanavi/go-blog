package routes

import (
	c "github.com/ahbanavi/go-blog/app/controllers"
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "Hello World!",
		})
	})

	api := app.Group("/api")

	users := api.Group("/users")
	users.Get("/", c.GetAllUsers)
	users.Get("/:id", c.GetUser)
	users.Post("/", c.CreateUser)
	users.Put("/:id", c.UpdateUser)
	users.Delete("/:id", c.DeleteUser)
}
