package models

import (
	"time"

	"github.com/backwardgo/kanban/ids"
)

type Board struct {
	Id ids.BoardId `db:"id" json:"id"`

	TeamId ids.TeamId `db:"team_id" json:"teamId"`

	Name string `db:"name" json:"name"`

	CreatedAt time.Time  `db:"created_at" json:"createdAt"`
	DeletedAt *time.Time `db:"deleted_at" json:"deletedAt,omitempty"`
	UpdatedAt time.Time  `db:"updated_at" json:"updatedAt"`

	CreatedBy ids.UserId `db:"created_by" json:"createdBy"`
}

func (m *Board) Errors() Errors {
	m.Normalize()

	e := NewErrors()

	if m.Id.Present() && m.Id.Invalid() {
		e["id"] = "is invalid"
	}

	if m.Name == "" {
		e["name"] = "is required"
	}

	switch {
	case m.CreatedBy.Blank():
		e["createdBy"] = "is required"
	case m.CreatedBy.Invalid():
		e["createdBy"] = "is invalid"
	}

	return e
}

func (m *Board) Normalize() {
	m.Name = trimSpace(m.Name)
}
