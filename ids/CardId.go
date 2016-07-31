package ids

type CardId int

func (id CardId) Blank() bool   { return id == 0 }
func (id CardId) Invalid() bool { return id < 0 }
func (id CardId) Present() bool { return id > 0 }

func CardIdIn(ids ...CardId) []CardId {
	return []CardId(ids)
}
