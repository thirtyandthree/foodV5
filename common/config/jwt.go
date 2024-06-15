package config

type Jwt struct {
	Key         string `yaml:"key"`
	TokenPrefix string `yaml:"token_prefix"`
	// 超时时间
	ExpireTime int `yaml:"expire_time"`
}
