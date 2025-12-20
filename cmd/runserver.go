package cmd

import (
	"context"
	"dewkit/config"
	"dewkit/internal/handlers"
	"fmt"
	"time"

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

func runserver() {
	fmt.Println("Staring dewkit server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	config.SetupDB(ctx)

	e := echo.New()
	api := e.Group("/api")
	// ws := e.Group("/ws")

	handlers.RegisterAPIRoutes(api)

	e.Logger.Fatal(e.Start(":8000"))
}