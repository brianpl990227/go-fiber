package userCases

import (
	"github.com/brianpl990227/hello-fiber/domain/entity"
	"github.com/brianpl990227/hello-fiber/repositories/userRepo"
)


func GetAllUsersHandler() (*entity.Users, error){

	users, errRepo := userRepo.GetAll()

	if errRepo != nil{
		return nil, errRepo
	}

	return users, nil

}