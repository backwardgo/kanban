package db

import "github.com/backwardgo/kanban/ids"

type CardFilter struct {
	CardIdIn    []ids.CardId
	ListIdIn    []ids.ListId
	CreatedByIn []ids.UserId
}
