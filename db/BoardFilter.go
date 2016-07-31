package db

import "github.com/backwardgo/kanban/ids"

type BoardFilter struct {
	BoardIdIn   []ids.BoardId
	TeamIdIn    []ids.TeamId
	CreatedByIn []ids.UserId
}
