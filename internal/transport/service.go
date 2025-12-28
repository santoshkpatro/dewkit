package transport

import (
	"context"
	"dewkit/config"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
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

func (s *Service) NewChatSession(ctx context.Context, chat ChatInitiateRequest) (*ChatSession, error) {
	project, err := s.GetProjectFromProjectCode(chat.ProjectCode)
	if err != nil {
		return nil, err
	}

	tx, err := s.DB.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	defer func() {
		if err != nil {
			_ = tx.Rollback()
		}
	}()

	var conversationId int
	conversationQuery := `
		INSERT INTO 
		conversations (customer_email, customer_full_name, status, project_id) 
		VALUES ($1, $2, $3, $4)
		RETURNING id
	`
	err = tx.QueryRow(conversationQuery, chat.CustomerEmail, chat.CustomerFullName, "open", project.ID).Scan(&conversationId)
	if err != nil {
		return nil, err
	}
	messageQuery := `
		INSERT INTO messages
			(conversation_id, sender_type, body)
		VALUES ($1, $2, $3)
	`
	_, err = tx.ExecContext(
		ctx,
		messageQuery,
		conversationId,
		"customer",
		chat.Message,
	)
	if err != nil {
		return nil, err
	}
	if err = tx.Commit(); err != nil {
		return nil, err
	}

	sessionId := strings.ReplaceAll(uuid.New().String(), "-", "")
	newChatSession := ChatSession{
		ConversationId: conversationId,
		SessionId:      sessionId,
		ProjectId:      project.ID,
		ExpiresAt:      time.Now().Add(30 * time.Minute),
	}

	key := fmt.Sprintf("chat:%s:session", newChatSession.SessionId)
	sessionData, err := json.Marshal(newChatSession)
	if err != nil {
		return nil, err
	}

	err = s.Cache.SetEx(ctx, key, sessionData, time.Until(newChatSession.ExpiresAt)).Err()
	if err != nil {
		return nil, err
	}

	return &newChatSession, nil
}

func (s *Service) GetProjectFromProjectCode(projectCode string) (*Project, error) {
	project := Project{}

	query := `
		SELECT id, name, code FROM projects WHERE code = $1
	`
	err := s.DB.Get(&project, query, projectCode)
	if err != nil {
		return nil, err
	}

	return &project, nil
}
