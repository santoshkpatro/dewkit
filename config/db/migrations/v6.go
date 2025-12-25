package migrations

import "github.com/jmoiron/sqlx"

var V6 = Migration{
	Version: 6,
	Up: func(tx *sqlx.Tx) error {
		_, err := tx.Exec(
			`
			ALTER TABLE projects
			ADD CONSTRAINT projects_name_unique UNIQUE (name);
			`,
		)
		return err
	},
	Down: func(tx *sqlx.Tx) error {
		_, err := tx.Exec(
			`
			ALTER TABLE projects
			DROP CONSTRAINT projects_name_unique;
			`,
		)
		return err
	},
}
