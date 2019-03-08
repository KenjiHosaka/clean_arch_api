package misc

import (
	"time"
	"github.com/jinzhu/gorm"
	"clean_arch_api/backend/environment"
)

func Migration(db *gorm.DB) {
	// バージョン管理のテーブルがあればmigrateするか確認
	migrationVersion := environment.MigrationVesion()
	if db.HasTable(SchemeVersion{}) {
		schemeVersion, err := schemeVersion(db)
		if err != nil || (migrationVersion <= schemeVersion) {
			return
		}
	}

	tx := db.Begin()
	var err error
	switch migrationVersion {
	case 1:
		err = migration001(tx)
	}

	if err != nil {
		tx.Rollback()
		return
	}
	// Scheme version更新
	updateSchemeVersion(tx, uint(migrationVersion), time.Now())
	tx.Commit()
}

type SchemeVersion struct {
	Version       uint      `gorm:"PRIMARY_KEY;type:int"`
	MigrationDate time.Time `gorm:"not null" json:"migration_date"`
}

func updateSchemeVersion(db *gorm.DB, version uint, migrationDate time.Time) error {
	values := SchemeVersion {
		Version       : version,
		MigrationDate : migrationDate,
	}
	err := db.Create(&values).Error
	return err
}

func schemeVersion(db *gorm.DB) (int, error) {
	result := SchemeVersion{}
	db = db.Last(&result)
	if db.RecordNotFound() {
		return 0, nil
	}
	return int(result.Version), db.Error
}
