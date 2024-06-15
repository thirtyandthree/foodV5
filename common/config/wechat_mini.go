package config

type Wechat struct {
	// 微信公众号
	AppID string `yaml:"app_id"`

	AppSecret string `yaml:"app_secret"`

	Token string `json:"token"`
	// 支付商
	MchId string `yaml:"mch_id"`

	MchKey string `yaml:"mch_key"`

	MchNumber string `yaml:"mch_number"`
	// 私钥路径
	PrivateKey string `yaml:"private_key"`
	Domain     string `yaml:"domain"`
}
