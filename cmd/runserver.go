package cmd

import (
	"context"
	"dewkit/config"
	"dewkit/internal/middlewares"
	"dewkit/internal/services/auth"
	"fmt"
	"time"

	"github.com/go-playground/validator/v10"
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

	e := echo.New()
	e.Validator = &AppValidator{
		Validator: validator.New(),
	}
	e.Use(middlewares.DBMiddleware(db))

	api := e.Group("/api")
	// ws := e.Group("/ws")

	auth.RegisterRoutes(api.Group("/auth"))

	e.Logger.Fatal(e.Start(":8000"))
}
