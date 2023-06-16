package utils

import (
	"gorm.io/gorm"
)

type Page struct {
	Page    int   `form:"page"`
	PerPage int   `form:"per_page"`
	Desc    bool  `form:"desc"`
	PageNum int64 `form:"-"`
	Total   int64 `form:"-"`
}

func (p *Page) LimitPageArgs() {
	if p.Page == 0 {
		p.Page = 1
	}

	switch {
	case p.PerPage > 100:
		p.PerPage = 100
	case p.PerPage <= 0:
		p.PerPage = 10
	}
}

func (p *Page) Paginate() func(db *gorm.DB) *gorm.DB {
	p.LimitPageArgs()

	return func(db *gorm.DB) *gorm.DB {
		offset := (p.Page - 1) * p.PerPage
		q := db.Offset(offset).Limit(p.PerPage)
		if p.Desc {
			q.Order("id desc")
		}

		return q
	}
}
