package db

import "github.com/backwardgo/kanban/models"

func CardCreate(
	txn Transaction,
	card *models.Card,
) error {
	if errors := card.Errors(); errors.Present() {
		return errors
	}

	err := txn.
		insertInto("cards").
		Blacklist("id", "created_at", "deleted_at", "updated_at").
		Record(card).
		Returning("*").
		QueryStruct(card)

	return translatePQError(err)
}
