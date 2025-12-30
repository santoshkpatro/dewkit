package conversations

import (
	"dewkit/config/middlewares"

	"github.com/labstack/echo/v4"
)

func RegisterAPIRoutes(g *echo.Group) {
	g.Use(middlewares.LoggedInMiddleware)
	g.Use(middlewares.ProjectPermissionMiddleware)

	g.GET("/conversations", ConversationListHandler)
	g.GET("/conversations/:conversationId/messages", ConversationMessageListHandler)
	g.POST("/conversations/:conversationId/messages", ConversationMessageCreateHandler)
}
