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

	query := txn.selekt(`users.*`).From(`users`)
	query = userFilter.refineQuery(query)

	if pager != nil {
		totalRecords, err := UserCount(txn, userFilter)
		if err != nil {
			return users, err
		}

		pager.setTotalRecords(totalRecords)
		query = pager.refineQuery(query)
	}

	if pager != nil && len(orderBy) == 0 {
		orderBy = append(orderBy, OrderBy{
			tableName: `users`,
			fieldName: `id`,
			direction: Ascending,
		})
	}

	for _, ob := range orderBy {
		clause := ob.Clause()
		if _, ok := userSliceOrderByWhitelist[clause]; ok {
			query = query.OrderBy(clause)
		}
	}

	err := query.QueryStructs(&users)
	if err == sql.ErrNoRows {
		err = nil
	}

	return users, err
}

var userSliceOrderByWhitelist = map[string]struct{}{
	`users.id ASC`:  struct{}{},
	`users.id DESC`: struct{}{},

	`users.first_name ASC`:  struct{}{},
	`users.first_name DESC`: struct{}{},

	`users.last_name ASC`:  struct{}{},
	`users.last_name DESC`: struct{}{},

	`users.initials ASC`:  struct{}{},
	`users.initials DESC`: struct{}{},

	`users.email ASC`:  struct{}{},
	`users.email DESC`: struct{}{},

	`users.created_at ASC`:  struct{}{},
	`users.created_at DESC`: struct{}{},

	`users.deleted_at ASC`:  struct{}{},
	`users.deleted_at DESC`: struct{}{},

	`users.updated_at ASC`:  struct{}{},
	`users.updated_at DESC`: struct{}{},
}
