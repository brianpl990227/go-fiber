package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/brianpl990227/hello-fiber/controllers"
)


func main() {
	app := fiber.New()

	app.Use(logger.New()) //app.User es para el middleware

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})


	userGroup := app.Group("/users")
	userGroup.Post("", controllers.RegisterUserEndpoint)
	userGroup.Get("", controllers.GetAllUsersEndpoint)

	app.Listen(":3000")
}
