package projects

import (
	"dewkit/config/middlewares"

	"github.com/labstack/echo/v4"
)

func RegisterAPIRoutes(g *echo.Group) {
	g.GET("", ProjectListHandler, middlewares.LoggedInMiddleware)
	g.POST("", ProjectCreateHandler, middlewares.LoggedInMiddleware)
}
