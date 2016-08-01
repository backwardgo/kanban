package models

import (
	"time"

	"github.com/backwardgo/kanban/ids"
)

type User struct {
	Id ids.UserId `db:"id" json:"id"`

	FirstName string `db:"first_name" json:"firstName"`
	LastName  string `db:"last_name" json:"lastName"`
	Initials  string `db:"initials" json:"initials"`
	Biography string `db:"biography" json:"biography"`

	Email          Email  `db:"email" json:"email"`
	PasswordDigest []byte `db:"password_digest" json:"-"`

	CreatedAt time.Time `db:"created_at" json:"createdAt"`
	UpdatedAt time.Time `db:"updated_at" json:"updatedAt"`
}

func (m *User) Errors() Errors {
	m.Normalize()

	e := NewErrors()

	if m.Id.Present() && m.Id.Invalid() {
		e["id"] = "is invalid"
	}

	if m.FirstName == "" {
		e["firstName"] = "is required"
	}

	if m.LastName == "" {
		e["lastName"] = "is required"
	}

	switch {
	case m.Email.Blank():
		e["email"] = "is required"
	case m.Email.Invalid():
		e["email"] = "is invalid"
	}

	return e
}

func (m *User) Normalize() {
	m.FirstName = trimSpace(m.FirstName)
	m.LastName = trimSpace(m.LastName)
	m.Initials = trimSpace(m.Initials)
	m.Biography = trimSpace(m.Biography)
	m.Email = m.Email.Normalize()
}
