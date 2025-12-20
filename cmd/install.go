package cmd

import (
	"context"
	"dewkit/config"
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/spf13/cobra"
)

var installCmd = &cobra.Command{
	Use:   "install",
	Short: "Install dewkit",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("Installing dewkit...")

		ctx, cancel := context.WithTimeout(context.Background(), 2*time.Minute)
		defer cancel()

		db, err := config.GetDB(ctx)
		if err != nil {
			return err
		}
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
	"system.maintenance": false,
}

func install(ctx context.Context, db *pgx.Conn) error {
	// 1. Check if settings table already exists
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
		fmt.Println("Skipping installation! Database already initialized")
		return nil
	}

	// 2. Load schema.sql
	schema, err := os.ReadFile("schema.sql")
	if err != nil {
		return fmt.Errorf("failed to read schema.sql: %w", err)
	}

	// 3. Execute schema.sql
	fmt.Println("Applying schema.sql...")
	if _, err := db.Exec(ctx, string(schema)); err != nil {
		return fmt.Errorf("failed to apply schema.sql: %w", err)
	}

	// 4. Seed initial settings (excluding db.version)
	tx, err := db.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	for key, value := range initialSettings {
		if key == "db.version" {
			continue
		}

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

	fmt.Println("Dewkit installed successfully 🎉")
	return nil
}
