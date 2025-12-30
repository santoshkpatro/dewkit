package transport

import (
	"dewkit/internal/models"
	"dewkit/internal/utils"
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/redis/go-redis/v9"
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

func ChatMessageSend(c echo.Context) error {
	ctx := c.Request().Context()
	sessionId := c.QueryParam("sessionId")
	cache := c.Get("cache").(*redis.Client)

	key := fmt.Sprintf("chat:%s:session", sessionId)

	sessionVal, err := cache.Get(ctx, key).Result()
	if err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{
			"error": "No session found.",
		})
	}

	var chatSession ChatSession
	err = json.Unmarshal([]byte(sessionVal), &chatSession)
	if err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{
			"error": "failed to send chat.",
		})
	}

	var req MessageRequest

	if err := c.Bind(&req); err != nil {
		slog.Debug(
			"failed to bind request",
			"err", err,
			"path", c.Path(),
		)

		return c.JSON(http.StatusBadRequest, echo.Map{
			"error": "Invalid request body",
		})
	}

	service := NewService()
	newMessage, err := service.SendMessage(chatSession.ConversationId, req)
	if err != nil {
		slog.Debug(
			"failed to send message",
			"err", err,
		)
		return c.JSON(http.StatusNotFound, echo.Map{
			"error": "failed to send chat.",
		})
	}
	conversationMessage := ConversationMessageResponse{
		ConversationId: chatSession.ConversationId,
		Message:        newMessage,
	}

	messageEvent, _ := utils.BuildEvent(models.EventMessageNew, newMessage)
	conversationMessageEvent, _ := utils.BuildEvent(models.EventMessageNew, conversationMessage)

	imboxChannel := fmt.Sprintf("project:%d:imbox", chatSession.ProjectId)
	conversationChannel := fmt.Sprintf("project:%d:conversation:%d", chatSession.ProjectId, chatSession.ConversationId)

	cache.Publish(ctx, conversationChannel, messageEvent)
	cache.Publish(ctx, imboxChannel, conversationMessageEvent)

	return c.JSON(http.StatusOK, newMessage)
}
