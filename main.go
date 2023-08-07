package main

import (
	"github.com/brianpl990227/hello-fiber/controllers"
	"github.com/brianpl990227/hello-fiber/db"
	"github.com/brianpl990227/hello-fiber/domain/entity"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	
	db.DbConnection()
	db.Context.AutoMigrate(entity.User{},
		entity.Task{})

	app := fiber.New()

	app.Use(logger.New()) //app.User es para el middleware

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	userGroup := app.Group("/users")
	userGroup.Post("", controllers.RegisterUserEndpoint)
	userGroup.Get("", controllers.GetAllUsersEndpoint)
	userGroup.Get("/:id", controllers.GetOneUserEndpoint)
	userGroup.Delete("/:id", controllers.DeleteUserHandler)
	taskGroup := app.Group("/tasks")
	taskGroup.Post("", controllers.AddTaskEndpoint)

	app.Listen(":3000")
}
