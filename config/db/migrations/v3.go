package migrations

import (
	"context"

	"github.com/jackc/pgx/v5"
)

var V3 = Migration{
	Version: 3,
	Up: func(tx pgx.Tx, ctx context.Context) error {
		_, err := tx.Exec(
			ctx,
			`
			CREATE TABLE conversations (
				id SERIAL PRIMARY KEY,
				customer_id INTEGER REFERENCES users(id) ON DELETE SET NULL,
				customer_full_name TEXT,
				customer_email TEXT,
				status TEXT NOT NULL,
				resolved_at TIMESTAMPTZ,
				archived_at TIMESTAMPTZ,
				assigned_to INTEGER REFERENCES users(id) ON DELETE SET NULL,

				created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
				updated_at TIMESTAMPTZ NOT NULL DEFAULT now()
			);
			`,
		)
		return err
	},
	Down: func(tx pgx.Tx, ctx context.Context) error {
		_, err := tx.Exec(
			ctx,
			`
			DROP TABLE IF EXISTS conversations;
			`,
		)
		return err
	},
}
