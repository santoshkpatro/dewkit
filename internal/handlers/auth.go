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
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request body."})
	}

	if err := c.Validate(&req); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, map[string]string{"error": "Invalid request."})
	}

	fmt.Println("User- ", req.Email, req.Password)

	return c.JSON(http.StatusOK, map[string]string{"hello": "world"})
}
