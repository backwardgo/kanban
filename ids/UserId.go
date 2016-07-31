package ids

type UserId int

func (id UserId) Blank() bool   { return id == 0 }
func (id UserId) Invalid() bool { return id < 0 }
func (id UserId) Present() bool { return id > 0 }
