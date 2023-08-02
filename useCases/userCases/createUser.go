package userCases

import (
	"github.com/brianpl990227/hello-fiber/domain/dto"
	"github.com/brianpl990227/hello-fiber/domain/entity"
	"github.com/brianpl990227/hello-fiber/repositories/userRepo"
	"github.com/google/uuid"
)

func CreateUserHandler(dto *dto.CreateUserDTO) (*entity.User, error) {

	user := entity.User{
		Id: uuid.NewString(),
		Firstname: dto.Firstname,
		Lastname: dto.Lastname,
	}
	
	createdUser, errRepo := userRepo.Add(&user)

	if errRepo != nil{
		return nil, errRepo
	}
	
	return createdUser, nil;

}