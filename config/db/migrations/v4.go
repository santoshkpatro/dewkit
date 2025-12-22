package migrations

import "github.com/jmoiron/sqlx"

var V4 = Migration{
	Version: 4,
	Up: func(tx *sqlx.Tx) error {
		_, err := tx.Exec(
			`
			ALTER TABLE users
			RENAME COLUMN salt TO password_salt;

			ALTER TABLE users
			ALTER COLUMN password_salt DROP NOT NULL;

			ALTER TABLE users
			ADD COLUMN IF NOT EXISTS password_changed_at TIMESTAMP NULL;

			ALTER TABLE users
			ADD COLUMN IF NOT EXISTS failed_login_attempts INT NOT NULL DEFAULT 0;
			`,
		)
		return err
	},
	Down: func(tx *sqlx.Tx) error {
		_, err := tx.Exec(
			`
			ALTER TABLE users
			DROP COLUMN IF EXISTS failed_login_attempts;

			ALTER TABLE users
			DROP COLUMN IF EXISTS password_changed_at;

			ALTER TABLE users
			RENAME COLUMN password_salt TO salt;

			ALTER TABLE users
			ALTER COLUMN salt SET NOT NULL;
			`,
		)
		return err
	},
}
