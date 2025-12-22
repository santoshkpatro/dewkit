package cmd

import (
	"context"
	"dewkit/config"
	"dewkit/config/db/migrations"
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/spf13/cobra"
)

var syncSchema bool

func getCurrentDBVersion(db *sqlx.DB) (int, error) {
	var version int

	err := db.Get(&version, `
		SELECT value::int
		FROM settings
		WHERE key = 'db.version'
	`)

	return version, err
}

var dbMigrateCmd = &cobra.Command{
	Use:   "db_migrate",
	Short: "DB Migration",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()
		db, err := config.GetDB(ctx)
		if err != nil {
			panic("Unable to connect to DB")
		}
		defer db.Close()

		currentVersion, err := getCurrentDBVersion(db)
		if err != nil {
			return fmt.Errorf("failed to read db.version: %w", err)
		}

		fmt.Printf("Current DB version: %d\n", currentVersion)

		for _, m := range migrations.All {
			if m.Version <= currentVersion {
				continue
			}

			fmt.Printf("Running migration %d...\n", m.Version)

			tx, err := db.Beginx()
			if err != nil {
				return err
			}

			// run migration
			if err := m.Up(tx); err != nil {
				_ = tx.Rollback()
				return fmt.Errorf("migration %d failed: %w", m.Version, err)
			}

			// update db.version
			_, err = tx.Exec(`
			UPDATE settings
			SET value = to_jsonb($1::int)
			WHERE key = 'db.version'
		`, m.Version)
			if err != nil {
				_ = tx.Rollback()
				return err
			}

			if err := tx.Commit(); err != nil {
				return err
			}

			fmt.Printf("Migration %d applied successfully\n", m.Version)
		}

		fmt.Println("All migrations complete")

		if syncSchema {
			fmt.Println("Syncing schema.sql...")

			if err := SchemaSync(); err != nil {
				fmt.Println("sync failed, please manually try to sync schema.sql")
			}
		}

		return nil
	},
}

func init() {
	dbMigrateCmd.Flags().BoolVar(
		&syncSchema,
		"sync",
		false,
		"Sync schema.sql after successful migrations",
	)

	rootCmd.AddCommand(dbMigrateCmd)
}
