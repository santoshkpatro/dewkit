package cmd

import (
	"context"
	"dewkit/config"
	"fmt"
	"os"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/spf13/cobra"
)

var installCmd = &cobra.Command{
	Use:   "install",
	Short: "Install dewkit",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("Installing dewkit...")
		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()

		db, err := config.GetDB(ctx)
		if err != nil {
			return fmt.Errorf("unable to connect to DB: %w", err)
		}
		defer db.Close()

		return install(db)
	},
}

func init() {
	rootCmd.AddCommand(installCmd)
}

func install(db *sqlx.DB) error {
	var exists bool
	err := db.Get(&exists, `
		SELECT EXISTS (
			SELECT FROM information_schema.tables
			WHERE table_name = 'settings'
		)
	`)
	if err != nil {
		return err
	}

	if exists {
		fmt.Println("Skipping installation! Database already initialized")
		return nil
	}

	schema, err := os.ReadFile("schema.sql")
	if err != nil {
		return fmt.Errorf("failed to read schema.sql: %w", err)
	}

	fmt.Println("Applying schema.sql...")
	if _, err := db.Exec(string(schema)); err != nil {
		return fmt.Errorf("failed to apply schema.sql: %w", err)
	}

	fmt.Println("Dewkit installed successfully 🎉")
	return nil
}
