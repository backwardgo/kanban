package db

import (
	"path"

	"github.com/DavidHuie/gomigrate"
	"github.com/backwardgo/kanban/env"
	"github.com/jmoiron/sqlx"
)

func newMigrator(databaseURL string) (*gomigrate.Migrator, error) {
	sqlxDB, err := sqlx.Connect(`postgres`, databaseURL)
	if err != nil {
		return nil, err
	}

	migrations := path.Join(env.KANBAN_HOME, "db/migrations")

	return gomigrate.NewMigrator(
		sqlxDB.DB,
		gomigrate.Postgres{},
		migrations,
	)
}
