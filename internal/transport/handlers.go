package transport

import (
	"log/slog"
	"net/http"

	"github.com/labstack/echo/v4"
)

func ChatInitiateHandler(c echo.Context) error {
	ctx := c.Request().Context()
	var req ChatInitiateRequest

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "invalid request body"})
	}

	service := NewService()
	chatSession, err := service.NewChatSession(ctx, req)
	if err != nil {
		slog.Error("failed to initiate chat", "err", err)
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"error": "Failed to initiate chat",
		})
	}

	return c.JSON(http.StatusOK, echo.Map{"sessionId": chatSession.SessionId})
}
