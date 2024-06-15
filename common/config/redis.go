package config

type Redis struct {
	Addr        string `yaml:"addr"`
	Password    string `yaml:"password"`
	Database    int    `yaml:"database"`
	MaxIdle     int    `yaml:"maxIdle"`
	IdleTimeout int    `yaml:"idle_timeout"`
}
