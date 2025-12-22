package auth

import (
	"log/slog"
	"net/http"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
)

func LoginHandler(c echo.Context) error {
	var req LoginRequest

	if err := c.Bind(&req); err != nil {
		slog.Debug(
			"failed to bind login request",
			"err", err,
			"path", c.Path(),
		)

		return c.JSON(http.StatusBadRequest, echo.Map{
			"error": "Invalid request body",
		})
	}

	if err := c.Validate(&req); err != nil {
		slog.Debug(
			"login request validation failed",
			"email", req.Email,
			"err", err,
		)

		return c.JSON(http.StatusUnprocessableEntity, echo.Map{
			"error": "Invalid request",
		})
	}

	authService := NewService()
	if err := authService.Authenticate(req.Email, req.Password); err != nil {
		slog.Warn(
			"authentication failed",
			"email", req.Email,
			"err", err,
		)

		return c.JSON(http.StatusUnauthorized, echo.Map{
			"error": "Invalid credentials",
		})
	}

	db := c.Get("db").(*sqlx.DB)

	var loggedInUser LoggedInUserResponse
	if err := db.Get(
		&loggedInUser,
		`SELECT email, full_name, role FROM users WHERE email = $1`,
		req.Email,
	); err != nil {
		slog.Error(
			"failed to fetch logged in user",
			"email", req.Email,
			"err", err,
		)

		return c.JSON(http.StatusInternalServerError, echo.Map{
			"error": "Failed to fetch logged in user information",
		})
	}

	slog.Info(
		"user logged in successfully",
		"email", loggedInUser.Email,
	)

	return c.JSON(http.StatusOK, loggedInUser)
}
