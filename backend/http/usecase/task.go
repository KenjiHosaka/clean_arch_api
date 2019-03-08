package usecase

import (
	"clean_arch_api/backend/domain/task"
	"clean_arch_api/backend/domain/task/entity"
	"clean_arch_api/backend/util/errors"
)

type Task struct {
	TaskRepo task.Repository
}

func NewTaskUsecase(taskRepo task.Repository) *Task {
	return &Task {
		taskRepo,
	}
}

func (usecase Task) GetTask(taskId uint) (entity.TaskEntity, errors.Error)  {
	return usecase.TaskRepo.FetchTaskByTaskId(taskId)
}

func (usecase Task) StoreNewTask(title, contents string) (entity.TaskEntity, errors.Error)  {
	return usecase.TaskRepo.Store(task.StoreData{Title: title, Contents: contents})
}