package db

import "database/sql"

func CountUsers(
	txn Transaction,
	userFilter UserFilter,
) (
	count uint,
	err error,
) {
	query := txn.
		selekt("count(users.*)").
		From("users")

	err = userFilter.
		refineQuery(query).
		QueryScalar(&count)

	if err == sql.ErrNoRows {
		err = nil
	}

	return
}
