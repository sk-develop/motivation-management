package logger

import (
	stackdriver "github.com/tommy351/zap-stackdriver"
	"go.uber.org/zap"
)

var (
	l *zap.Logger
	s *zap.SugaredLogger

	Debug     func(args ...interface{})
	Info      func(args ...interface{})
	Warn      func(args ...interface{})
	Error     func(args ...interface{})
	Debugf    func(tmpl string, args ...interface{})
	Infof     func(tmpl string, args ...interface{})
	Warnf     func(tmpl string, args ...interface{})
	Errorf    func(tmpl string, args ...interface{})
	InfoQuick func(msg string, fields ...zap.Field)
	WarnQuick func(msg string, fields ...zap.Field)
)

func SetUp() *zap.Logger {
	config := zap.NewProductionConfig()
	config.EncoderConfig = stackdriver.EncoderConfig

	var err error
	l, err = config.Build()
	if err != nil {
		panic(err)
	}
	s = l.Sugar()

	Debug = s.Debug
	Info = s.Info
	Warn = s.Warn
	Error = s.Error
	Debugf = s.Debugf
	Infof = s.Infof
	Warnf = s.Warnf
	Errorf = s.Errorf
	InfoQuick = l.Info
	WarnQuick = l.Warn

	return l
}
