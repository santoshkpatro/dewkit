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
	err = json.Unmarshal([]byte(sessionVal), &chatSession)
	if err != nil {
		return c.JSON(http.StatusNotFound, echo.Map{
			"error": "failed to initiate chat.",
		})
	}

	ws, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		return err
	}
	defer ws.Close()

	conversationChannel := fmt.Sprintf("project:%d:conversation:%d", chatSession.ProjectId, chatSession.ConversationId)
	imboxChannel := fmt.Sprintf("project:%d:imbox", chatSession.ProjectId)

	// conversationSub := cache.Subscribe(ctx, conversationChannel)
	// defer conversationSub.Close()

	// go func() {
	// 	for msg := range conversationSub.Channel() {
	// 		_ = ws.WriteMessage(websocket.TextMessage, []byte(msg.Payload))
	// 	}
	// }()

	for {
		_, msg, err := ws.ReadMessage()
		if err != nil {
			break
		}

		fmt.Println("Msg: ", string(msg))

		cache.Publish(ctx, conversationChannel, msg)
		cache.Publish(ctx, imboxChannel, msg)
	}

	return nil
}
