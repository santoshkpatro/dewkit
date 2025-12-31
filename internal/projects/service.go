package projects

import (
	"dewkit/config"
	"dewkit/internal/utils"
	"fmt"
	"strings"

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

func (s *Service) CreateProject(ownerId string, data ProjectCreateRequest) (ProjectListResponse, error) {

	tx, err := s.DB.Beginx()
	if err != nil {
		return ProjectListResponse{}, err
	}

	defer func() {
		if err != nil {
			_ = tx.Rollback()
		}
	}()

	var projectID string

	insertProjectQuery := `
		INSERT INTO projects (id, name, description, code, created_by_id)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id;
	`

	err = tx.QueryRowx(
		insertProjectQuery,
		utils.NewID("prj"),
		data.Name,
		data.Description,
		generateProjectCode(),
		ownerId,
	).Scan(&projectID)

	if err != nil {
		return ProjectListResponse{}, err
	}

	insertMemberQuery := `
		INSERT INTO project_members (id, project_id, user_id, role)
		VALUES ($1, $2, $3, 'admin');
	`

	_, err = tx.Exec(insertMemberQuery, utils.NewID("prjmem"), projectID, ownerId)
	if err != nil {
		return ProjectListResponse{}, err
	}

	if err = tx.Commit(); err != nil {
		return ProjectListResponse{}, err
	}

	// 2️⃣ Read: single SELECT for response
	return s.GetProjectResponse(ownerId, projectID)
}

func (s *Service) GetProjectResponse(userId string, projectId string) (ProjectListResponse, error) {

	query := `
		SELECT
			p.id,
			p.name,
			p.description,
			p.code,
			pm.role AS member_role,
			p.created_at
		FROM project_members pm
		JOIN projects p ON p.id = pm.project_id
		WHERE pm.user_id = $1 AND p.id = $2;
	`

	var project ProjectListResponse
	err := s.DB.Get(&project, query, userId, projectId)
	return project, err
}

func (s *Service) ListUserProjects(userId int) ([]ProjectListResponse, error) {
	query := `
		SELECT
			p.id,
			p.name,
			p.description,
			p.code,
			pm.role as member_role,
			p.created_at
		FROM project_members pm
		JOIN projects p
		ON p.id = pm.project_id
		WHERE pm.user_id = $1
		ORDER BY p.created_at DESC;
	`
	projects := []ProjectListResponse{}
	if err := s.DB.Select(
		&projects,
		query,
		userId,
	); err != nil {
		return nil, err
	}

	return projects, nil
}

func generateProjectCode() string {
	return strings.ToUpper(
		strings.ReplaceAll(uuid.New().String(), "-", ""),
	)
}

func (s *Service) ListMembers(projectId string, currentUserID *string) ([]ProjectMemberResponse, error) {
	baseQuery := `
		SELECT 
			pm.id, 
			u.email, 
			u.full_name, 
			pm.role 
		FROM project_members pm
		JOIN users u ON pm.user_id = u.id
		WHERE pm.project_id = $1
	`

	args := []any{projectId}
	argPos := 2

	if currentUserID != nil {
		baseQuery += fmt.Sprintf(" AND pm.user_id != $%d", argPos)
		args = append(args, *currentUserID)
	}

	members := []ProjectMemberResponse{}
	if err := s.DB.Select(&members, baseQuery, args...); err != nil {
		return nil, err
	}

	return members, nil

}
