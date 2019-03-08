package entity

import (
	"clean_arch_api/backend/domain/task/model"
	"clean_arch_api/backend/util/errors"
)

type TaskEntity struct {
	TaskID      uint   `json:"taskId"`
	Title       string `json:"title"`
	Contents    string `json:"contents"`
	Status      uint   `json:"status"`
	StatusLabel string `json:"statusLabel"`
}

type Status uint

func (status Status) label() string {
	switch status {
	case model.NotStarted:
		return "not started"
	case model.Doing:
		return "doing"
	case model.Finished:
		return "finished"
	case model.Closed:
		return "closed"
	default:
		return "unknown"
	}
}

func (entity *TaskEntity) Mapping(task model.Task) errors.Error {
	entity.TaskID = task.ID
	entity.Title = task.Title
	entity.Contents = task.Contents
	status := Status(task.Status)
	entity.Status = uint(status)
	entity.StatusLabel = status.label()
	return nil
}