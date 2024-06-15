package repo

import (
	"foodV5/common/entity"
	"gorm.io/gorm"
)

type UserRepo struct {
	DB   *gorm.DB
	Repo *Repo
}

func (u *UserRepo) FindById(id int64) (user *entity.User, err error) {
	err = u.DB.Where("id=?", id).Find(&user).Error
	return

}

func (u *UserRepo) CreateWithTx(db *gorm.DB, user *entity.User) error {
	u.Repo.CreateTime(&user.Time)
	return db.Create(user).Error
}

func (u *UserRepo) Update(user *entity.User) error {
	return u.DB.Updates(user).Error
}
func (u *UserRepo) UpdateWithTx(db *gorm.DB, user *entity.User) error {
	return db.Updates(user).Error
}

// FindByMiniOpenId 根据小程序openid查询
func (u *UserRepo) FindByMiniOpenId(openId string) (user *entity.User, err error) {
	err = u.DB.Where("mini_openid=?", openId).Find(&user).Error
	return
}

// FindByUnionId 根据unionId查询,可能是先注册了小程序
func (u *UserRepo) FindByUnionId(unionId string) (user *entity.User, err error) {
	err = u.DB.
		Where("union_id = ?", unionId).
		Find(&user).Error
	return
}

func (u *UserRepo) DeleteByUserId(userId int64) error {
	return u.DB.Delete(&entity.User{Id: userId}).Error
}

// BalanceIncWithTx 增加钱财
func (u *UserRepo) BalanceIncWithTx(db *gorm.DB, userId int64, money float64) error {
	return db.Model(&entity.User{}).
		Where("id=?", userId).
		UpdateColumn("balance", gorm.Expr("balance+?", money)).Error
}

// BalanceDecWithTx 减少
func (u *UserRepo) BalanceDecWithTx(db *gorm.DB, userId int64, money float64) error {
	return db.Model(&entity.User{}).
		Where("id=?", userId).
		UpdateColumn("balance", gorm.Expr("balance-?", money)).Error
}

func (u *UserRepo) IntegralIncWithTx(db *gorm.DB, userId int64, integral int) error {
	return db.Model(&entity.User{}).
		Where("id=?", userId).
		UpdateColumn("integral", gorm.Expr("integral+?", integral)).Error
}

func (u *UserRepo) IntegralDecWithTx(db *gorm.DB, userId int64, integral int) error {
	return db.Model(&entity.User{}).
		Where("id=?", userId).
		UpdateColumn("integral", gorm.Expr("integral-?", integral)).Error
}

func (u *UserRepo) RewardIncWithTx(db *gorm.DB, userId int64, money float64) error {
	return db.Model(&entity.User{}).
		Where("id=?", userId).
		UpdateColumn("reward", gorm.Expr("reward+?", money)).Error
}

// BalanceLockIncWithTx 锁定
func (u *UserRepo) BalanceLockIncWithTx(db *gorm.DB, userId int64, money float64) error {
	return db.Model(&entity.User{}).
		Where("id=?", userId).
		UpdateColumn("balance_lock", gorm.Expr("balance_lock+?", money)).Error
}

func (u *UserRepo) BalanceLockDecWithTx(db *gorm.DB, userId int64, money float64) error {
	return db.Model(&entity.User{}).
		Where("id=?", userId).
		UpdateColumn("balance_lock", gorm.Expr("balance_lock-?", money)).Error
}
