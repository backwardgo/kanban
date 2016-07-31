package ids

type ListId int

func (id ListId) Blank() bool   { return id == 0 }
func (id ListId) Invalid() bool { return id < 0 }
func (id ListId) Present() bool { return id > 0 }
