package misc

import (
	"clean_arch_api/backend/domain/task/model"
	"github.com/jinzhu/gorm"
)

func migration001(tx *gorm.DB) error {
	// テーブル作成
	err := tx.AutoMigrate(
		&model.Task{},
		&SchemeVersion{},
	).Error

	if err != nil {
		return err
	}

	return err
}
