package taskCases

import (
	"github.com/brianpl990227/hello-fiber/domain/dto"
	"github.com/brianpl990227/hello-fiber/domain/entity"
	"github.com/brianpl990227/hello-fiber/repositories/taskRepo"
	"github.com/google/uuid"
)

func CreateTaskHandler(dto *dto.CreateTaskDTO) (*entity.Task, error) {

	entity := entity.Task{

		Title:       dto.Title,
		Description: dto.Description,
		Done:        dto.Done,
		UserID:      dto.UserID,
		ID: uuid.NewString(),
	}

	task, errCreate := taskrepo.CreateTask(&entity)

	if errCreate != nil {
		return nil, errCreate
	}

	return task, nil

}
