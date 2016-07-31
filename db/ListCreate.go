package db

import "github.com/backwardgo/kanban/models"

func ListCreate(
	txn Transaction,
	list *models.List,
) error {
	if errors := list.Errors(); errors.Present() {
		return errors
	}

	err := txn.
		insertInto("lists").
		Blacklist("id", "created_at", "deleted_at", "updated_at").
		Record(list).
		Returning("*").
		QueryStruct(list)

	return translatePQError(err)
}
