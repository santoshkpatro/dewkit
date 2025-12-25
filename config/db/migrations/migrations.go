package migrations

import (
	"github.com/jmoiron/sqlx"
)

type Migration struct {
	Version int
	Up      func(tx *sqlx.Tx) error
	Down    func(tx *sqlx.Tx) error
}

var All = []Migration{
	V1,
	V2,
	V3,
	V4,
	V5,
	V6,
}
