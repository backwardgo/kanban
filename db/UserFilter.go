package db

import (
	"github.com/backwardgo/kanban/ids"
	dat "gopkg.in/mgutz/dat.v1"
)

type UserFilter struct {
	DeletedAtIsNull bool
	UserIdIn        []ids.UserId
}

func (f UserFilter) refineQuery(query *dat.SelectBuilder) *dat.SelectBuilder {
	if f.DeletedAtIsNull {
		query = query.Where("users.deleted_at IS NULL")
	}

	if len(f.UserIdIn) > 0 {
		query = query.Where("users.id in $1", f.UserIdIn)
	}

	return query
}
