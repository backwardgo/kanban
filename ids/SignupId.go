package ids

type SignupId int

func (id SignupId) Blank() bool   { return id == 0 }
func (id SignupId) Invalid() bool { return id < 0 }
func (id SignupId) Present() bool { return id > 0 }
