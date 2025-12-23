package auth

import (
	"dewkit/config/middlewares"

	"github.com/labstack/echo/v4"
)

func RegisterAPIRoutes(g *echo.Group) {
	g.POST("/login", LoginHandler)
	g.GET("/profile", ProfileHandler, middlewares.LoggedInMiddleware)
	g.GET("/meta", MetaHandler)
}
