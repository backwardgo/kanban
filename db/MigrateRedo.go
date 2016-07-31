package db

import "log"

func MigrateRedo(databaseURL string) {
	migrator, err := newMigrator(databaseURL)
	if err != nil {
		log.Fatal(err)
	}

	if err = migrator.Rollback(); err != nil {
		log.Fatal(err)
	}

	if err = migrator.Migrate(); err != nil {
		log.Fatal(err)
	}
}
