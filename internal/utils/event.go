package utils

import (
	"dewkit/internal/models"
	"encoding/json"
	"errors"
	"reflect"
	"time"
)

func BuildEvent(eventType string, data interface{}) (string, error) {

	if data == nil {
		return "", errors.New("event data cannot be nil")
	}

	t := reflect.TypeOf(data)

	// Allow struct or pointer to struct
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	if t.Kind() != reflect.Struct {
		return "", errors.New("event data must be a struct or pointer to struct")
	}

	payload, err := json.Marshal(data)
	if err != nil {
		return "", err
	}

	event := models.Event{
		Type:      eventType,
		Payload:   payload,
		Timestamp: time.Now().UTC(),
	}

	bytes, err := json.Marshal(event)
	if err != nil {
		return "", err
	}

	return string(bytes), nil
}
