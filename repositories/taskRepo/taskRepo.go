package taskrepo

import (
	"github.com/brianpl990227/hello-fiber/db"
	"github.com/brianpl990227/hello-fiber/domain/entity"
	"errors"
)

func CreateTask(task *entity.Task) (*entity.Task, error){

	createdTask := db.Context.Create(task)

	err := createdTask.Error

	if(err != nil){
		return nil, errors.New(err.Error());
	}


	return task, nil

}