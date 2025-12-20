package handlers

import (
	"dewkit/internal/services/auth"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (app *AppContext) LoginHandler(c echo.Context) error {
	var req auth.LoginRequest

	if err := c.Bind(&req); err != nil {
		return err
	}

	fmt.Println("User-", req.Email, req.Password)

	return c.JSON(http.StatusOK, map[string]string{"hello": "world"})
}
