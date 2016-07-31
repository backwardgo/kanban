package db

import "github.com/backwardgo/kanban/ids"

type ListFilter struct {
	ListIdIn    []ids.ListId
	TeamIdIn    []ids.TeamId
	BoardIdIn   []ids.BoardId
	CreatedByIn []ids.UserId
}
