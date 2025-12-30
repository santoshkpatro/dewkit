package conversations

import (
	"dewkit/config"

	"github.com/jmoiron/sqlx"
	"github.com/redis/go-redis/v9"
)

type Service struct {
	DB    *sqlx.DB
	Cache *redis.Client
}

func NewService() *Service {
	db := config.DB
	cache := config.Cache
	return &Service{DB: db, Cache: cache}
}

func (s *Service) ListActiveConversations(projectId int, status string) ([]ConversationListResponse, error) {
	query := `
		SELECT *
		FROM (
			SELECT DISTINCT ON (c.id)
				c.id,
				c.status,
				c.customer_email,
				c.customer_full_name,
				c.created_at,

				m.id         AS last_message_id,
				m.body       AS last_message_body,
				m.created_at AS last_message_created_at

			FROM conversations c
			LEFT JOIN messages m
				ON m.conversation_id = c.id

			WHERE c.project_id = $1
			  AND c.status = $2

			ORDER BY
				c.id,
				m.created_at DESC
		) t
		ORDER BY
			last_message_created_at DESC NULLS LAST;
	`

	conversations := []ConversationListResponse{}
	err := s.DB.Select(&conversations, query, projectId, status)
	if err != nil {
		return nil, err
	}

	return conversations, nil
}

func (s *Service) ConversationMessages(conversationId int) ([]MessageResponse, error) {
	query := `
		SELECT 
			m.id,
			m.sender_type,
			m.body,
			m.created_at
		FROM messages m
		WHERE m.conversation_id = $1
		AND m.is_internal = false
		ORDER BY m.created_at ASC
	`

	messages := []MessageResponse{}
	err := s.DB.Select(&messages, query, conversationId)
	if err != nil {
		return nil, err
	}

	return messages, nil

}

func (s *Service) CreateConversationMessage(conversationId int, userId int, message MessageRequest) (MessageResponse, error) {
	query := `
		INSERT INTO messages
		(conversation_id, sender_type, sender_staff_id, body)
		VALUES ($1, $2, $3, $4)
		RETURNING id;
	`
	tx, err := s.DB.Beginx()
	if err != nil {
		return MessageResponse{}, err
	}
	defer func() {
		if err != nil {
			_ = tx.Rollback()
		}
	}()

	var messageId int
	err = tx.QueryRowx(query, conversationId, message.SenderType, userId, message.Body).Scan(&messageId)
	if err != nil {
		return MessageResponse{}, err
	}
	if err = tx.Commit(); err != nil {
		return MessageResponse{}, err
	}

	var newMessage MessageResponse
	err = s.DB.Get(&newMessage, "SELECT id, body, sender_type, created_at FROM messages WHERE id = $1;", messageId)
	if err != nil {
		return MessageResponse{}, err
	}

	return newMessage, nil
}
