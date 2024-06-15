package repo

import (
	"foodV5/common/entity"
	"gorm.io/gorm"
)

type ConfigRepo struct {
	DB *gorm.DB
}

// FindOne 查询最新的配置,全局的
func (c *ConfigRepo) FindOne() (conf *entity.Config, err error) {
	err = c.DB.Order("id desc").Find(&conf).Error
	return
}
