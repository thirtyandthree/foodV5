package config

type Gorm struct {
	Debug             bool   `yaml:"debug"`
	DBType            string `yaml:"db_type"`
	MaxLifetime       int    `yaml:"max_lifetime"`
	MaxOpenConns      int    `yaml:"max_open_conns"`
	MaxIdleConns      int    `yaml:"max_idle_conns"`
	TablePrefix       string `yaml:"table_prefix"`
	EnableAutoMigrate bool   `yaml:"enable_auto_migrate"`
}
