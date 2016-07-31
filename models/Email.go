package models

import (
	"database/sql/driver"
	"regexp"
)

var emailValidator = regexp.MustCompile(`^([^@\s]+)@((?:[-a-z0-9]+\.)+[a-z]{2,})$`)

type Email string

func (e Email) Blank() bool {
	return e == ""
}

func (e Email) Present() bool {
	return !e.Blank()
}

func (e Email) Normalize() Email {
	return Email(toLower(trimSpace(string(e))))
}

func (e Email) Invalid() bool {
	return !e.Valid()
}

func (e Email) Valid() bool {
	return emailValidator.MatchString(string(e))
}

func (id Email) Value() (driver.Value, error) {
	return string(id), nil
}

func (id *Email) Scan(src interface{}) error {
	*id = Email(sprintf("%s", src))
	return nil
}
