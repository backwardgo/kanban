package db

import "github.com/backwardgo/kanban/models"

func CreateUser(
	txn Transaction,
	user *models.User,
) error {
	if errors := user.Errors(); errors.Present() {
		return errors
	}

	err := txn.
		insertInto("users").
		Blacklist("id", "created_at", "deleted_at", "updated_at").
		Record(user).
		Returning("*").
		QueryStruct(user)

	return translatePQError(err)
}
