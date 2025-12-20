package cmd

import (
	"context"
	"dewkit/config"
	"encoding/json"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/spf13/cobra"
)

var installCmd = &cobra.Command{
	Use:   "install",
	Short: "Install dewkit",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("Installing dewkit...")
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		db, _ := config.GetDB(ctx)
		defer db.Close(ctx)

		return install(ctx, db)
	},
}

func init() {
	rootCmd.AddCommand(installCmd)
}

var initialSettings = map[string]any{
	"app.baseUrl":        "https://dewkit.app",
	"app.supportEmail":   "support@dewkit.app",
	"db.version":         0,
	"system.maintenance": false,
}

func install(ctx context.Context, db *pgx.Conn) error {
	// 1. Checking if settings table exists
	var exists bool
	err := db.QueryRow(ctx, `
		SELECT EXISTS (
			SELECT FROM information_schema.tables
			WHERE table_name = 'settings'
		)
	`).Scan(&exists)
	if err != nil {
		return err
	}

	if exists {
		fmt.Println("Skipping installation! Settings has been initiated already")
		return nil
	}

	// 2. Creating a table called settings
	_, err = db.Exec(ctx, `
		CREATE TABLE settings (
			key TEXT PRIMARY KEY,
			value JSONB NOT NULL
		);
	`)
	if err != nil {
		return err
	}
	fmt.Println("Creating settings table")

	// 3. Seeding with some initial set of data
	tx, err := db.Begin(ctx)
	if err != nil {
		return err
	}

	defer tx.Rollback(ctx)

	for key, value := range initialSettings {
		data, err := json.Marshal(value)
		if err != nil {
			return err
		}
		_, err = tx.Exec(ctx, `
			INSERT INTO settings (key, value)
			VALUES ($1, $2)
			ON CONFLICT (key) DO NOTHING
		`, key, data)
		if err != nil {
			return err
		}

		fmt.Printf("Inserted setting: %s\n", key)
	}

	if err := tx.Commit(ctx); err != nil {
		return err
	}

	fmt.Println("Dewkit installed successfully ...")
	return nil
}
