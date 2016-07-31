package models

type Signup struct {
	FirstName string `db:"first_name" json:"firstName"`
	LastName  string `db:"last_name" json:"lastName"`
	Email     Email  `db:"email" json:"email"`
	Password  string `db:"-" json:"password"`
}

func (m *Signup) Errors() Errors {
	m.Normalize()

	e := NewErrors()

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

	switch {
	case m.Password == "":
		e["password"] = "is required"
	case len(m.Password) < minPasswordLen:
		e["password"] = "is too short"
	}

	return e
}

func (m *Signup) Normalize() {
	m.FirstName = trimSpace(m.FirstName)
	m.LastName = trimSpace(m.LastName)
	m.Password = trimSpace(m.Password)
	m.Email = m.Email.Normalize()
}
