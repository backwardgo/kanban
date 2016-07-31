package db

import (
	"github.com/backwardgo/kanban/models"
	"github.com/lib/pq"
)

func translatePQError(err error) error {
	if err == nil {
		return nil
	}

	if pqe, ok := err.(*pq.Error); ok {

		switch pqe.Code.Name() {
		case "unique_violation":
			switch pqe.Constraint {

			case "index_users_on_email":
				return models.Errors{"email": "is conflicted"}

			}
		}
	}

	return err
}
