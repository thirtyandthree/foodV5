package constant

const FinanceActInCome = 1
const FinanceActExpend = 2

var FinanceAct = map[uint8]string{
	FinanceActInCome: "收入",
	FinanceActExpend: "支出",
}

const WithdrawStatusSuccess = 1
const WithdrawStatusWait = 2
const WithdrawStatusReject = 3

var WithdrawStatus = map[int]string{
	WithdrawStatusSuccess: "成功",
	WithdrawStatusWait:    "待处理",
	WithdrawStatusReject:  "拒绝",
}

const WithdrawTypeAlipay = 1
const WithdrawTypeWechat = 2

var WithdrawType = map[int]string{
	WithdrawTypeAlipay: "支付宝",
	WithdrawTypeWechat: "微信",
}
