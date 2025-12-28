package cmd

import (
	"context"
	"dewkit/config"
	"dewkit/config/middlewares"
	"dewkit/internal/auth"
	"dewkit/internal/projects"
	"dewkit/internal/transport"
	"fmt"
	"net/http"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/spf13/cobra"
)

var runserverCmd = &cobra.Command{
	Use:   "runserver",
	Short: "Run dewkit server",
	Run: func(cmd *cobra.Command, args []string) {
		runserver()
	},
}

func init() {
	rootCmd.AddCommand(runserverCmd)
}

type AppValidator struct {
	Validator *validator.Validate
}

func (cv *AppValidator) Validate(i interface{}) error {
	return cv.Validator.Struct(i)
}

func runserver() {
	fmt.Println("Staring dewkit server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	db, err := config.SetupDB(ctx)
	if err != nil {
		panic("Failed to setup DB")
	}
	cache, err := config.SetupCache(ctx)
	if err != nil {
		panic("Failed to setup Cache")
	}

	isProd := config.GetEnvDefault("ENV", "production") == "production"
	store := sessions.NewCookieStore([]byte(config.GetEnv("SECRET_KEY")))
	store.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   86400,
		HttpOnly: true,
		Secure:   isProd, // true in prod (HTTPS)
		SameSite: http.SameSiteLaxMode,
	}

	e := echo.New()
	e.Validator = &AppValidator{
		Validator: validator.New(),
	}
	e.Use(middlewares.DBMiddleware(db))
	e.Use(middlewares.CacheMiddleware(cache))
	e.Use(session.Middleware(store))

	api := e.Group("/api")
	ws := e.Group("/ws")

	auth.RegisterAPIRoutes(api.Group("/auth"))

	projects.RegisterAPIRoutes(api.Group("/projects"))
	projects.RegisterWSRoutes(ws.Group("/projects"))

	transport.RegisterWSRoutes(ws.Group("/transport"))
	transport.RegisterAPIRoutes(api.Group("/transport"))

	e.Logger.Fatal(e.Start(":8000"))
}
