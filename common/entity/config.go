package entity

// Config 配置表
type Config struct {
	Id                int64   `json:"-" gorm:"primaryKey,autoIncrement"`
	InviteBonusPoints int     `json:"invite_bonus_points"`
	QuestionPoints    int     `json:"question_points"`
	DayPoints         int     `json:"day_points"`
	InitPoints        int     `json:"init_points"`
	AdPoints          int     `json:"ad_points"`
	DayAdPoints       int     `json:"day_ad_points"`
	BannerAd          string  `json:"banner_ad"`
	PopAd             string  `json:"pop_ad"`
	ImpelAd           string  `json:"impel_ad"`
	LevelOne          float64 `json:"level_one"`
	LevelSecond       float64 `json:"level_second"`

	// 最大体现
	WithdrawMin float64 `json:"withdraw_min"`
	WithdrawMax float64 `json:"withdraw_max"`
	WithdrawFee int     `json:"withdraw_fee"`

	// 邀请的跳转页面
	InvitePage string `json:"invite_page" gorm:"type:varchar(255);comment:跳转页面"`
	// 参数
	InviteScene string `json:"invite_scene" gorm:"type:varchar(255)"`
	Time
}
