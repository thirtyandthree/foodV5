package entity

// 积分表
type Integral struct {
	Id          int64  `json:"id" gorm:"primaryKey,autoIncrement"`
	UserId      int64  `json:"user_id"` // 用户id
	Act         uint8  `json:"act"`     // 操作
	ActText     string `json:"act_text" gorm:"-"`
	Integral    int    `json:"integral"`    // 积分
	Balance     int    `json:"balance"`     // 金额
	Description string `json:"description"` // 说明

	Time
}
