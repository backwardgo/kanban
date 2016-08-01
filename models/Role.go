package models

import "database/sql/driver"

type Role string

func (r Role) Blank() bool {
	return r == ""
}

func (r Role) Normalize() Role {
	return Role(toLower(trimSpace(string(r))))
}

func (r Role) Present() bool {
	return r != ""
}

func (r Role) Valid() bool {
	return r == "owner" ||
		r == "admin" ||
		r == "virtual" ||
		r == "member" ||
		r == "observer"
}

func (r Role) Value() (driver.Value, error) {
	return string(r), nil
}

func (r *Role) Scan(src interface{}) error {
	*r = Role(sprintf("%s", src))
	return nil
}
