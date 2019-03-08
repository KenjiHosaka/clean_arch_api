package registory

import (
	"clean_arch_api/backend/domain/task"
	"clean_arch_api/backend/domain/task/dao"
	"github.com/jinzhu/gorm"
)

type Repository interface {
	NewTaskRepository() task.Repository
}

type repositoryImp struct {
	database *gorm.DB
}

func NewRepository(database *gorm.DB) Repository {
	return &repositoryImp{database}
}
func (r *repositoryImp) NewTaskRepository() task.Repository {
	return &dao.Task{Database: r.database}
}