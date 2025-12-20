package cmd

import (
	"context"
	"dewkit/config"
	"dewkit/config/db/migrations"
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"time"

	"github.com/spf13/cobra"
)

var dbSchemaSync = &cobra.Command{
	Use:   "db_schema_sync",
	Short: "DB Migration",
	RunE: func(cmd *cobra.Command, args []string) error {
		return SchemaSync()
	},
}

func cleanSchemaRegex(schema string) string {
	replacements := []struct {
		re   *regexp.Regexp
		repl string
	}{
		// 0. Remove pg_catalog.set_config(search_path...)
		{regexp.MustCompile(
			`(?m)^SELECT\s+pg_catalog\.set_config\('search_path',\s*''\s*,\s*false\);\s*`,
		), ``},

		// 1. Remove all SET statements
		{regexp.MustCompile(`(?m)^SET\s+.*?;\s*`), ``},

		// 2. Remove SQL line comments (-- ...)
		{regexp.MustCompile(`(?m)--.*$`), ``},

		// 3. Remove SQL block comments (/* ... */)
		{regexp.MustCompile(`(?s)/\*.*?\*/`), ``},

		// 4. Remove COMMENT ON statements
		{regexp.MustCompile(`(?m)^COMMENT ON .*?;\s*`), ``},

		// 5. Remove public. schema prefix
		{regexp.MustCompile(`\bpublic\.`), ``},

		// 6. Collapse excessive blank lines
		{regexp.MustCompile(`\n{3,}`), "\n\n"},
	}

	clean := schema
	for _, r := range replacements {
		clean = r.re.ReplaceAllString(clean, r.repl)
	}

	return clean
}

func getLatestMigrationVersion() int {
	latest := 0
	for _, m := range migrations.All {
		if m.Version > latest {
			latest = m.Version
		}
	}
	return latest
}

func SchemaSync() error {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Minute)
	defer cancel()

	dsn := config.GetEnv("DB_URL")

	cmd := exec.CommandContext(
		ctx,
		"pg_dump",
		"--schema-only",
		"--no-owner",
		"--no-privileges",
		dsn,
	)

	out, err := cmd.Output()
	if err != nil {
		return fmt.Errorf("pg_dump failed: %w", err)
	}

	cleaned := cleanSchemaRegex(string(out))

	latestVersion := getLatestMigrationVersion()

	// Append db.version update so fresh installs are up-to-date
	cleaned += fmt.Sprintf(`
-- Ensure db.version is set to latest migration
INSERT INTO settings (key, value)
VALUES ('db.version', to_jsonb(%d::int))
ON CONFLICT (key)
DO UPDATE SET value = EXCLUDED.value;
	`, latestVersion)

	return os.WriteFile("schema.sql", []byte(cleaned), 0644)
}

func init() {
	rootCmd.AddCommand(dbSchemaSync)
}
