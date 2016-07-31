package db

import "log"

func MigrateUp(databaseURL string) {
	migrator, err := newMigrator(databaseURL)
	if err != nil {
		log.Fatal(err)
	}

	if err = migrator.Migrate(); err != nil {
		log.Fatal(err)
	}
}
