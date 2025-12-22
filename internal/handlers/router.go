package handlers

import (
	"dewkit/config"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
)

type AppContext struct {
	DB *sqlx.DB
}

func RegisterAPIRoutes(api *echo.Group) {
	app := AppContext{
		DB: config.DB,
	}

	api.POST("/auth/login", app.LoginHandler)
}
