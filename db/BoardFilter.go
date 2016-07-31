package db

import "github.com/backwardgo/kanban/ids"

type BoardFilter struct {
	IdIn []ids.BoardId

	TeamIdIn []ids.TeamId
}
