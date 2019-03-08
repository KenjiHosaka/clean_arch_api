package middleware

import (
	"clean_arch_api/backend/fork/golang/sync/singleflight"

	"github.com/labstack/echo"
)

func SetRequestBagToContext(requestBag *singleflight.Group) echo.MiddlewareFunc {
	return func (handlerFunc echo.HandlerFunc) echo.HandlerFunc{
		return func(c echo.Context) error {
			c.Set("REQUEST_BAG", requestBag)
			return handlerFunc(c)
		}
	}
}
