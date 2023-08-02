package userRepo

import (
	"errors"
	"math/rand"
	"github.com/brianpl990227/hello-fiber/domain/entity"
)

var users entity.Users


func Add(user *entity.User) (*entity.User, error){

	aux := rand.Intn(10)
	
	if aux % 2 == 0 {
		return nil, errors.New("Error al conectar con la base de datos")
	}

	users = append(users, *user)
	
	return user, nil

}

