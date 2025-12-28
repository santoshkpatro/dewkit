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

func (s *Service) ConversationMessages(conversationId int) {

}
