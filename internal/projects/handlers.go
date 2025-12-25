package projects

import (
	"log/slog"
	"net/http"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
)

func ProjectCreateHandler(c echo.Context) error {
	userID := c.Get("user_id").(int)
	var req ProjectCreateRequest

	if err := c.Bind(&req); err != nil {
		slog.Debug(
			"failed to bind login request",
			"err", err,
			"path", c.Path(),
		)

		return c.JSON(http.StatusBadRequest, echo.Map{
			"error": "Invalid request body",
		})
	}

	service := NewService()
	project, err := service.CreateProject(userID, req)
	if err != nil {
		slog.Error("failed to create project", "err", err)
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"error": "Failed to create project",
		})
	}

	return c.JSON(http.StatusCreated, project)
}

func ProjectListHandler(c echo.Context) error {
	userID := c.Get("user_id").(int)
	db := c.Get("db").(*sqlx.DB)

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
	if err := db.Select(
		&projects,
		query,
		userID,
	); err != nil {
		slog.Error("failed to fetch projects", "err", err)
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"error": "Failed to fetch projects",
		})
	}

	return c.JSON(http.StatusOK, projects)
}
