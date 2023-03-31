package zlog

import (
	"fmt"
	"log"
	"os"
	"strings"
	"sync"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const envLevel = "ZLOG"   // info,module_name=debug
const envDay = "ZLOG_DAY" // path_dir,module_name=path_dir

var gBaseMutex sync.RWMutex                            // global
var gBaseLoggers = make(map[string]*zap.SugaredLogger) // [subsystem]=logger
var gBaseLevels = make(map[string]zap.AtomicLevel)     // [subsystem]=level
var gBaseDays = make(map[string]*dayWriter)            // [subsystem]=LogByDay
var gSetUp = false

var encoderConfig = zapcore.EncoderConfig{
	TimeKey:       "time",
	LevelKey:      "level",
	NameKey:       "logger",
	CallerKey:     "linenum",
	MessageKey:    "msg",
	StacktraceKey: "trace",
	LineEnding:    zapcore.DefaultLineEnding,
	EncodeLevel:   zapcore.CapitalLevelEncoder, // level 格式
	// EncodeTime:     zapcore.RFC3339NanoTimeEncoder, // 日期 格式
	EncodeTime: func(t time.Time, pae zapcore.PrimitiveArrayEncoder) {
		pae.AppendString(t.Format("2006-01-02T15:04:05.000")) //
	},
	EncodeDuration: zapcore.SecondsDurationEncoder, //
	EncodeCaller:   zapcore.ShortCallerEncoder,     // 路径格式
	EncodeName:     zapcore.FullNameEncoder,
}

func getBaseLogger(name string, skipN int) (*zap.SugaredLogger, zap.AtomicLevel) {

	gBaseMutex.Lock()
	defer gBaseMutex.Unlock()

	baseSugarLogger, ok := gBaseLoggers[name]
	if ok {
		return baseSugarLogger, gBaseLevels[name]
	}

	// creat new

	// lv
	var lv zap.AtomicLevel
	if base_lv, ok := gBaseLevels[name]; !ok {
		if defaut_lv, ok3 := gBaseLevels[""]; ok3 {
			lv = defaut_lv
		} else {
			lv = zap.NewAtomicLevelAt(zap.DebugLevel)
		}
	} else {
		lv = base_lv
	}
	// day
	var ws zapcore.WriteSyncer
	if ld, ok2 := gBaseDays[name]; !ok2 {
		if defaut_ld, ok3 := gBaseDays[""]; ok3 {
			ws = zapcore.NewMultiWriteSyncer(zapcore.AddSync(defaut_ld))
		} else {
			ws = zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout))
		}
	} else {
		ws = zapcore.NewMultiWriteSyncer(zapcore.AddSync(ld))
	}
	// core
	core := zapcore.NewCore(
		zapcore.NewConsoleEncoder(encoderConfig),
		ws,
		lv,
	)

	baseSugarLogger = zap.New(core, zap.AddCaller()).
		Sugar().
		WithOptions(zap.AddCallerSkip(skipN))
	gBaseLoggers[name] = baseSugarLogger

	return baseSugarLogger, gBaseLevels[name]
}

// str=info,module_name=debug
func SetLevel(str string) error {
	for _, kv := range strings.Split(str, ",") {
		parts := strings.Split(kv, "=")

		moduleName := ""
		lvStr := ""
		switch len(parts) {
		case 1:
			moduleName = ""
			lvStr = parts[0]
		case 2:
			moduleName = parts[0]
			lvStr = parts[1]
		}

		lv, err := zapcore.ParseLevel(lvStr)
		if err != nil {
			return fmt.Errorf("ParseLevel: %v, %v", kv, err)
		}

		gBaseMutex.Lock()
		baseLv, ok := gBaseLevels[moduleName]
		if !ok {
			baseLv = zap.NewAtomicLevelAt(lv)
			gBaseLevels[moduleName] = baseLv
		} else {
			baseLv.SetLevel(lv)
		}
		gBaseMutex.Unlock()
	}
	return nil
}

// str=path_dir,module_name=path_dir
func iniDay(str string) error {
	for _, kv := range strings.Split(str, ",") {
		parts := strings.Split(kv, "=")

		moduleName := ""
		logDir := ""
		switch len(parts) {
		case 1:
			moduleName = ""
			logDir = parts[0]
		case 2:
			moduleName = parts[0]
			logDir = parts[1]
		}

		ld, err := newDayWriter(logDir, "log")
		if err != nil {
			return fmt.Errorf("newDayWriter: %v, %v", kv, err)
		}
		gBaseDays[moduleName] = ld
	}
	return nil
}

func setUp() {
	// TODO: init form config

	gBaseMutex.Lock()
	if gSetUp {
		gBaseMutex.Unlock()
		return
	}
	gSetUp = true
	gBaseMutex.Unlock()

	// first
	dayStr := os.Getenv(envDay)
	if dayStr != "" {
		if err := iniDay(dayStr); err != nil {
			log.Printf("zlog iniDay: %v", err)
			return
		}
	}
	// second
	lvStr := os.Getenv(envLevel)
	if lvStr == "" {
		lvStr = "debug"
	}
	if err := SetLevel(lvStr); err != nil {
		log.Printf("zlog SetLevel: %v", err)
	}
}
