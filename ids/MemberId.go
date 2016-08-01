package ids

type MemberId int

func (id MemberId) Blank() bool   { return id <= 0 }
func (id MemberId) Invalid() bool { return id < 0 }
func (id MemberId) Present() bool { return id > 0 }

func MemberIdIn(ids ...MemberId) []MemberId {
	return []MemberId(ids)
}
