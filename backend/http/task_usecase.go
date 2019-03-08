package http

import (
	"clean_arch_api/backend/domain/task/entity"
	"clean_arch_api/backend/util/errors"
)

type TaskUsecase interface {
	GetTask(taskId uint) (entity.TaskEntity, errors.Error)
	StoreNewTask(title, contents string) (entity.TaskEntity, errors.Error)
}