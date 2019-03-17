package handler

import (
	"clean_arch_api/backend/http/handler/response"
	"github.com/labstack/echo/v4"
	"net/http"
)

func WelcomeMessage(c echo.Context) error {
	return c.JSON(http.StatusOK, response.SuccessResult("Welcome to sample !!"))
}
