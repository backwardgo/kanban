package db

import (
	"database/sql"

	"github.com/backwardgo/kanban/models"
)

func UserList(
	txn Transaction,
	userFilter UserFilter,
) (
	[]models.User,
	error,
) {
	users := []models.User{}

	query := txn.
		selekt("users.*").
		From("users")

	err := userFilter.
		refineQuery(query).
		QueryStructs(&users)

	if err == sql.ErrNoRows {
		err = nil
	}

	return users, err
}
