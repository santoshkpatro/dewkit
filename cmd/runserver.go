package cmd

import (
	"dewkit/internal/handlers"
	"fmt"

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

	e := echo.New()
	api := e.Group("/api")
	// ws := e.Group("/ws")

	handlers.RegisterAPIRoutes(api)

	e.Logger.Fatal(e.Start(":8000"))
}