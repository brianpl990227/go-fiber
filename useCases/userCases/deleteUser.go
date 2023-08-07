package userCases

import (
	"github.com/brianpl990227/hello-fiber/domain/entity"
	"github.com/brianpl990227/hello-fiber/domain/exception"
	"github.com/brianpl990227/hello-fiber/repositories/userRepo"
	"github.com/gofiber/fiber/v2"
)


func DeleteUserHandler(id string) (*entity.User, *exception.ErrorResponse){

	user, err := userRepo.GetOne(id);
	
	if(err != nil){
		return nil, &exception.ErrorResponse{
			Status: fiber.StatusInternalServerError,
			Message: err.Error(),
		}
	}

	if(user.ID == ""){
		return nil, &exception.ErrorResponse{
			Status: fiber.StatusNotFound,
			Message: "No se encontró el usuario",
		}
	}

	deletedUser, errDelete := userRepo.Delete(user)

	if(errDelete != nil){
		return nil, &exception.ErrorResponse{
			Status: fiber.StatusInternalServerError,
			Message: err.Error(),
		}
	}

	return deletedUser, nil

}