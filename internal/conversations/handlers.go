package conversations

import (
	"dewkit/internal/models"
	"dewkit/internal/utils"
	"fmt"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/redis/go-redis/v9"
)

func ConversationListHandler(c echo.Context) error {
	projectID := c.Get("project_id").(int)
	status := c.QueryParam("status")

	service := NewService()
	conversations, err := service.ListActiveConversations(projectID, status)
	if err != nil {
		slog.Error("failed to list conversation", "err", err)
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"error": "Failed to list conversations",
		})
	}

	return c.JSON(http.StatusOK, conversations)
}

func ConversationMessageListHandler(c echo.Context) error {
	// projectID := c.Get("project_id").(int)
	conversationIDStr := c.Param("conversationId")
	conversationID, err := strconv.Atoi(conversationIDStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "invalid conversationId",
		})
	}

	service := NewService()
	messages, err := service.ConversationMessages(conversationID)
	if err != nil {
		slog.Error("failed to list messages", "err", err)
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"error": "Failed to list messages",
		})
	}

	return c.JSON(http.StatusOK, messages)
}

func ConversationMessageCreateHandler(c echo.Context) error {
	ctx := c.Request().Context()
	cache := c.Get("cache").(*redis.Client)
	projectID := c.Get("project_id").(int)

	userId := c.Get("user_id").(int)
	conversationIDStr := c.Param("conversationId")
	conversationID, err := strconv.Atoi(conversationIDStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "invalid conversationId",
		})
	}
	var req MessageRequest

	if err := c.Bind(&req); err != nil {
		slog.Debug(
			"failed to bind message reqyest",
			"err", err,
		)

		return c.JSON(http.StatusBadRequest, echo.Map{
			"error": "Invalid request body",
		})
	}

	service := NewService()
	newMessage, err := service.CreateConversationMessage(conversationID, userId, req)

	if err != nil {
		slog.Debug(
			"failed to create message",
			"err", err,
			"path", c.Path(),
		)

		return c.JSON(http.StatusBadGateway, echo.Map{
			"error": "failed to create message",
		})
	}

	conversationMessage := ConversationMessageResponse{
		ConversationId: conversationID,
		Message:        newMessage,
	}

	messageEvent, _ := utils.BuildEvent(models.EventMessageNew, newMessage)
	conversationMessageEvent, _ := utils.BuildEvent(models.EventMessageNew, conversationMessage)

	imboxChannel := fmt.Sprintf("project:%d:imbox", projectID)
	conversationChannel := fmt.Sprintf("project:%d:conversation:%d", projectID, conversationID)

	cache.Publish(ctx, conversationChannel, messageEvent)
	cache.Publish(ctx, imboxChannel, conversationMessageEvent)

	return c.JSON(http.StatusOK, newMessage)
}
