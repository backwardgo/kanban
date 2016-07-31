package db

import "github.com/backwardgo/kanban/ids"

type ListFilter struct {
	IdIn []ids.ListId

	BoardIdIn []ids.BoardId
}
