package migrations

import (
	"context"

	"github.com/jackc/pgx/v5"
)

type Migration struct {
	Version int
	Up      func(tx pgx.Tx, ctx context.Context) error
	Down    func(tx pgx.Tx, ctx context.Context) error
}

var All = []Migration{
	V1,
	V2,
}
