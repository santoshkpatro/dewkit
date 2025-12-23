package projects

import (
	"dewkit/config"

	"github.com/jmoiron/sqlx"
)

type Service struct {
	DB *sqlx.DB
}

func NewService() *Service {
	db := config.DB
	return &Service{DB: db}
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
