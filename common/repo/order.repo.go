package repo

import (
	"foodV5/common/dto"
	"foodV5/common/entity"
	"gorm.io/gorm"
)

type OrderRepo struct {
	Repo *Repo
	Db   *gorm.DB
}

func (o *OrderRepo) ExecuteOrder(db *gorm.DB) (order *entity.Order, err error) {
	err = db.Find(order).Error
	return
}

func (o *OrderRepo) ExecuteOrders(db *gorm.DB) (order []*entity.Order, err error) {
	err = db.Find(&order).Error
	return
}
func (o *OrderRepo) List(userId int64, dto dto.Dto) (order []*entity.Order, err error) {
	err = o.Db.Scopes(paginate(dto.Page, dto.Limit)).Where("user_id = ?", userId).
		Order("create_time desc").
		Find(&order).Error
	return
}

// Create 创建订单
func (o *OrderRepo) Create(order *entity.Order) error {
	o.Repo.CreateTime(&order.Time)
	return o.Db.Create(order).Error
}

func (o *OrderRepo) CreateWithTx(tx *gorm.DB, order *entity.Order) error {
	o.Repo.CreateTime(&order.Time)
	return tx.Create(order).Error
}

// FindById 根据订单id查询
func (o *OrderRepo) FindById(id int64) (order *entity.Order, err error) {
	err = o.Db.Where("id = ?", id).Find(&order).Error
	return
}

// FindByNumber 根据订单号查询订单
func (o *OrderRepo) FindByNumber(number string) (order *entity.Order, err error) {
	err = o.Db.Where("number = ?", number).Find(&order).Error
	return
}

// FindByNumberWithTx 查询，带事物
func (o *OrderRepo) FindByNumberWithTx(tx *gorm.DB, number string) (order *entity.Order, err error) {
	err = tx.Where("number = ?", number).Find(&order).Error
	return
}

func (o *OrderRepo) UpdateWithTx(db *gorm.DB, order *entity.Order) error {
	return db.Updates(order).Error
}
