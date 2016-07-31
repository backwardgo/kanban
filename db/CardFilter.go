package db

import "github.com/backwardgo/kanban/ids"

type CardFilter struct {
	IdIn     []ids.CardId
	TeamIdIn []ids.TeamId
	ListIdIn []ids.ListId
}
