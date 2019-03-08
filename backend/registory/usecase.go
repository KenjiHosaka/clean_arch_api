package registory

import (
	"clean_arch_api/backend/http"
	"clean_arch_api/backend/http/usecase"
)

type Usecase interface {
	NewTaskUsecase() http.TaskUsecase
}

type usecaseImp struct {
	Repository
}

func NewUsecase(repository Repository) Usecase {
	return &usecaseImp{repository}
}

func (u *usecaseImp) NewTaskUsecase() http.TaskUsecase {
	return &usecase.Task{u.Repository.NewTaskRepository()}
}