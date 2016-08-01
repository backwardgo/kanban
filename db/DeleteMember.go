package db

import (
	"github.com/backwardgo/kanban/models"
	dat "gopkg.in/mgutz/dat.v1"
)

func DeleteMember(
	txn Transaction,
	member *models.Member,
) error {
	// TODO if member == nil
	// TODO if member.Id.Blank()

	err := txn.
		update("members").
		Set("deleted_at", dat.NOW).
		Where("id = $1", member.Id).
		Where("deleted_at IS NULL").
		Returning("*").
		QueryStruct(member)

	return translatePQError(err)
}
