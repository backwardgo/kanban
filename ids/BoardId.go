package ids

type BoardId int

func (id BoardId) Blank() bool   { return id <= 0 }
func (id BoardId) Invalid() bool { return id < 0 }
func (id BoardId) Present() bool { return id > 0 }

func BoardIdIn(ids ...BoardId) []BoardId {
	return []BoardId(ids)
}
