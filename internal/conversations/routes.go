package conversations

import (
	"dewkit/config/middlewares"

	"github.com/labstack/echo/v4"
)

func RegisterAPIRoutes(g *echo.Group) {
	g.Use(middlewares.LoggedInMiddleware)

	g.GET("/conversations", ConversationListHandler, middlewares.ProjectPermissionMiddleware)
}
