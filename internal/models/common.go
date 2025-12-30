package models

import (
	"encoding/json"
	"time"
)

const (
	EventMessageNew = "message.new"
)

type Event struct {
	Type      string          `json:"type"`
	Payload   json.RawMessage `json:"payload"`
	Timestamp time.Time       `json:"timestamp"`
}
