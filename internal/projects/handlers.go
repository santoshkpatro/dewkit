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

	db := c.Get("db").(*sqlx.DB)
	var exists bool
	err := db.Get(
		&exists,
		`SELECT EXISTS (SELECT 1 FROM projects WHERE name = $1)`,
		req.Name,
	)
	if err != nil {
		slog.Error("Failed to check for project name uniqueness", "err", err)
		return c.JSON(http.StatusBadGateway, echo.Map{
			"error": "Failed to creae project!",
		})
	}
	if exists {
		slog.Warn("A project already exists with the given name")
		return c.JSON(http.StatusConflict, echo.Map{"error": "A project already exists with the given name"})
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

func ProjectMembersHandler(c echo.Context) error {
	// ctx := c.Request().Context()

	userID := c.Get("user_id").(int)
	projectID := c.Get("project_id").(int)
	// cache := c.Get("cache").(*redis.Client)

	service := NewService()

	// cacheKey := fmt.Sprintf("project:members:%d", projectID)

	// 1️⃣ Try cache first
	// if cached, err := cache.Get(ctx, cacheKey).Result(); err == nil {
	// 	var members []ProjectMemberResponse
	// 	if err := json.Unmarshal([]byte(cached), &members); err == nil {
	// 		return c.JSON(http.StatusOK, members)
	// 	}
	// }

	// 2️⃣ Cache miss → fetch from service / DB
	members, err := service.ListMembers(projectID, &userID)
	if err != nil {
		slog.Error("failed to list members", "err", err)
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"error": "Failed to list members",
		})
	}

	// 3️⃣ Store in cache (TTL = 10 seconds)
	// if data, err := json.Marshal(members); err == nil {
	// 	_ = cache.Set(ctx, cacheKey, data, 10*time.Second).Err()
	// }

	return c.JSON(http.StatusOK, members)
}
