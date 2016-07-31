package models

import (
	"time"

	"github.com/backwardgo/kanban/ids"
)

type Signup struct {
	Id ids.SignupId `db:"id" json:"id"`

	FirstName string `db:"first_name" json:"firstName"`
	LastName  string `db:"last_name" json:"lastName"`
	Password  string `db:"-" json:"password"`

	CreatedAt time.Time  `db:"created_at" json:"createdAt"`
	DeletedAt *time.Time `db:"deleted_at" json:"deletedAt,omitempty"`
	UpdatedAt time.Time  `db:"updated_at" json:"updatedAt"`
}

func (m *Signup) Errors() Errors {
	m.Normalize()

	e := NewErrors()

	if m.Id.Present() && m.Id.Invalid() {
		e["id"] = "is invalid"
	}

	if m.FirstName == "" {
		e["firstName"] = "is required"
	}

	if m.FirstName == "" {
		e["lastName"] = "is required"
	}

	switch {
	case m.Password == "":
		e["password"] = "is required"
	case len(m.Password) < 6:
		e["password"] = "is too short"
	}

	return e
}

func (m *Signup) Normalize() {
	m.FirstName = trimSpace(m.FirstName)
	m.LastName = trimSpace(m.LastName)
}
