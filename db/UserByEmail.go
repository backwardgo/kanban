package db

import "github.com/backwardgo/kanban/models"

func UserByEmail(
	txn Transaction,
	email models.Email,
) (
	user models.User,
	err error,
) {
	err = txn.
		selekt("*").
		From("users").
		Where("email = $1", email).
		QueryStruct(&user)

	return
}
