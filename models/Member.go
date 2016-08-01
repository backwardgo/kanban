package models

import (
	"time"

	"github.com/backwardgo/kanban/ids"
)

type Member struct {
	Id ids.UserId `db:"id" json:"id"`

	BoardId ids.BoardId `db:"board_id" json:"boardId"`
	UserId  ids.UserId  `db:"user_id" json:"userId"`
	Role    Role        `db:"role" json:"role"`

	CreatedAt time.Time  `db:"created_at" json:"createdAt"`
	DeletedAt *time.Time `db:"deleted_at" json:"deletedAt,omitempty"`
	UpdatedAt time.Time  `db:"updated_at" json:"updatedAt"`

	CreatedBy ids.UserId `db:"created_by" json:"createdBy"`
}

func (m *Member) Errors() Errors {
	m.Normalize()

	e := NewErrors()

	if m.Id.Present() && m.Id.Invalid() {
		e["id"] = "is invalid"
	}

	switch {
	case m.BoardId.Blank():
		e["boardId"] = "is required"
	case m.BoardId.Invalid():
		e["boardId"] = "is invalid"
	}

	switch {
	case m.Role.Blank():
		e["role"] = "is required"
	case m.Role.Invalid():
		e["role"] = "is invalid"
	}

	switch {
	case m.UserId.Blank():
		e["userId"] = "is required"
	case m.UserId.Invalid():
		e["userId"] = "is invalid"
	}

	switch {
	case m.CreatedBy.Blank():
		e["createdBy"] = "is required"
	case m.CreatedBy.Invalid():
		e["createdBy"] = "is invalid"
	}

	return e
}

func (m *Member) Normalize() {
}
