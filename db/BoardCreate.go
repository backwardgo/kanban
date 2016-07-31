package db

import "github.com/backwardgo/kanban/models"

func BoardCreate(
	txn Transaction,
	board *models.Board,
) error {
	if errors := board.Errors(); errors.Present() {
		return errors
	}

	err := txn.
		insertInto("boards").
		Blacklist("id", "created_at", "deleted_at", "updated_at").
		Record(board).
		Returning("*").
		QueryStruct(board)

	return translatePQError(err)
}
