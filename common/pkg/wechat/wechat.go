package wechat

import (
	"context"
	"foodV5/common/config"
	"foodV5/common/pkg/logs"
	"github.com/silenceper/wechat/v2"
	"github.com/silenceper/wechat/v2/cache"
	"github.com/silenceper/wechat/v2/miniprogram"
	miniConfig "github.com/silenceper/wechat/v2/miniprogram/config"
	"github.com/silenceper/wechat/v2/officialaccount"
	config2 "github.com/silenceper/wechat/v2/officialaccount/config"
	"github.com/wechatpay-apiv3/wechatpay-go/core"
	"github.com/wechatpay-apiv3/wechatpay-go/core/option"
	"github.com/wechatpay-apiv3/wechatpay-go/utils"
)

func NewWeChatClient() *core.Client {
	mini := config.C.WechatMini
	mchPrivateKey, err := utils.LoadPrivateKeyWithPath(mini.PrivateKey)
	if err != nil {
		logs.Log.Fatal("加载私钥文件错误" + err.Error())
		return nil
	}

	ctx := context.Background()
	// 使用商户私钥等初始化 client，并使它具有自动定时获取微信支付平台证书的能力
	opts := []core.ClientOption{
		option.WithWechatPayAutoAuthCipher(mini.MchId, mini.MchNumber, mchPrivateKey, mini.MchKey),
	}
	client, err := core.NewClient(ctx, opts...)
	if err != nil {
		logs.Log.Fatalf("创建微信支付客户端失败,原因:%v", err)
	}

	return client
}

// NewWechatMini 创建微信的,这个用来做登陆获取信息啥的
func NewWechatMini() *miniprogram.MiniProgram {
	wc := wechat.NewWechat()

	redisOpts := &cache.RedisOpts{
		Host:        config.C.Redis.Addr,
		Password:    config.C.Redis.Password,
		Database:    0,
		MaxIdle:     10,
		IdleTimeout: 60, //second
	}
	redisCache := cache.NewRedis(context.Background(), redisOpts)
	cfg := &miniConfig.Config{
		AppID:     config.C.WechatMini.AppID,
		AppSecret: config.C.WechatMini.AppSecret,
		Cache:     redisCache,
	}
	return wc.GetMiniProgram(cfg)
}

func NewWechatOfficialAccount() *officialaccount.OfficialAccount {
	wc := wechat.NewWechat()
	redisOpts := &cache.RedisOpts{
		Host:        config.C.Redis.Addr,
		Password:    config.C.Redis.Password,
		Database:    0,
		MaxIdle:     10,
		IdleTimeout: 60, //second
	}
	redisCache := cache.NewRedis(context.Background(), redisOpts)
	cfg := &config2.Config{
		AppID:     config.C.WechatOfficialAccount.AppID,
		AppSecret: config.C.WechatOfficialAccount.AppSecret,
		Token:     config.C.WechatOfficialAccount.Token,
		Cache:     redisCache,
	}
	return wc.GetOfficialAccount(cfg)
}
