package log

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"path/filepath"
	"runtime"
	"strings"
	"veric-backend/logic/config"
)

var logger *zap.Logger

func init() {
	var err error

	if config.Get().Debug.Enable {
		developmentConfig := zap.NewDevelopmentConfig()
		developmentConfig.EncoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout("15:04:05")
		developmentConfig.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder

		if _, file, _, ok := runtime.Caller(0); ok {
			basePath := filepath.Dir(filepath.Dir(filepath.Dir(file))) + "/"
			developmentConfig.EncoderConfig.EncodeCaller = func(caller zapcore.EntryCaller, encoder zapcore.PrimitiveArrayEncoder) {
				rel, err := filepath.Rel(basePath, caller.File)
				if err != nil {
					encoder.AppendString(caller.FullPath())
				} else {
					if strings.HasPrefix(rel, "vendor/") {
						encoder.AppendString(fmt.Sprintf("[mod] %s:%d", rel[7:], caller.Line))
					} else {
						encoder.AppendString(fmt.Sprintf("%s:%d", rel, caller.Line))
					}
				}
			}
		}

		if !config.Get().Debug.Verbose {
			developmentConfig.Level.SetLevel(zapcore.InfoLevel)
		}

		logger, err = developmentConfig.Build()
		if err != nil {
			panic(err)
		}
	} else {
		logger, err = zap.NewProduction()
		if err != nil {
			panic(err)
		}
	}

	logger.Info("init log config...")
}

func GetLogger() *zap.Logger {
	return logger
}
