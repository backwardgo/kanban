package models

import (
	"database/sql/driver"

	"github.com/backwardgo/kanban/models/roles"
)

type Role string

func (r Role) Blank() bool {
	return r == ""
}

func (r Role) Invalid() bool {
	return !r.Valid()
}

func (r Role) Normalize() Role {
	return Role(toLower(trimSpace(string(r))))
}

func (r Role) Present() bool {
	return r != ""
}

func (r Role) Valid() bool {
	return r == roles.Admin ||
		r == roles.Default ||
		r == roles.Observer ||
		r == roles.Virtual
}

func (r Role) Value() (driver.Value, error) {
	return string(r), nil
}

func (r *Role) Scan(src interface{}) error {
	*r = Role(sprintf("%s", src))
	return nil
}
