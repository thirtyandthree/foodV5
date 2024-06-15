package dto

// AccountDto 账户登陆接收信息
type AccountDto struct {
	Code string `json:"code"  binding:"required" label:"code不能为空"`

	From string `json:"from"`
}
