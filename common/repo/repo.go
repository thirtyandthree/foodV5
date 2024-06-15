package repo

import (
	"foodV5/common/entity"
	"gorm.io/gorm"
	"time"
)

type Repo struct {
}

// CreateTime 创建时间...
func (*Repo) CreateTime(t *entity.Time) {
	// 创建时间,年月日
	t.Year = uint16(time.Now().Year())
	t.Month = uint16(time.Now().Month())
	t.Day = uint16(time.Now().Day())
	t.CreateTime = time.Now().Format("2006-01-02 15:04:05")
}

// 分页
func paginate(page int, limit int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if page == 0 {
			page = 1
		}
		if limit > 100 {
			limit = 100
		}
		if limit <= 0 {
			limit = 15
		}
		return db.Limit(limit).Offset((page - 1) * limit)
	}
}

// 获取表名字
func tableName(db *gorm.DB, table string) string {
	return db.NamingStrategy.TableName(table)
}
