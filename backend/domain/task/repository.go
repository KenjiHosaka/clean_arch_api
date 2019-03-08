package task

import (
	"clean_arch_api/backend/domain/task/entity"
	"clean_arch_api/backend/util/errors"
)

type Repository interface {
	FetchTaskByTaskId(uint) (entity.TaskEntity, errors.Error)
	Store(StoreData) (entity.TaskEntity, errors.Error)
}

type StoreData struct {
	Title    string
	Contents string
}