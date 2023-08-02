package controllers

import (
	"github.com/brianpl990227/hello-fiber/domain/dto"
	"github.com/brianpl990227/hello-fiber/domain/exception"
	"github.com/brianpl990227/hello-fiber/useCases/userCases"
	"github.com/gofiber/fiber/v2"
)

func RegisterUserEndpoint(c *fiber.Ctx) error {
	dto := &dto.CreateUserDTO{}

	err := c.BodyParser(dto)

	if err != nil {
		return err
	}

	user, errCase := userCases.CreateUserHandler(dto)

	if errCase != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(exception.ErrorResponse{
			Status: fiber.StatusInternalServerError,
			Message: errCase.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(user)

}

func GetAllUsersEndpoint(c *fiber.Ctx) error {
	

	users, errCase := userCases.GetAllUsersHandler()

	if errCase != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(exception.ErrorResponse{
			Status: fiber.StatusInternalServerError,
			Message: errCase.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(users)

}

// func handleCreateUser(c *fiber.Ctx) error {
// 	user := domain{}

// 	err := c.BodyParser(&user)

// 	if err != nil {
// 		return err
// 	}
// 	user.Id = uuid.NewString()
// 	users = append(users, user)
// 	return c.Status(fiber.StatusOK).JSON(users)
// }
