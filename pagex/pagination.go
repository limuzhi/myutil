/*
 * @PackageName: pagex
 * @Description:
 * @Author: limuzhi
 * @Date: 2022/12/1 14:53
 */

package pagex

type Pagination struct {
	Page, PageSize int
}

func (p *Pagination) GetPage() int {
	if p.Page <= 0 {
		p.Page = 1
	}
	return p.Page
}

func (p *Pagination) GetPageSize() int {
	if p.PageSize <= 0 {
		p.PageSize = 10
	}
	if p.PageSize > 10000 {
		p.PageSize = 10000
	}
	return p.PageSize
}
