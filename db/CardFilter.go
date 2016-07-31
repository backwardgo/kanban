package db

import "github.com/backwardgo/kanban/ids"

type CardFilter struct {
	CardIdIn    []ids.CardId
	TeamIdIn    []ids.TeamId
	ListIdIn    []ids.ListId
	CreatedByIn []ids.UserId
}
