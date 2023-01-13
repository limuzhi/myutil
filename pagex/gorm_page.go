/*
 * @PackageName: pagex
 * @Description:
 * @Author: limuzhi
 * @Date: 2022/12/1 14:58
 */

package pagex

import "gorm.io/gorm"

func Paginate(page, pageSize int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if page <= 0 {
			page = 1
		}
		if pageSize <= 0 {
			pageSize = 10
		}
		if pageSize > 10000 {
			pageSize = 10000
		}
		return db.Offset((page - 1) * pageSize).Limit(pageSize)
	}
}
