package res

import (
	"clean_arch_api/backend/domain/task/entity"
)

type TaskEntity struct {
	Status string `json:"status" example:"success"`
	Payload entity.TaskEntity `json:"payload"`
}