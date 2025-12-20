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
				ADD COLUMN created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
				ADD COLUMN updated_at TIMESTAMPTZ NOT NULL DEFAULT now();
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
