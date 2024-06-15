package logs

import (
	"fmt"
	"foodV5/common/config"
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"time"
)

var Log *zap.SugaredLogger

func NewLog() *zap.SugaredLogger {
	core := zapcore.NewCore(enc(), ws(config.C.Log), enab())
	logger := zap.New(core, zap.AddCaller())
	Log = logger.Sugar()
	return Log
}

func ws(cfg config.Log) zapcore.WriteSyncer {
	logFileName := fmt.Sprintf("./%s/%v.log", cfg.Filename, time.Now().Format(cfg.TimeFormat))
	lumberJackLogger := &lumberjack.Logger{
		Filename:   logFileName,
		MaxSize:    cfg.MaxSize,
		MaxBackups: cfg.MaxBackups,
		MaxAge:     cfg.MaxAge,
	}
	return zapcore.AddSync(lumberJackLogger)
}

func enc() zapcore.Encoder {
	cfg := zap.NewProductionEncoderConfig()
	cfg.TimeKey = "time"
	cfg.EncodeTime = zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05")
	return zapcore.NewJSONEncoder(cfg)
}

func enab() zapcore.LevelEnabler {
	return zapcore.DebugLevel
}
