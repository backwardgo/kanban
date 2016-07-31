package db

import dat "gopkg.in/mgutz/dat.v1"

func NewPager(page, perPage uint) Pager {
	if page < 1 {
		page = 1
	}

	if perPage < 1 {
		perPage = 200
	}

	return &pager{
		page:    page,
		perPage: perPage,
	}
}

type Pager interface {
	Page() uint
	PerPage() uint
	TotalRecords() uint

	refineQuery(*dat.SelectBuilder) *dat.SelectBuilder
	setTotalRecords(uint)
}

type pager struct {
	page         uint
	perPage      uint
	totalRecords uint
}

func (p *pager) Page() uint {
	return p.page
}

func (p *pager) PerPage() uint {
	return p.perPage
}

func (p *pager) TotalRecords() uint {
	return p.totalRecords
}

func (p *pager) refineQuery(query *dat.SelectBuilder) *dat.SelectBuilder {
	var (
		page    = uint64(p.page)
		perPage = uint64(p.perPage)
	)

	return query.Paginate(page, perPage)
}

func (p *pager) setTotalRecords(totalRecords uint) {
	p.totalRecords = totalRecords
}
