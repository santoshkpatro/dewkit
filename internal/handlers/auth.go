package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (app *AppContext) LoginHandler(c echo.Context) error {
	return c.String(http.StatusOK, "Hello")
}