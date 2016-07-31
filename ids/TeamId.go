package ids

type TeamId int

func (id TeamId) Blank() bool   { return id == 0 }
func (id TeamId) Invalid() bool { return id < 0 }
func (id TeamId) Present() bool { return id > 0 }

func TeamIdIn(ids ...TeamId) []TeamId {
	return []TeamId(ids)
}
