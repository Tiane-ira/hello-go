package zlog

import (
	"io"
	"os"
	"path/filepath"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var Logger *zap.Logger

const (
	defaultLevel      = zap.DebugLevel
	defaultTimeLayout = "2006-01-02 15:04:05"
)

type option struct {
	level          zapcore.Level
	fields         map[string]string
	file           io.Writer
	timeLayout     string
	disableConsole bool
}

type Option func(*option)

func WithDebugLevel() Option {
	return func(o *option) {
		o.level = zapcore.DebugLevel
	}
}

func WithInfoLevel() Option {
	return func(o *option) {
		o.level = zapcore.InfoLevel
	}
}

func WithWarnLevel() Option {
	return func(o *option) {
		o.level = zapcore.WarnLevel
	}
}

func WithErrLevel() Option {
	return func(o *option) {
		o.level = zapcore.ErrorLevel
	}
}

func WithFields(k, v string) Option {
	return func(o *option) {
		o.fields[k] = v
	}
}

// WithFile 打印日志到文件,不存在目录会自动创建
func WithFile(file string) Option {
	dir := filepath.Dir(file)
	err := os.MkdirAll(dir, 0766)
	if err != nil {
		panic(err)
	}

	f, err := os.OpenFile(file, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0766)
	if err != nil {
		panic(err)
	}
	return func(o *option) {
		o.file = zapcore.Lock(f)
	}

}

// WithFileRotate 打印日志到文件,自动滚动大小和清理日志
func WithFileRotate(file string) Option {
	dir := filepath.Dir(file)
	err := os.MkdirAll(dir, 0766)
	if err != nil {
		panic(err)
	}

	return func(o *option) {
		o.file = &lumberjack.Logger{
			Filename:   file,
			MaxSize:    128,  // 单个文件最大128M
			MaxBackups: 3,    // 文件数量最大值
			MaxAge:     28,   // 日志保留天数最大值
			LocalTime:  true, // 使用本地时间文件文件是否过期
			Compress:   true, // 是否压缩,默认不压缩
		}
	}

}

func WithTimeLayout(layout string) Option {
	return func(o *option) {
		o.timeLayout = layout
	}
}

func WithDisableConsole() Option {
	return func(o *option) {
		o.disableConsole = true
	}
}

func InitJsonZap(opts ...Option) {
	opt := &option{
		level:      defaultLevel,
		fields:     make(map[string]string),
		timeLayout: defaultTimeLayout,
	}

	for _, f := range opts {
		f(opt)
	}
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:       "time",
		LevelKey:      "level",
		NameKey:       "logger",
		CallerKey:     "caller",
		MessageKey:    "msg",
		StacktraceKey: "stacktrace",
		LineEnding:    zapcore.DefaultLineEnding,
		EncodeLevel:   zapcore.LowercaseLevelEncoder,
		EncodeTime: func(t time.Time, pae zapcore.PrimitiveArrayEncoder) {
			pae.AppendString(t.Format(opt.timeLayout))
		},
		EncodeDuration: zapcore.MillisDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}

	jsonEncoder := zapcore.NewJSONEncoder(encoderConfig)

	low := zap.LevelEnablerFunc(func(l zapcore.Level) bool {
		return l >= opt.level && l < zap.ErrorLevel
	})

	high := zap.LevelEnablerFunc(func(l zapcore.Level) bool {
		return l >= opt.level && l >= zap.ErrorLevel
	})

	stdout := zapcore.Lock(os.Stdout)
	stderr := zapcore.Lock(os.Stderr)

	core := zapcore.NewTee()

	// 输出到控制台
	if !opt.disableConsole {
		core = zapcore.NewTee(
			zapcore.NewCore(
				jsonEncoder,
				zapcore.NewMultiWriteSyncer(stdout),
				low,
			),
			zapcore.NewCore(
				jsonEncoder,
				zapcore.NewMultiWriteSyncer(stderr),
				high,
			),
		)
	}

	if opt.file != nil {
		core = zapcore.NewTee(
			core,
			zapcore.NewCore(jsonEncoder,
				zapcore.AddSync(opt.file),
				zap.LevelEnablerFunc(func(l zapcore.Level) bool {
					return l >= opt.level
				})),
		)
	}

	Logger = zap.New(core, zap.AddCaller(), zap.ErrorOutput(stderr))

	// 添加额外属性
	for k, v := range opt.fields {
		Logger = Logger.WithOptions(zap.Fields(zapcore.Field{Key: k, Type: zapcore.StringType, String: v}))
	}

}
