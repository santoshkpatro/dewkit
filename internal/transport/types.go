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

type MessageRequest struct {
	Body string `json:"body"`
}

type MessageResponse struct {
	ID         int       `json:"id" db:"id"`
	Body       string    `json:"body" db:"body"`
	SenderType string    `json:"senderType" db:"sender_type"`
	CreatedAt  time.Time `json:"createdAt" db:"created_at"`
}

type ConversationMessageResponse struct {
	ConversationId int             `json:"conversationId"`
	Message        MessageResponse `json:"message"`
}
