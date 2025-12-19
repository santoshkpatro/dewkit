package handlers

import (
	"dewkit/config"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/labstack/echo/v4"
)

type AppContext struct {
	DB *pgxpool.Pool
}

func RegisterAPIRoutes(api *echo.Group) {
	app := AppContext{
		DB: config.DB,
	}

	api.POST("/auth/login", app.LoginHandler)
}