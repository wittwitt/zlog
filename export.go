package zlog

var std *Logger

func init() {
	std = NewLoggerWithSkip("", 2) // export warp 1 layer ( export.Debug-> std.Debug), so +1
}

func Debug(args ...interface{}) {
	std.Debug(args...)
}

func Info(args ...interface{}) {
	std.Info(args...)
}

func Warn(args ...interface{}) {
	std.Warn(args...)
}

func Error(args ...interface{}) {
	std.Error(args...)
}

func Panic(args ...interface{}) {
	std.Panic(args...)
}

func Fatal(args ...interface{}) {
	std.Fatal(args...)
}

//

func Debugf(format string, args ...interface{}) {
	std.Debugf(format, args...)
}

func Infof(format string, args ...interface{}) {
	std.Infof(format, args...)
}

func Warnf(format string, args ...interface{}) {
	std.Warnf(format, args...)
}

func Errorf(format string, args ...interface{}) {
	std.Errorf(format, args...)
}

func Panicf(format string, args ...interface{}) {
	std.Panicf(format, args...)
}

func Fatalf(format string, args ...interface{}) {
	std.Fatalf(format, args...)
}

//

func Debugw(msg string, kv ...interface{}) {
	std.Debugw(msg, kv...)
}

func Infow(msg string, kv ...interface{}) {
	std.Infow(msg, kv...)
}

func Warnw(msg string, kv ...interface{}) {
	std.Warnw(msg, kv...)
}

func Errorw(msg string, kv ...interface{}) {
	std.Errorw(msg, kv...)
}

func DPanicw(msg string, kv ...interface{}) {
	std.DPanicw(msg, kv...)
}

func Panicw(msg string, kv ...interface{}) {
	std.Panicw(msg, kv...)
}

func Fatalw(msg string, kv ...interface{}) {
	std.Fatalw(msg, kv...)
}

//

func Debugln(args ...interface{}) {
	std.Debugln(args...)
}

func Infoln(args ...interface{}) {
	std.Infoln(args...)
}

func Warnln(args ...interface{}) {
	std.Warnln(args...)
}

func Errorln(args ...interface{}) {
	std.Errorln(args...)
}

func DPanicln(args ...interface{}) {
	std.DPanicln(args...)
}

func Panicln(args ...interface{}) {
	std.Panicln(args...)
}

func Fatalln(args ...interface{}) {
	std.Fatalln(args...)
}
