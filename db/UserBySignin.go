package db

import (
	"database/sql"

	"github.com/backwardgo/kanban/models"
	"github.com/backwardgo/kanban/passwords"
)

func UserBySignin(
	txn Transaction,
	signin models.Signin,
) (
	models.User,
	error,
) {
	var nobody models.User

	if errors := signin.Errors(); errors.Present() {
		return nobody, errors
	}

	user, err := UserByEmail(txn, signin.Email)
	if err != nil {
		return nobody, err
	}

	match := passwords.Equal(
		signin.PasswordDigest(),
		user.PasswordDigest,
	)

	if !match {
		return nobody, sql.ErrNoRows
	}

	return user, nil
}
