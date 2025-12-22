package migrations

import "github.com/jmoiron/sqlx"

var V1 = Migration{
	Version: 1,

	Up: func(tx *sqlx.Tx) error {
		_, err := tx.Exec(`
			CREATE TABLE users (
				id BIGSERIAL PRIMARY KEY,
				email TEXT NOT NULL UNIQUE,
				full_name TEXT NOT NULL,
				password_hash TEXT NOT NULL,
				salt TEXT NOT NULL,
				is_active BOOLEAN NOT NULL DEFAULT TRUE
			);
		`)
		return err
	},

	Down: func(tx *sqlx.Tx) error {
		_, err := tx.Exec(`DROP TABLE IF EXISTS users;`)
		return err
	},
}
