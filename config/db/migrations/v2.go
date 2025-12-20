package migrations

import (
	"context"

	"github.com/jackc/pgx/v5"
)

var V2 = Migration{
	Version: 2,
	Up: func(tx pgx.Tx, ctx context.Context) error {
		_, err := tx.Exec(
			ctx,
			`
			ALTER TABLE users
				ADD COLUMN is_password_expired BOOLEAN NOT NULL DEFAULT FALSE,
				ADD COLUMN last_login_at TIMESTAMPTZ,
				ADD COLUMN role TEXT NOT NULL DEFAULT 'staff';
			`,
		)
		return err
	},
	Down: func(tx pgx.Tx, ctx context.Context) error {
		_, err := tx.Exec(
			ctx,
			`
			ALTER TABLE users
				DROP COLUMN IF EXISTS is_password_expired,
				DROP COLUMN IF EXISTS last_login_at,
				DROP COLUMN IF EXISTS role;
			`,
		)
		return err
	},
}
