package migrations

import "github.com/jmoiron/sqlx"

var V9 = Migration{
	Version: 9,
	Up: func(tx *sqlx.Tx) error {
		_, err := tx.Exec(
			`
			ALTER TABLE projects
			DROP COLUMN IF EXISTS code;
			`,
		)
		return err
	},
	Down: func(tx *sqlx.Tx) error {
		_, err := tx.Exec(
			`
			ALTER TABLE projects
			ADD COLUMN code TEXT;
			`,
		)
		return err
	},
}
