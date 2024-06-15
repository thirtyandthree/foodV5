package entity

// User 用户表
type User struct {
	Id int64 `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`

	MiniOpenid string `json:"mini_openid" gorm:"size:100;comment:小程序用户openid"`

	UnionId string `json:"union_id" gorm:"size:100;comment:之间的unionId"`

	OfficialOpenid string `json:"official_openid" gorm:"size:100;comment:公众号用户openid"`

	// 积分
	Integral int `json:"integral"`
	// 钱财,余额
	Balance float64 `json:"balance"`
	// 锁定余额
	BalanceLock float64 `json:"balance_lock"`
	// 返现金额
	CashBack float64 `json:"cash_back"`
	// 已提现金额
	Withdraw float64 `json:"withdraw"`

	InviteReward float64 `json:"invite_reward"` // 邀请奖励比例
	Reward       float64 `json:"reward"`        // 返现总金额

	// 是不是会员
	//IsVip      uint8      `json:"is_vip"`
	//VipExpires *time.Time `json:"-"`
	// 超时时间
	//VipExpire string `json:"vip_expires" gorm:"-"`
	Uid      string `json:"uid" gorm:"-"`
	Avatar   string `json:"avatar"`
	UserName string `json:"user_name"`
	//Phone    string `json:"phone"`
	// 时间
	Time
}

type UserFinance struct {
	Id          int64   `json:"id" gorm:"primaryKey,autoIncrement"`
	UserId      int64   `json:"user_id"`
	Act         uint8   `json:"act"`
	ActText     string  `json:"act_text" gorm:"-"`
	Amount      float64 `json:"amount"`
	Balance     float64 `json:"balance"`
	Description string  `json:"description"`
	OrderNo     string  `json:"order_no"`
	Time
}

// UserAct  完整的用户操作
type UserAct struct {
	Id int64 `json:"id" gorm:"primaryKey,autoIncrement"`
	// 用户的id
	UserId int64 `json:"user_id"`
	// 操作
	Act uint8 `json:"act"`
	// 操作信息
	ActText string `json:"act_text" gorm:"-"`

	// 说明备注
	Description string `json:"description" gorm:"type:varchar(255)"`
	Time
}
