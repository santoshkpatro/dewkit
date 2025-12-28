package conversations

import (
	"log/slog"
	"net/http"

	"github.com/labstack/echo/v4"
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

// func ConversationMessageListHandler(c echo.Context) error {
// 	projectID := c.Get("project_id").(int)
// 	conversationIDStr := c.Param("conversationId")
// 	conversationID, err := strconv.Atoi(conversationIDStr)
// 	if err != nil {
// 		return c.JSON(http.StatusBadRequest, map[string]string{
// 			"error": "invalid conversationId",
// 		})
// 	}

// 	service := NewService()

// }
