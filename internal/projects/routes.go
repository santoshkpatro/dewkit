package projects

import (
	"dewkit/config/middlewares"

	"github.com/labstack/echo/v4"
)

func RegisterAPIRoutes(g *echo.Group) {
	g.Use(middlewares.LoggedInMiddleware)

	g.GET("/:projectId/members", ProjectMembersHandler, middlewares.ProjectPermissionMiddleware)
	g.GET("", ProjectListHandler)
	g.POST("", ProjectCreateHandler)
}

func RegisterWSRoutes(g *echo.Group) {
	g.Use(middlewares.LoggedInMiddleware)

	g.GET("/:projectId/imbox", ImboxConsumer, middlewares.ProjectPermissionMiddleware)
}
