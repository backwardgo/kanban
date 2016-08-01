package db

import (
	"github.com/backwardgo/kanban/ids"
	dat "gopkg.in/mgutz/dat.v1"
)

type MemberFilter struct {
	DeletedAtIsNull bool

	BoardIdIn  []ids.BoardId
	MemberIdIn []ids.MemberId
	UserIdIn   []ids.UserId
}

func (f MemberFilter) refineQuery(query *dat.SelectBuilder) *dat.SelectBuilder {
	if f.DeletedAtIsNull {
		query = query.Where("members.deleted_at IS NULL")
	}

	if len(f.BoardIdIn) > 0 {
		query = query.Where("members.board_id in $1", f.BoardIdIn)
	}

	if len(f.MemberIdIn) > 0 {
		query = query.Where("members.id in $1", f.MemberIdIn)
	}

	if len(f.UserIdIn) > 0 {
		query = query.Where("members.user_id in $1", f.UserIdIn)
	}

	return query
}
