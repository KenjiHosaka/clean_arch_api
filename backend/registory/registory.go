package registory

import (
	"github.com/jinzhu/gorm"
)

func Inject(database * gorm.DB) Usecase {
	return NewUsecase(NewRepository(database))
}
