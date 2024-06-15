package errors

var (
	Success       = 0
	ErrorCode     = 100001 // 参数错误
	BusinessCode  = 100002 // 业务错误
	AuthorizeCode = 401001 // 登录认证
	TimesOutCode  = 401002 // VIP过期

	DataCreate = Error(ErrorCode, "数据创建失败")
	DataQuery  = Error(ErrorCode, "数据查询失败")
	DataDelete = Error(ErrorCode, "删除失败")
	DataUpdate = Error(ErrorCode, "数据更新失败")
	DataEmpty  = Error(ErrorCode, "该数据不存在")

	UserUnLoginError     = Error(AuthorizeCode, "请登录")
	TokenError           = Error(AuthorizeCode, "token无效,请重新登陆")
	UserLoginError       = Error(BusinessCode, "登陆失败")
	UserCreateError      = Error(BusinessCode, "自动注册失败")
	UserUpdateError      = Error(BusinessCode, "用户功能扩展失败")
	WechatGetOpenIdError = Error(BusinessCode, "微信授权登陆失败")

	WithdrawCreate  = Error(BusinessCode, "申请提现失败")
	WithdrawBalance = Error(BusinessCode, "余额不足您提现的金额")

	WechatOrderCreate       = Error(BusinessCode, "微信订单创建失败")
	WechatNotifyProcessed   = Error(BusinessCode, "微信支付回调处理失败，该订单已处理")
	WechatNotifyAmountWrong = Error(BusinessCode, "微信支付回调处理失败，金额不一致")

	QuestionNotVipOrIntegral = Error(TimesOutCode, "您的免费次数已用完请邀请好友或者开通会员")

	MsgTxtCheck = Error(BusinessCode, "内容检查失败")
)

type Err struct {
	Code int
	Msg  string
}

func (e *Err) Error() string {
	return e.Msg
}

func Error(code int, msg string) *Err {
	return &Err{
		Code: code,
		Msg:  msg,
	}
}
