package db

import (
	"github.com/backwardgo/kanban/ids"
	"github.com/backwardgo/kanban/models"
)

func UserById(
	txn Transaction,
	userId ids.UserId,
) (
	user models.User,
	err error,
) {
	err = txn.
		selekt("*").
		From("users").
		Where("id = $1", userId).
		QueryStruct(&user)

	return
}
