package model

import (
	"clean_arch_api/backend/util/errors"
	"github.com/jinzhu/gorm"
	"time"
)

const (
	NotStarted = iota
	Doing
	Finished
	Closed
)

type Task struct {
	ID        uint      `gorm:"primary_key;AUTO_INCREMENT" json:"id"`
	Title     string    `gorm:"type:varchar(100);not null" json:"title"`
	Contents  string    `gorm:"type:text" json:"contents"`
	Status    uint      `gorm:"not null" json:"status"`
	CreatedAt time.Time `gorm:"not null;default:CURRENT_TIMESTAMP" json:"createdAt"`
	UpdatedAt time.Time `gorm:"not null;default:CURRENT_TIMESTAMP" json:"updatedAt"`
}

func (self *Task) GetById(database *gorm.DB) (result Task, err errors.Error) {
	if self.ID < 1 {
		err = errors.NewBadRequestError()
		return
	}

	if database.Where("id = ?", self.ID).
		Find(&result).Error != nil {
		err = errors.NewInternalServerError()
	}
	return
}

func (self *Task) Create(database *gorm.DB) errors.Error {
	if self.Title == "" {
		return errors.NewBadRequestError()
	}
	if database.Create(self).Error != nil {
		return errors.NewInternalServerError()
	}
	return nil
}