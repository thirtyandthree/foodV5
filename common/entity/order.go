package entity

import "time"

type GoodsInfo struct {
	OrderName string `json:"order_name" gorm:"type:varchar(255)"` // 订单商品

	Price float64 `json:"price"` // 单价格

	Freight float64 `json:"freight" gorm:"type:decimal(18,2);default:0;comment:配送费"` // 订单下的配送费，避免因为商品修改了而改变
}

// Order 订单
type Order struct {
	Id int64 `json:"id" gorm:"primaryKey;autoIncrement;comment:订单ID"`

	UserId int64 `json:"user_id"` // 哪个用户下的订单

	RiderId int64 `json:"rider_id"` // 骑手id-哪个骑手接的订单

	GoodsId int64 `json:"goods_id" gorm:"not null"` // 商品id

	BusinessId int64 `json:"business_id"` // 商家id

	TimeStampId int64 `json:"time_stamp_id"` // 时间选择的id

	Number string `json:"number" gorm:"type:varchar(100)"` // 自己内部订单编号

	GoodsInfo

	FinalAmount float64 `gorm:"column:final_amount" json:"final_amount"` // 订单总价格
	// 订单状态
	Status uint8 `gorm:"column:status;default:0;type:int(11)" json:"status"` // 状态0关闭1启用

	StatusText string `json:"status_text" gorm:"type:varchar(255)"` // 状态文字

	PayTime *time.Time `json:"pay_time"  gorm:"type:datetime;null"` //支付时间

	// 是否被接了该订单,默认是0
	IsConnect uint8 `json:"is_connect" gorm:"default:0"`

	TransactionNumber string `gorm:"column:transaction_number" json:"transaction_number"` // 流水号,腾讯回调的号

	PayAmount float64 `json:"pay_amount"` // 实际成交支付金额

	ExpireTime int64 `gorm:"column:expire_time" json:"expire_time"` // 超时时间

	PrintStatus uint8 `json:"print_status" gorm:"default:0"` // 打印状态,是否打印...

	Taste string `json:"taste"` // 选择的口味

	Path string `json:"path"` // 配送地址

	Phone string `json:"phone"` // 联系方式

	Name string `json:"name"` // 收货人

	Closed       uint8 `gorm:"column:closed" json:"closed"`    // 是否关闭
	RefundStatus uint8 `json:"refund_status" gorm:"default:0"` // 退款,默认0

	// 退款累计金额
	RefundAmount float64 `json:"refund_amount"`
	// 订单真实所剩金额
	ResidueAmount float64 `json:"residue_amount"`
	// 是否展示,软删除就是这个样子,1展示2隐藏
	Allow uint8 `json:"allow"`

	Time
}
