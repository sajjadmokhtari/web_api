package logging

import (
	"GOLANG_CLEAN_WEB_API/src/config"

	"gopkg.in/natefinch/lumberjack.v2"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// test comment
var logLevelMap = map[string]zapcore.Level{
	"debug": zapcore.DebugLevel,
	"info":  zapcore.InfoLevel,
	"warn":  zapcore.WarnLevel,
	"error": zapcore.ErrorLevel,
	"fatal": zapcore.FatalLevel,
}// نگاشت سطح‌های لاگ

type ZapLogger struct {
	cfg    *config.Config
	logger *zap.SugaredLogger
}

func NewZapLogger(cfg *config.Config) *ZapLogger {
	logger := &ZapLogger{cfg: cfg}
	logger.Init()
	return logger
}

func (l *ZapLogger) getLogLevel() zapcore.Level {
	level, exists := logLevelMap[l.cfg.Logger.Level]
	if !exists {
		return zapcore.DebugLevel
	}
	return level
}

func (l *ZapLogger) Init() {
	w := zapcore.AddSync(&lumberjack.Logger{
		Filename:   l.cfg.Logger.Filepath,// ادرس اون فایلی که توی داکر یمل هست
		MaxSize:    1, // megabytes
		MaxAge:     5, // days
		LocalTime:  true,
		MaxBackups: 10,
		Compress:   true,
	})

	config := zap.NewProductionEncoderConfig()
	config.EncodeTime = zapcore.ISO8601TimeEncoder

	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(config),
		w,
		l.getLogLevel(),
	)

	logger := zap.New(core, zap.AddCaller(),
		zap.AddCallerSkip(1), zap.AddStacktrace(zapcore.ErrorLevel)).Sugar()
	l.logger = logger

}
func (l *ZapLogger) Debug(cat Category, sub SubCategory, msg string, extra map[ExtraKey]interface{}) {
	params := prepareLogKeys(extra, cat, sub)
	l.logger.Debugw(msg, params...)
}
func (l *ZapLogger) Debugf(template string, args ...interface{}) {
	l.logger.Debugf(template, args)

}
func (l *ZapLogger) Info(cat Category, sub SubCategory, msg string, extra map[ExtraKey]interface{}) {
	params := prepareLogKeys(extra, cat, sub)
	l.logger.Infow(msg, params...)
}
func (l *ZapLogger) Infof(template string, args ...interface{}) {
	l.logger.Infof(template, args)

}
func (l *ZapLogger) Warn(cat Category, sub SubCategory, msg string, extra map[ExtraKey]interface{}) {
	params := prepareLogKeys(extra, cat, sub)
	l.logger.Warnw(msg, params...)
}
func (l *ZapLogger) Warnf(template string, args ...interface{}) {
	l.logger.Warnf(template, args)

}
func (l *ZapLogger) Error(cat Category, sub SubCategory, msg string, extra map[ExtraKey]interface{}) {
	params := prepareLogKeys(extra, cat, sub)
	l.logger.Errorw(msg, params...)
}
func (l *ZapLogger) Errorf(template string, args ...interface{}) {
	l.logger.Errorf(template, args)

}
func (l *ZapLogger) Fatal(cat Category, sub SubCategory, msg string, extra map[ExtraKey]interface{}) {
	params := prepareLogKeys(extra, cat, sub)

	l.logger.Fatalw(msg, params...)
}
func (l *ZapLogger) Fatalf(template string, args ...interface{}) {
	l.logger.Fatalf(template, args)

}
func prepareLogKeys(extra map[ExtraKey]interface{}, cat Category, sub SubCategory) []interface{} {
	if extra == nil {
		extra = make(map[ExtraKey]interface{}, 0)
	}
	extra["Category"] = cat
	extra["Sub"] = sub
	params := mapToZapParams(extra)
    return params
}
