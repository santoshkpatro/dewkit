package migrations

import "github.com/jmoiron/sqlx"

var V2 = Migration{
	Version: 2,
	Up: func(tx *sqlx.Tx) error {
		_, err := tx.Exec(
			`
			CREATE TYPE user_role AS ENUM ('admin', 'staff', 'superuser');

			ALTER TABLE users
				ADD COLUMN is_password_expired BOOLEAN NOT NULL DEFAULT FALSE,
				ADD COLUMN last_login_at TIMESTAMPTZ,
				ADD COLUMN role TEXT NOT NULL DEFAULT 'staff',
				ADD COLUMN created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
				ADD COLUMN updated_at TIMESTAMPTZ NOT NULL DEFAULT now();
			`,
		)
		return err
	},
	Down: func(tx *sqlx.Tx) error {
		_, err := tx.Exec(
			`
			ALTER TABLE users
				DROP COLUMN IF EXISTS is_password_expired,
				DROP COLUMN IF EXISTS last_login_at,
				DROP COLUMN IF EXISTS role,
				DROP COLUMN IF EXISTS created_at,
				DROP COLUMN IF EXISTS updated_at;

			DROP TYPE IF EXISTS user_role;
			`,
		)
		return err
	},
}
