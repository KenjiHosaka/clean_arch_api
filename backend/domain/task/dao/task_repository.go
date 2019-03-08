package dao

import (
	"clean_arch_api/backend/domain/task"
	"clean_arch_api/backend/domain/task/entity"
	"clean_arch_api/backend/domain/task/model"
	"clean_arch_api/backend/util/errors"
	"github.com/jinzhu/gorm"
)

type Task struct {
	Database *gorm.DB
}

func (t *Task) FetchTaskByTaskId(taskId uint) (entity entity.TaskEntity, err errors.Error) {
	q := model.Task{ ID: taskId }
	rec, err := q.GetById(t.Database); if err != nil {
		return
	}
	err = entity.Mapping(rec); if err != nil {
		return
	}
	return
}

func (t *Task) Store(data task.StoreData) (entity entity.TaskEntity, err errors.Error) {
	newRec := model.Task{
		Title   : data.Title,
		Contents: data.Contents,
	}

	tx := t.Database.Begin()
	err = newRec.Create(tx); if err != nil {
		tx.Rollback()
		return
	}
	err = entity.Mapping(newRec); if err != nil {
		tx.Rollback()
		return
	}
	tx.Commit()
	return
}