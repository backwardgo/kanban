package db

import (
	"database/sql"

	"github.com/backwardgo/kanban/models"
)

func UserSlice(
	txn Transaction,
	userFilter UserFilter,
	pager Pager,
	orderBy ...OrderBy,
) (
	[]models.User,
	error,
) {
	users := []models.User{}

	query := txn.selekt("users.*").From("users")
	query = userFilter.refineQuery(query)

	if pager != nil {
		totalRecords, err := UserCount(txn, userFilter)
		if err != nil {
			return users, err
		}

		pager.setTotalRecords(totalRecords)
		query = pager.refineQuery(query)

		if len(orderBy) == 0 {
			query = query.OrderBy(`users.id asc`)
		}
	}

	for i := range orderBy {
		clause := orderBy[i].String()
		query = query.OrderBy(clause)
	}

	err := query.QueryStructs(&users)
	if err == sql.ErrNoRows {
		err = nil
	}

	return users, err
}
