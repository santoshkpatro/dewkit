package cmd

import (
	"context"
	"dewkit/config"
	"dewkit/config/db/migrations"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/spf13/cobra"
)

func getCurrentDBVersion(ctx context.Context, db *pgx.Conn) (int, error) {
	var version int

	err := db.QueryRow(ctx, `
		SELECT value::int
		FROM settings
		WHERE key = 'db.version'
	`).Scan(&version)

	return version, err
}

var dbMigrateCmd = &cobra.Command{
	Use:   "db_migrate",
	Short: "DB Migration",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		db, err := config.GetDB(ctx)
		if err != nil {
			return err
		}
		defer db.Close(ctx)

		currentVersion, err := getCurrentDBVersion(ctx, db)
		if err != nil {
			return fmt.Errorf("failed to read db.version: %w", err)
		}

		fmt.Printf("Current DB version: %d\n", currentVersion)

		for _, m := range migrations.All {
			if m.Version <= currentVersion {
				continue
			}

			fmt.Printf("Running migration %d...\n", m.Version)

			tx, err := db.Begin(ctx)
			if err != nil {
				return err
			}

			if err := m.Up(tx, ctx); err != nil {
				tx.Rollback(ctx)
				return fmt.Errorf("migration %d failed: %w", m.Version, err)
			}

			_, err = tx.Exec(ctx, `
				UPDATE settings
				SET value = to_jsonb($1::int)
				WHERE key = 'db.version'
			`, m.Version)
			if err != nil {
				tx.Rollback(ctx)
				return err
			}

			if err := tx.Commit(ctx); err != nil {
				return err
			}

			fmt.Printf("Migration %d applied successfully\n", m.Version)
		}

		fmt.Println("All migrations complete")
		return nil
	},
}

func init() {
	rootCmd.AddCommand(dbMigrateCmd)
}
