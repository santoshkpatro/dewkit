package migrations

import "github.com/jmoiron/sqlx"

var V7 = Migration{
	Version: 7,
	Up: func(tx *sqlx.Tx) error {
		_, err := tx.Exec(
			`
			ALTER TABLE users
				ADD COLUMN is_customer BOOLEAN NOT NULL DEFAULT FALSE,
				ADD COLUMN customer_identifier TEXT,
				ALTER COLUMN password_hash DROP NOT NULL;

			ALTER TABLE users
				ADD CONSTRAINT users_customer_identifier_unique UNIQUE (customer_identifier);

			CREATE INDEX idx_users_is_customer ON users (is_customer);
			`,
		)
		return err
	},
	Down: func(tx *sqlx.Tx) error {
		_, err := tx.Exec(
			`
			DROP INDEX IF EXISTS idx_users_is_customer;

			ALTER TABLE users
				DROP CONSTRAINT IF EXISTS users_customer_identifier_unique,
				DROP COLUMN IF EXISTS is_customer,
				DROP COLUMN IF EXISTS customer_identifier;
				ALTER COLUMN password_hash SET NOT NULL;
			`,
		)
		return err
	},
}
