package transport

import "time"

type ChatInitiateRequest struct {
	ProjectCode      string `json:"projectCode"`
	CustomerEmail    string `json:"customerEmail"`
	CustomerFullName string `json:"customerFullName"`
	Message          string `json:"message"`
}

type ChatSession struct {
	ProjectId      int       `json:"projectId"`
	SessionId      string    `json:"sessionId"`
	ConversationId int       `json:"conversationId"`
	ExpiresAt      time.Time `json:"expiresAt"`
}
