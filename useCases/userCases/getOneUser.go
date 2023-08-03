package userCases

import (
	"github.com/brianpl990227/hello-fiber/domain/entity"
	"github.com/brianpl990227/hello-fiber/repositories/userRepo"
)


func GetOneUserHandler(id string) (*entity.User, error){
	
	user, err := userRepo.GetOne(id)

	if(err != nil){
		return nil, err
	}

	return user, nil

}