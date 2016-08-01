package db

import "github.com/backwardgo/kanban/ids"

type BoardFilter struct {
	BoardIdIn   []ids.BoardId
	CreatedByIn []ids.UserId
}
