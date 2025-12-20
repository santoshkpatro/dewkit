package migrations

import (
	"context"

	"github.com/jackc/pgx/v5"
)

var V1 = Migration{
	Version: 1,
	Up: func(tx pgx.Tx, ctx context.Context) error {
		_, err := tx.Exec(
			ctx,
			`
			CREATE TABLE users (
				id SERIAL PRIMARY KEY,
				email TEXT NOT NULL UNIQUE,
				full_name TEXT NOT NULL,
				password_hash TEXT NOT NULL,
				salt TEXT NOT NULL,
				is_active BOOLEAN NOT NULL DEFAULT TRUE
			);
			`,
		)
		return err
	},
	Down: func(tx pgx.Tx, ctx context.Context) error {
		_, err := tx.Exec(
			ctx,
			`DROP TABLE IF EXISTS users;`,
		)
		return err
	},
}
