package config

type Server struct {
	Host               string `yaml:"host"`
	Port               int    `yaml:"port"`
	ShutdownTimeout    int    `yaml:"shutdown_timeout"`
	MaxContentLength   int64  `yaml:"max_content_length"`
	MaxReqLoggerLength int    `yaml:"max_req_logger_length"`
	MaxResLoggerLength int    `yaml:"max_res_logger_length"`
	IsClose            bool   `yaml:"is_close"`
	CloseMessage       string `yaml:"close_message"`
}
