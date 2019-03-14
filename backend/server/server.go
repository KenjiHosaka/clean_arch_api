package server

import (
	"clean_arch_api/backend/db/misc"
	"clean_arch_api/backend/environment"
	"clean_arch_api/backend/fork/golang/sync/singleflight"
	mw "clean_arch_api/backend/server/middleware"
	"clean_arch_api/backend/server/router"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func SetUp(db *gorm.DB, requestBag *singleflight.Group) *echo.Echo {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(mw.SetRequestBagToContext(requestBag))
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{environment.OriginSiteUrl(), environment.AdminSiteUrl()},
		AllowMethods: []string{echo.GET, echo.PUT, echo.POST, echo.DELETE, echo.PATCH},
	}))

	misc.Migration(db)

	router.Routing(e, db)
	return e
}
