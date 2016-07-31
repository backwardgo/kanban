package db

import "github.com/backwardgo/kanban/ids"

type CardFilter struct {
	IdIn []ids.CardId

	ListIdIn []ids.ListId
}
