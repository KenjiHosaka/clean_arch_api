package util

import (
	"github.com/labstack/echo"
	"strconv"
)

func GetUintParam(c echo.Context, name string) (uint, error) {
	val, err := strconv.Atoi(c.Param(name))
	return uint(val), err
}
