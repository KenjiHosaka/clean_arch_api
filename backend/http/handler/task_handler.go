package handler

import (
	usecase "clean_arch_api/backend/http"
	"clean_arch_api/backend/http/handler/request"
	"clean_arch_api/backend/http/handler/response"
	"clean_arch_api/backend/registory"
	"clean_arch_api/backend/util"
	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
	"net/http"
)

type TaskHandler struct {
	taskUsecase usecase.TaskUsecase
}

func NewTaskHandler(u registory.Usecase) TaskHandler {
	return TaskHandler{u.NewTaskUsecase()}
}

func (th TaskHandler) NewTask(c echo.Context) error {

	body := request.NewTaskRequest{}
	e := c.Bind(&body)

	if e != nil {
		log.Error(e.Error())
		return c.JSON(http.StatusBadRequest, response.ErrorResult(e.Error()))
	}

	entity, err := th.taskUsecase.StoreNewTask(body.Title, body.Contents)

	if err != nil {
		log.Error(err.Message())
		return c.JSON(err.Code(), response.ErrorResult(err.Message()))
	}

	return c.JSON(http.StatusOK, response.SuccessResult(entity))
}

func (th TaskHandler) GetTask(c echo.Context) error {

	taskId, e:= util.GetUintParam(c, "taskId")

	if e != nil {
		log.Error(e.Error())
		return c.JSON(http.StatusBadRequest, response.ErrorResult(e.Error()))
	}

	entity, err := th.taskUsecase.GetTask(taskId)

	if err != nil {
		log.Error(err.Message())
		return c.JSON(err.Code(), response.ErrorResult(err.Message()))
	}

	return c.JSON(http.StatusOK, response.SuccessResult(entity))
}