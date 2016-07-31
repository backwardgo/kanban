package db

import "github.com/backwardgo/kanban/ids"

type TeamFilter struct {
	TeamIdIn    []ids.TeamId
	CreatedByIn []ids.UserId
}
