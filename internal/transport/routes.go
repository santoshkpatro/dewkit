package transport

import "github.com/labstack/echo/v4"

func RegisterAPIRoutes(g *echo.Group) {
	g.POST("/chat/initiate", ChatInitiateHandler)
	g.POST("/chat/message", ChatMessageSend)
}

func RegisterWSRoutes(g *echo.Group) {
	g.GET("/chat", ChatConsumer)
}
