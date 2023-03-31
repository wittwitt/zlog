package zlog

import "go.uber.org/zap"

type Logger struct {
	baseLogger *zap.SugaredLogger
}

func NewLogger(name string) *Logger {
	return NewLoggerWithSkip(name, 1)
}

func NewLoggerWithSkip(name string, skipN int) *Logger {
	setUp()

	baseLogger, _ := getBaseLogger(name, skipN)
	return &Logger{
		baseLogger: baseLogger,
	}
}

func (p *Logger) Debug(args ...interface{}) {
	p.baseLogger.Debug(args...)
}

func (p *Logger) Info(args ...interface{}) {
	p.baseLogger.Info(args...)
}

func (p *Logger) Warn(args ...interface{}) {
	p.baseLogger.Warn(args...)
}

func (p *Logger) Error(args ...interface{}) {
	p.baseLogger.Error(args...)
}

func (p *Logger) Panic(args ...interface{}) {
	p.baseLogger.Panic(args...)
}

func (p *Logger) Fatal(args ...interface{}) {
	p.baseLogger.Fatal(args...)
}

func (p *Logger) Debugf(format string, args ...interface{}) {
	p.baseLogger.Debugf(format, args...)
}

func (p *Logger) Infof(format string, args ...interface{}) {
	p.baseLogger.Infof(format, args...)
}

func (p *Logger) Warnf(format string, args ...interface{}) {
	p.baseLogger.Warnf(format, args...)
}

func (p *Logger) Errorf(format string, args ...interface{}) {
	p.baseLogger.Errorf(format, args...)
}

func (p *Logger) Panicf(format string, args ...interface{}) {
	p.baseLogger.Panicf(format, args...)
}

func (p *Logger) Fatalf(format string, args ...interface{}) {
	p.baseLogger.Fatalf(format, args...)
}

func (p *Logger) Debugw(msg string, kv ...interface{}) {
	p.baseLogger.Debugw(msg, kv...)
}

func (p *Logger) Infow(msg string, kv ...interface{}) {
	p.baseLogger.Infow(msg, kv...)
}

func (p *Logger) Warnw(msg string, kv ...interface{}) {
	p.baseLogger.Warnw(msg, kv...)
}

func (p *Logger) Errorw(msg string, kv ...interface{}) {
	p.baseLogger.Errorw(msg, kv...)
}

func (p *Logger) DPanicw(msg string, kv ...interface{}) {
	p.baseLogger.DPanicw(msg, kv...)
}

func (p *Logger) Panicw(msg string, kv ...interface{}) {
	p.baseLogger.Panicw(msg, kv...)
}

func (p *Logger) Fatalw(msg string, kv ...interface{}) {
	p.baseLogger.Fatalw(msg, kv...)
}

func (p *Logger) Debugln(args ...interface{}) {
	p.baseLogger.Debugln(args...)
}

func (p *Logger) Infoln(args ...interface{}) {
	p.baseLogger.Infoln(args...)
}

func (p *Logger) Warnln(args ...interface{}) {
	p.baseLogger.Warnln(args...)
}

func (p *Logger) Errorln(args ...interface{}) {
	p.baseLogger.Errorln(args...)
}

func (p *Logger) DPanicln(args ...interface{}) {
	p.baseLogger.DPanicln(args...)
}

func (p *Logger) Panicln(args ...interface{}) {
	p.baseLogger.Panicln(args...)
}

func (p *Logger) Fatalln(args ...interface{}) {
	p.baseLogger.Fatalln(args...)
}
