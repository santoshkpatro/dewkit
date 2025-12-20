package migrations

import (
	"context"

	"github.com/jackc/pgx/v5"
)

var V4 = Migration{
	Version: 4,
	Up: func(tx pgx.Tx, ctx context.Context) error {
		_, err := tx.Exec(
			ctx,
			`
			ALTER TABLE users
				created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
				updated_at TIMESTAMPTZ NOT NULL DEFAULT now();
			`,
		)
		return err
	},
	Down: func(tx pgx.Tx, ctx context.Context) error {
		_, err := tx.Exec(
			ctx,
			`
			ALTER TABLE users
				DROP COLUMN IF EXISTS created_at,
				DROP COLUMN IF EXISTS updated_at;
			`,
		)
		return err
	},
}
