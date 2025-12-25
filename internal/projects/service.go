package projects

import (
	"dewkit/config"
	"strings"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type Service struct {
	DB *sqlx.DB
}

func NewService() *Service {
	db := config.DB
	return &Service{DB: db}
}

func (s *Service) CreateProject(ownerId int, data ProjectCreateRequest) (ProjectListResponse, error) {

	tx, err := s.DB.Beginx()
	if err != nil {
		return ProjectListResponse{}, err
	}

	defer func() {
		if err != nil {
			_ = tx.Rollback()
		}
	}()

	var projectID int

	insertProjectQuery := `
		INSERT INTO projects (name, description, code, created_by_id)
		VALUES ($1, $2, $3, $4)
		RETURNING id;
	`

	err = tx.QueryRowx(
		insertProjectQuery,
		data.Name,
		data.Description,
		generateProjectCode(),
		ownerId,
	).Scan(&projectID)

	if err != nil {
		return ProjectListResponse{}, err
	}

	insertMemberQuery := `
		INSERT INTO project_members (project_id, user_id, role)
		VALUES ($1, $2, 'admin');
	`

	_, err = tx.Exec(insertMemberQuery, projectID, ownerId)
	if err != nil {
		return ProjectListResponse{}, err
	}

	if err = tx.Commit(); err != nil {
		return ProjectListResponse{}, err
	}

	// 2️⃣ Read: single SELECT for response
	return s.GetProjectResponse(ownerId, projectID)
}

func (s *Service) GetProjectResponse(userId int, projectId int) (ProjectListResponse, error) {

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
