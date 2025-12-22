package middlewares

import (
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
)

func DBMiddleware(db *sqlx.DB) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Set("db", db)
			return next(c)
		}
	}
}
