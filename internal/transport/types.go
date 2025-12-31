package transport

import "time"

type ChatInitiateRequest struct {
	ProjectCode      string `json:"projectCode"`
	CustomerEmail    string `json:"customerEmail"`
	CustomerFullName string `json:"customerFullName"`
	Message          string `json:"message"`
}

type ChatSession struct {
	ProjectId      string    `json:"projectId"`
	SessionId      string    `json:"sessionId"`
	ConversationId string    `json:"conversationId"`
	ExpiresAt      time.Time `json:"expiresAt"`
}

type MessageRequest struct {
	Body string `json:"body"`
}

type MessageResponse struct {
	ID         string    `json:"id" db:"id"`
	Body       string    `json:"body" db:"body"`
	SenderType string    `json:"senderType" db:"sender_type"`
	CreatedAt  time.Time `json:"createdAt" db:"created_at"`
}

type ConversationMessageResponse struct {
	ConversationId string          `json:"conversationId"`
	Message        MessageResponse `json:"message"`
}
