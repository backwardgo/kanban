package db

import "github.com/backwardgo/kanban/ids"

type ListFilter struct {
	ListIdIn    []ids.ListId
	BoardIdIn   []ids.BoardId
	CreatedByIn []ids.UserId
}
