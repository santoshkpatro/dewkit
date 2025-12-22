package migrations

import "github.com/jmoiron/sqlx"

var V3 = Migration{
	Version: 3,
	Up: func(tx *sqlx.Tx) error {
		_, err := tx.Exec(
			`
			CREATE TYPE conversation_status AS ENUM (
				'open',
				'pending',
				'resolved',
				'archived'
			);

			CREATE TABLE conversations (
				id SERIAL PRIMARY KEY,
				customer_id BIGINT REFERENCES users(id) ON DELETE SET NULL,
				customer_full_name TEXT,
				customer_email TEXT,
				status TEXT NOT NULL,
				resolved_at TIMESTAMPTZ,
				archived_at TIMESTAMPTZ,
				assigned_to BIGINT REFERENCES users(id) ON DELETE SET NULL,

				created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
				updated_at TIMESTAMPTZ NOT NULL DEFAULT now()
			);
			`,
		)
		return err
	},
	Down: func(tx *sqlx.Tx) error {
		_, err := tx.Exec(
			`
			DROP TABLE IF EXISTS conversations;
			DROP TYPE IF EXISTS conversation_status;
			`,
		)
		return err
	},
}
