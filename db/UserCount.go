package db

import "database/sql"

func UserCount(
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

	return count, err
}
