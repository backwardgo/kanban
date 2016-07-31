package models

import (
	"time"

	"github.com/backwardgo/kanban/ids"
)

type Card struct {
	Id     ids.CardId `db:"id" json:"id"`
	TeamId ids.TeamId `db:"team_id" json:"teamId"`

	ListId ids.ListId `db:"list_id" json:"listId"`
	Title  string     `db:"title" json:"title"`

	CreatedAt time.Time  `db:"created_at" json:"createdAt"`
	DeletedAt *time.Time `db:"deleted_at" json:"deletedAt,omitempty"`
	UpdatedAt time.Time  `db:"updated_at" json:"updatedAt"`

	CreatedBy ids.UserId `db:"created_by" json:"createdBy"`
}

func (m *Card) Errors() Errors {
	m.Normalize()

	e := NewErrors()

	if m.Id.Present() && m.Id.Invalid() {
		e["id"] = "is invalid"
	}

	switch {
	case m.TeamId.Blank():
		e["teamId"] = "is required"
	case m.TeamId.Invalid():
		e["teamId"] = "is invalid"
	}

	switch {
	case m.ListId.Blank():
		e["listId"] = "is required"
	case m.ListId.Invalid():
		e["listId"] = "is invalid"
	}

	if m.Title == "" {
		e["title"] = "is required"
	}

	switch {
	case m.CreatedBy.Blank():
		e["createdBy"] = "is required"
	case m.CreatedBy.Invalid():
		e["createdBy"] = "is invalid"
	}

	return e
}

func (m *Card) Normalize() {
	m.Title = trimSpace(m.Title)
}
