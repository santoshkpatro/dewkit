package conversations

import "time"

type LastMessageResponse struct {
	ID        int       `db:"last_message_id" json:"id"`
	Body      string    `db:"last_message_body" json:"body"`
	CreatedAt time.Time `db:"last_message_created_at" json:"createdAt"`
}

type ConversationListResponse struct {
	ID               int                 `db:"id" json:"id"`
	Status           string              `db:"status" json:"status"`
	CustomerEmail    string              `db:"customer_email" json:"customerEmail"`
	CustomerFullName string              `db:"customer_full_name" json:"customerFullName"`
	CreatedAt        time.Time           `db:"created_at" json:"createdAt"`
	LastMessage      LastMessageResponse `db:",inline" json:"lastMessage"`
}

type MessageResponse struct {
	ID         int    `json:"id" db:"id"`
	Body       string `json:"body" db:"body"`
	SenderType string `json:"senderType" db:"sender_type"`
}
