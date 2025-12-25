package projects

import (
	"fmt"
	"log/slog"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"github.com/redis/go-redis/v9"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // dev only
	},
}

func ImboxConsumer(c echo.Context) error {
	// userID := c.Get("user_id").(int)
	// projectID := c.Get("project_id").(int)
	projectID := c.Param("projectId")

	ctx := c.Request().Context()
	cache := c.Get("cache").(*redis.Client)
	db := c.Get("db").(*sqlx.DB)

	ws, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		return err
	}
	defer ws.Close()

	var projectCode string
	query := `
		SELECT code FROM projects WHERE id = $1
	`
	err = db.Get(&projectCode, query, projectID)
	if err != nil {
		slog.Warn("No project found with the given project id")
		c.JSON(http.StatusNotFound, echo.Map{"error": "No project found."})
	}

	channel := fmt.Sprintf("project:%s:imbox", projectCode)
	sub := cache.Subscribe(ctx, channel)
	defer sub.Close()

	go func() {
		for msg := range sub.Channel() {
			_ = ws.WriteMessage(websocket.TextMessage, []byte(msg.Payload))
		}
	}()

	for {
		_, msg, err := ws.ReadMessage()
		if err != nil {
			break
		}

		cache.Publish(ctx, channel, msg)
	}

	return nil
}
