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

func GetOneUserEndpoint(c *fiber.Ctx) error {
	
	id := c.Params("id")
	  /*
	 	Si fuera un QueryString 
	  	name := c.Query("name")
	  	age := c.Query("age")
	  */
	  users, errCase := userCases.GetOneUserHandler(id)

	  if errCase != nil {
		  return c.Status(fiber.StatusInternalServerError).JSON(exception.ErrorResponse{
			  Status: fiber.StatusInternalServerError,
			  Message: errCase.Error(),
		  })
	  }
  
	  return c.Status(fiber.StatusOK).JSON(users)

}

func DeleteUserHandler(c *fiber.Ctx) error {

	id := c.Params("id")

	user, errCase := userCases.DeleteUserHandler(id)

	if errCase != nil {
		return c.Status(errCase.Status).JSON(errCase)
	}

	return c.Status(fiber.StatusOK).JSON(user)

}
