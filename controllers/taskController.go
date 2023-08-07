package controllers

import (
	"github.com/brianpl990227/hello-fiber/domain/dto"
	"github.com/brianpl990227/hello-fiber/domain/exception"
	"github.com/brianpl990227/hello-fiber/useCases/taskCases"
	"github.com/gofiber/fiber/v2"
)

func AddTaskEndpoint(c *fiber.Ctx)error {
	
	dto := &dto.CreateTaskDTO{}

	err := c.BodyParser(dto)

	if err != nil {
		return err
	}

	result, err := taskCases.CreateTaskHandler(dto)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(exception.ErrorResponse{
			Status: fiber.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(result)
	
}