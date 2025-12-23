package middlewares

import (
	"net/http"

	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

func LoggedInMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		sess, err := session.Get("session", c)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{
				"error": "session error",
			})
		}

		// Check authentication
		auth, ok := sess.Values["authenticated"].(bool)
		if !ok || !auth {
			return c.JSON(http.StatusUnauthorized, map[string]string{
				"error": "unauthorized",
			})
		}

		// Extract user_id
		userID, ok := sess.Values["user_id"]
		if !ok {
			return c.JSON(http.StatusUnauthorized, map[string]string{
				"error": "invalid session",
			})
		}

		// Attach to Echo context
		c.Set("user_id", userID)

		return next(c)
	}
}
