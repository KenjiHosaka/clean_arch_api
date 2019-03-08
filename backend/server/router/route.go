package router

import (
	_ "clean_arch_api/backend/docs"
	"clean_arch_api/backend/http/handler"
	"clean_arch_api/backend/registory"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"github.com/swaggo/echo-swagger"
)

func Routing(r *echo.Echo, database * gorm.DB) {
	// rootAPIを用意
	r.GET("/", handler.WelcomeMessage)

	uc := registory.Inject(database)
	taskHandler := handler.NewTaskHandler(uc)

	taskRoute := r.Group("/tasks")
	taskRoute.POST("", taskHandler.NewTask)
	taskRoute.GET("/:taskId", taskHandler.GetTask)

	// swagger
	r.GET("/swagger/*", echoSwagger.WrapHandler)
}
