package middlewares

import (
	"database/sql"
	"dewkit/config"
	"net/http"

	"github.com/labstack/echo/v4"
)

func ProjectPermissionMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		db := config.DB

		// 1️⃣ Extract user ID from context
		userID, ok := c.Get("user_id").(string)
		if !ok {
			return echo.NewHTTPError(
				http.StatusUnauthorized,
				"Authentication required",
			)
		}

		// 2️⃣ Extract project ID from URL
		projectID := c.Param("projectId")

		// 3️⃣ Query membership + role
		var role string
		query := `
				SELECT role
				FROM project_members
				WHERE project_id = $1 AND user_id = $2;
			`

		err := db.Get(&role, query, projectID, userID)
		if err == sql.ErrNoRows {
			return echo.NewHTTPError(
				http.StatusForbidden,
				"You do not have sufficient permissions to access this project",
			)
		}
		if err != nil {
			return err // internal server error
		}

		// 4️⃣ Attach useful context
		c.Set("project_id", projectID)
		c.Set("project_role", role)

		// 5️⃣ Continue
		return next(c)
	}
}
