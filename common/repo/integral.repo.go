package repo

import (
	"foodV5/common/dto"
	"foodV5/common/entity"
	"gorm.io/gorm"
	"strconv"
	"time"
)

type IntegralRepo struct {
	DB   *gorm.DB
	Repo *Repo
}

// List 获取用户积分列表
func (i *IntegralRepo) List(search *dto.IntegralSearch) (lists []*entity.Integral, err error) {
	err = i.DB.Scopes(paginate(search.Page, 15)).
		Where("user_id=?", search.UserId).
		Order("create_time desc").
		Find(&lists).Error
	return
}

func (i *IntegralRepo) CreateWithTx(db *gorm.DB, integral *entity.Integral) error {
	i.Repo.CreateTime(&integral.Time)
	return db.Create(integral).Error
}

// FindUserDayIntegral 获取用户的日积分量
func (i *IntegralRepo) FindUserDayIntegral(userId int64) (num int, err error) {
	t := time.Now()
	monthStr := time.Now().Format("01")
	month, _ := strconv.Atoi(monthStr)
	err = i.DB.Model(&entity.Integral{}).
		Where("user_id=?", userId).
		Where("year=?", t.Year()).
		Where("month=?", month).
		Where("day=?", t.Day()).
		Select("COALESCE(SUM(`integral`),0)").
		Scan(&num).Error
	return
}
