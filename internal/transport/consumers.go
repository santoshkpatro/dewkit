package transport

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
	"github.com/redis/go-redis/v9"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // dev only
	},
}

func ChatConsumer(c echo.Context) error {
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
	if err := json.Unmarshal([]byte(sessionVal), &chatSession); err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"error": "failed to initiate chat.",
		})
	}

	ws, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		return err
	}
	defer ws.Close()

	conversationChannel := fmt.Sprintf(
		"project:%s:conversation:%s",
		chatSession.ProjectId,
		chatSession.ConversationId,
	)

	sub := cache.Subscribe(ctx, conversationChannel)
	defer sub.Close()

	// Listen and forward Redis → WebSocket
	for {
		select {
		case msg, ok := <-sub.Channel():
			if !ok {
				return nil // Redis subscription closed
			}
			if err := ws.WriteMessage(websocket.TextMessage, []byte(msg.Payload)); err != nil {
				return nil // client disconnected
			}

		case <-ctx.Done():
			return nil // request cancelled
		}
	}
}
