package auth

import (
	"fmt"
	"net/http"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
)

func LoginHandler(c echo.Context) error {
	var req LoginRequest

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request body."})
	}

	if err := c.Validate(&req); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, map[string]string{"error": "Invalid request."})
	}

	authService := NewService()
	err := authService.Authenticate(req.Email, req.Password)
	if err != nil {
		fmt.Println("Erro", err)
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid credentials"})
	}

	db := c.Get("db").(*sqlx.DB)
	var loggedInUser LoggedInUserResponse
	err = db.Get(&loggedInUser, `SELECT email, full_name, role FROM users WHERE email = $1`, req.Email)
	if err != nil {
		fmt.Println("Erro", err)
		return c.JSON(http.StatusBadGateway, map[string]string{"error": "Failed to fetch logged in user information"})
	}

	return c.JSON(http.StatusOK, loggedInUser)
}
