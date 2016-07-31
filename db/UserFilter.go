package db

import (
	"github.com/backwardgo/kanban/ids"
	dat "gopkg.in/mgutz/dat.v1"
)

type UserFilter struct {
	IdIn []ids.UserId
}

func (f *UserFilter) refineQuery(query *dat.SelectBuilder) *dat.SelectBuilder {
	if len(f.IdIn) > 0 {
		query = query.Where("users.id in $1", f.IdIn)
	}

	return query
}
