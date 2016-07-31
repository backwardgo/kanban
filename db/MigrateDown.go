package db

import "log"

func MigrateDown(databaseURL string) {
	migrator, err := newMigrator(databaseURL)
	if err != nil {
		log.Fatal(err)
	}

	err = migrator.Rollback()
	if err != nil {
		log.Fatal(err)
	}
}
