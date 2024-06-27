package logging

import (
	"github.com/meiti-x/golang-web-api/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

type zapLogger struct {
	cfg    *config.Config
	logger *zap.SugaredLogger
}

// warn implements Logger.
func (l *zapLogger) warn(cat Category, sub SubCategory, msg string, extra map[ExtraKey]interface{}) {
	panic("unimplemented")
}

// warnf implements Logger.
func (l *zapLogger) warnf(template string, args ...interface{}) {
	panic("unimplemented")
}

var logLevelMap = map[string]zapcore.Level{
	"debug": zap.DebugLevel,
	"info":  zap.InfoLevel,
	"warn":  zap.InfoLevel,
	"error": zap.ErrorLevel,
	"fatal": zap.FatalLevel,
}

func (l *zapLogger) getLoglevel() zapcore.Level {
	level, exist := logLevelMap[l.cfg.Logger.Level]

	if !exist {
		return zapcore.DebugLevel
	}
	return level
}

func (l *zapLogger) Init() {
	w := zapcore.AddSync(&lumberjack.Logger{
		Filename:   l.cfg.Logger.FilePath,
		MaxSize:    1,
		MaxAge:     5,
		MaxBackups: 10,
		LocalTime:  true,
		Compress:   true,
	})
	config := zap.NewProductionEncoderConfig()
	config.EncodeTime = zapcore.ISO8601TimeEncoder

	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(config),
		w,
		l.getLoglevel(),
	)
	logger := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1), zap.AddStacktrace(zapcore.ErrorLevel)).Sugar()

	l.logger = logger
}
func (l *zapLogger) Debug(cat Category, sub SubCategory, msg string, extra map[ExtraKey]interface{}) {
	params := prepareLogInfo(cat, sub, extra)

	l.logger.Debugw(msg, params...)
}

func (l *zapLogger) Debugf(template string, args ...interface{}) {
	l.logger.Debugf(template, args)
}

func (l *zapLogger) Info(cat Category, sub SubCategory, msg string, extra map[ExtraKey]interface{}) {
	params := prepareLogInfo(cat, sub, extra)
	l.logger.Infow(msg, params...)
}

func (l *zapLogger) Infof(template string, args ...interface{}) {
	l.logger.Infof(template, args)
}

func (l *zapLogger) Warn(cat Category, sub SubCategory, msg string, extra map[ExtraKey]interface{}) {
	params := prepareLogInfo(cat, sub, extra)
	l.logger.Warnw(msg, params...)
}

func (l *zapLogger) Warnf(template string, args ...interface{}) {
	l.logger.Warnf(template, args)
}

func (l *zapLogger) Error(cat Category, sub SubCategory, msg string, extra map[ExtraKey]interface{}) {
	params := prepareLogInfo(cat, sub, extra)
	l.logger.Errorw(msg, params...)
}

func (l *zapLogger) Errorf(template string, args ...interface{}) {
	l.logger.Errorf(template, args)
}

func (l *zapLogger) Fatal(cat Category, sub SubCategory, msg string, extra map[ExtraKey]interface{}) {
	params := prepareLogInfo(cat, sub, extra)
	l.logger.Fatalw(msg, params...)
}

func (l *zapLogger) Fatalf(template string, args ...interface{}) {
	l.logger.Fatalf(template, args)
}

func newZapLogger(cfg *config.Config) *zapLogger {
	logger := &zapLogger{cfg: cfg}

	logger.Init()
	return logger
}

func prepareLogInfo(cat Category, sub SubCategory, extra map[ExtraKey]interface{}) []interface{} {
	if extra == nil {
		extra = make(map[ExtraKey]interface{})
	}
	extra["Category"] = cat
	extra["SubCategory"] = sub

	return logParamsToZapParams(extra)
}
