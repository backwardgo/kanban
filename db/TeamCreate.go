package db

import "github.com/backwardgo/kanban/models"

func TeamCreate(
	txn Transaction,
	team *models.Team,
) error {
	if errors := team.Errors(); errors.Present() {
		return errors
	}

	err := txn.
		insertInto("teams").
		Blacklist("id", "created_at", "deleted_at", "updated_at").
		Record(team).
		Returning("*").
		QueryStruct(team)

	return translatePQError(err)
}
