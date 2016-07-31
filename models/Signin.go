package models

type Signin struct {
	Email    Email  `db:"-" json:"email"`
	Password string `db:"-" json:"password"`
}

func (m *Signin) Errors() Errors {
	m.Normalize()

	e := NewErrors()

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

func (m *Signin) Normalize() {
	m.Email = m.Email.Normalize()
	m.Password = trimSpace(m.Password)
}
