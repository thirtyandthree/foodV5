package constant

const InviteRewardDone = 1
const InviteRewardPending = 2

var InviteReward = map[uint8]string{
	InviteRewardDone:    "已奖励",
	InviteRewardPending: "未奖励",
}
