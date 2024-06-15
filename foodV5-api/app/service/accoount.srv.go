package service

import (
	"context"
	"fmt"
	"foodV5/common/config"
	"foodV5/common/entity"
	"foodV5/common/pkg/constant"
	"foodV5/common/pkg/errors"
	"foodV5/common/pkg/jwt"
	"foodV5/common/pkg/logs"
	"foodV5/common/repo"
	"foodV5/common/service"
	"github.com/go-redis/redis/v8"
	"github.com/silenceper/wechat/v2/miniprogram"
	"gorm.io/gorm"
	"time"
)

// AccountService 账户服务
type AccountService struct {
	Mini          *miniprogram.MiniProgram
	Redis         *redis.Client
	DB            *gorm.DB
	UserRepo      *repo.UserRepo
	ConfigRepo    *repo.ConfigRepo
	InviteRepo    *repo.InviteRepo
	InviteService *service.InviteService
	IntegralRepo  *repo.IntegralRepo
}

// MiniLogin 小程序登陆
func (accountService *AccountService) MiniLogin(jsCode string, inviteFrom int64) (token string, err error) {
	session, err := accountService.Mini.GetAuth().Code2Session(jsCode)
	if err != nil {
		err = errors.WechatGetOpenIdError
		return
	}
	// 从数据库查找用户,如果用户是union的话,没有使用过小程序，那么接着走小程序逻辑
	user, err := accountService.findUserInfo(session.OpenID, session.UnionID)
	// 用户不存在自动注册
	if user.Id <= 0 {
		user, err = accountService.autoCreate(session.OpenID, session.UnionID, inviteFrom)
		if err != nil {
			return
		}
	}
	// 用户🈶union但是没有注册小程序
	if user.MiniOpenid == "" {
		// 更新状态,更新用户的小程序openid
		err := accountService.updateWechatWithUser(user, session.OpenID, inviteFrom)
		if err != nil {
			return "", err
		}
	}
	// 返回登陆凭证
	token, err = jwt.GetTokenByUserId(user.Id)
	return
}

// 更新微信配置信息
func (accountService *AccountService) updateWechatWithUser(user *entity.User, openId string, inviteFrom int64) (err error) {
	// 给邀请人奖励积分
	conf, err := accountService.ConfigRepo.FindOne()
	if err != nil {
		logs.Log.Error(err)
		err = errors.UserCreateError
		return
	}
	tx := accountService.DB.Begin()
	user.MiniOpenid = openId // 设置小程序openid
	// 前往更新,设置奖励
	err = accountService.UserRepo.UpdateWithTx(tx, user)
	if err != nil {
		logs.Log.Error(err)
		err = errors.UserUpdateError
		tx.Rollback()
		return err
	}
	// 添加信息

	// 处理邀请好友
	if err = accountService.invite(tx, conf, user.Id, inviteFrom); err != nil {
		tx.Rollback()
		logs.Log.Error(err)
		err = errors.UserCreateError
		return
	}
	tx.Commit()
	return nil
}

// 查询微信用户信息
func (accountService *AccountService) findUserInfo(openId, unionId string) (user *entity.User, err error) {
	user, err = accountService.UserRepo.FindByMiniOpenId(openId)
	if err != nil {
		err = errors.DataQuery
		return
	}

	// 根据微信openid查询没有的话...
	if user.Id > 0 {
		return
	}
	// 根据unionId查询
	user, err = accountService.UserRepo.FindByUnionId(unionId)
	if err != nil {
		err = errors.UserLoginError
		return
	}
	return
}

// 自动创建用户信息
func (accountService *AccountService) autoCreate(openId, unionID string, inviteFrom int64) (user *entity.User, err error) {
	// 给邀请人奖励多少钱
	conf, err := accountService.ConfigRepo.FindOne()
	if err != nil {
		logs.Log.Error(err)
		err = errors.UserCreateError
		return
	}
	// 开启事物,进行更新
	tx := accountService.DB.Begin()
	user = &entity.User{
		MiniOpenid: openId,
		UnionId:    unionID,
	}
	// 创建用户表信息
	if err = accountService.UserRepo.CreateWithTx(tx, user); err != nil {
		tx.Rollback()
		logs.Log.Error(err)
		err = errors.UserCreateError
		return
	}

	// 处理邀请好友
	if err = accountService.invite(tx, conf, user.Id, inviteFrom); err != nil {
		tx.Rollback()
		logs.Log.Error(err)
		err = errors.UserCreateError
		return
	}
	tx.Commit()

	return
}

// 邀请好友,给予奖励,奖励积分或者余额什么的
func (accountService *AccountService) invite(tx *gorm.DB, conf *entity.Config, userId, inviteFrom int64) (err error) {
	if inviteFrom <= 0 {
		return nil
	}
	// 查找邀请用户
	user, err := accountService.UserRepo.FindById(inviteFrom)
	if err != nil {
		logs.Log.Error(err)
		return
	}
	if user.Id <= 0 {
		return
	}

	// 关联邀请关系
	// 绑定最上层关系
	levelOne := int64(0)
	userInvite, err := accountService.InviteRepo.FindByTo(inviteFrom)
	if err != nil {
		return
	}
	if userInvite.FromUser > 0 {
		levelOne = userInvite.FromUser
	}
	invite := &entity.Invite{
		FromUser: inviteFrom,
		ToUser:   userId,
		IsReward: constant.InviteRewardPending,
		LevelOne: levelOne,
	}
	// 新增邀请记录
	if err = accountService.InviteService.Plus(tx, invite); err != nil {
		return
	}

	// 判断当天积分是否上限
	sum, err := accountService.IntegralRepo.FindUserDayIntegral(inviteFrom)
	if err != nil || sum >= conf.DayPoints {
		return
	}
	// 添加积分
	if err = accountService.UserRepo.IntegralIncWithTx(tx, inviteFrom, conf.InviteBonusPoints); err != nil {
		return
	}

	if conf.InviteBonusPoints > 0 {
		integral := &entity.Integral{
			UserId:      inviteFrom,
			Act:         constant.FinanceActInCome,
			Integral:    conf.InviteBonusPoints,
			Balance:     user.Integral + conf.InviteBonusPoints,
			Description: "邀请好友注册奖励",
		}
		// 插入管理
		if err = accountService.IntegralRepo.CreateWithTx(tx, integral); err != nil {
			return
		}
	}
	return
}

// SetTokenInfo 设置token信息
func (accountService *AccountService) SetTokenInfo(token string, duration time.Duration) {
	// 1.0 获取token的id
	uid, _ := jwt.GetUserIdByToken(token)
	accountService.Redis.SetEX(context.TODO(), fmt.Sprintf("%s%d", config.C.Jwt.TokenPrefix, uid), "", duration)
}
