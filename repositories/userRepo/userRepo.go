package userRepo

import (
	"errors"

	"github.com/brianpl990227/hello-fiber/db"
	"github.com/brianpl990227/hello-fiber/domain/entity"
)



func Add(user *entity.User) (*entity.User, error) {

	createdUser := db.Context.Create(user)
	
	err := createdUser.Error

	if(err != nil){
		return nil, errors.New(err.Error());
	}


	return user, nil

}

func GetAll() (*[]entity.User, error) {

	var users []entity.User

	//Para ordenar de mayor a menor, por defecto es de menor a mayor (sin del desc)
	//Offset es el Skip y Limit es el Take
	getAll := db.Context.Preload("Tasks").Order("firstname desc").Offset(0).Limit(5).Find(&users)
	
	errAll := getAll.Error

	if(errAll != nil){
		return nil, errors.New(errAll.Error())
	}

	return &users, nil

}

func GetOne(id string) (*entity.User, error){
	var user entity.User;

	getOne := db.Context.Find(&user, "id = ?", id)

	if getOne.Error != nil {
		return nil, errors.New(getOne.Error.Error())
	}

	/*
		Buscar un usuario con id == 10 y name == "Brian"
		var user entity.User;
		db.Where("id = ? AND name = ?", 10, "Brian").First(&user)
	*/
	return &user, nil
}

func Delete(user *entity.User) (*entity.User, error){

	//El Unscoped elimina el registro de la BD sin borrado suave cuando esta DeletedAt
	errDelete := db.Context.Delete(user); 
	err := errDelete.Error
	if err != nil{
		return nil, errors.New(err.Error())
	}

	return user, nil;

	

}
