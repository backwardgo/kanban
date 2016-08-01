package db

import "github.com/backwardgo/kanban/models"

func CreateMember(
	txn Transaction,
	member *models.Member,
) error {
	if errors := member.Errors(); errors.Present() {
		return errors
	}

	err := txn.
		insertInto("members").
		Blacklist("id", "created_at", "deleted_at", "updated_at").
		Record(member).
		Returning("*").
		QueryStruct(member)

	return translatePQError(err)
}
