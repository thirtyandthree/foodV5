package config

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"io/ioutil"
)

type Config struct {
	Server                Server `yaml:"server"`
	CORS                  CORS   `yaml:"cors"`
	Gorm                  Gorm   `yaml:"gorm"`
	MySQL                 MySQL  `yaml:"mysql"`
	WechatMini            Wechat `yaml:"wechat_mini"`
	WechatOfficialAccount Wechat `yaml:"wechat_officialaccount"`
	Redis                 Redis  `yaml:"redis"`
	Log                   Log    `yaml:"log"`
	File                  File   `yaml:"file"`
	Jwt                   Jwt    `yaml:"jwt"`
}

var (
	C = new(Config)
)

func Load(path string) (err error) {
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Printf("配置文件打开失败：%v\n", err)
		return
	}
	err = yaml.Unmarshal(bytes, C)
	if err != nil {
		fmt.Printf("配置文件解析失败：%v\n", err)
		return
	}
	return nil
}
