package transport

import "github.com/labstack/echo/v4"

func RegisterAPIRoutes(g *echo.Group) {
	g.POST("/chat/initiate", ChatInitiateHandler)
}

func RegisterWSRoutes(g *echo.Group) {
	g.GET("/chat", ChatConsumer)
}
