/*
 * @PackageName: logs
 * @Description:
 * @Author: limuzhi
 * @Date: 2022/12/3 11:39
 */

package logs

import (
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

var _ log.Logger = (*ZapLogger)(nil)

// ZapLogger is a logger impl.
type ZapLogger struct {
	log           *zap.Logger
	Sync          func() error
	infoFilePath  string
	errorFilePath string
	fileSuffix    string
	logMode       string //dev test pro
	logLevel      zapcore.Level
}

func NewLogger(opts ...OptionFunc) log.Logger {
	//配置zap日志库的编码器
	o := &ZapLogger{
		log:           nil,
		Sync:          nil,
		infoFilePath:  "info",
		errorFilePath: "error",
		fileSuffix:    ".log",
		logLevel:      zap.InfoLevel,
		logMode:       "dev",
	}
	for _, opt := range opts {
		opt(o)
	}
	encoder := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		MessageKey:     "msg",
		StacktraceKey:  "stack",
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.CapitalLevelEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.FullCallerEncoder,
	}
	logger := o.NewZapLogger(
		encoder,
		zap.NewAtomicLevelAt(o.logLevel),
		zap.AddStacktrace(
			zap.NewAtomicLevelAt(zapcore.ErrorLevel)),
		zap.AddCaller(),
		zap.AddCallerSkip(4),
		zap.Development(),
	)
	return log.With(logger)
}

type OptionFunc func(*ZapLogger)

func WithInfoFilePath(infoFilePath string) OptionFunc {
	return func(o *ZapLogger) {
		o.infoFilePath = infoFilePath
	}
}

func WithErrorFilePath(errorFilePath string) OptionFunc {
	return func(o *ZapLogger) {
		o.errorFilePath = errorFilePath
	}
}

func WithFileSuffix(fileSuffix string) OptionFunc {
	return func(o *ZapLogger) {
		o.fileSuffix = fileSuffix
	}
}

func WithLogLevel(logLevel zapcore.Level) OptionFunc {
	return func(o *ZapLogger) {
		o.logLevel = logLevel
	}
}

func WithLogMode(logMode string) OptionFunc {
	return func(o *ZapLogger) {
		o.logMode = logMode
	}
}

// 日志自动切割，采用 lumberjack 实现的
func (l *ZapLogger) getLogWriter(filename string) zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   filename, //文件保存路径
		MaxSize:    500,      //日志的最大大小（M）
		MaxBackups: 30,       //日志的最大保存数量
		MaxAge:     7,        //日志文件存储最大天数
		Compress:   true,     //是否执行压缩
		//LocalTime:  true,
	}
	return zapcore.AddSync(lumberJackLogger)
}

// NewZapLogger return a zap logger.
func (l *ZapLogger) NewZapLogger(encoder zapcore.EncoderConfig, level zap.AtomicLevel, opts ...zap.Option) *ZapLogger {
	var coreArr = []zapcore.Core{}
	//设置日志级别
	//日志都会在console中展示，本地环境打开
	logLevel := level.Level()

	//控制台日志
	if l.logMode == "dev" {
		devZapCore := zapcore.NewCore(
			zapcore.NewConsoleEncoder(encoder),                             //编码器配置
			zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout)), level) // //打印到控制台 线上可注释掉
		coreArr = append(coreArr, devZapCore)
	}
	//日志切割
	infoFileName := fmt.Sprintf("%s%s", l.infoFilePath, l.fileSuffix)
	infoSyncer := l.getLogWriter(infoFileName)
	// 实现两个判断日志等级的interface (其实 zapcore.*Level 自身就是 interface)
	infoLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl < zapcore.WarnLevel && lvl >= logLevel
	})
	infoZapCore := zapcore.NewCore(zapcore.NewJSONEncoder(encoder), zapcore.AddSync(infoSyncer), infoLevel)
	coreArr = append(coreArr, infoZapCore)

	errorFileName := fmt.Sprintf("%s%s", l.errorFilePath, l.fileSuffix)
	errorSyncer := l.getLogWriter(errorFileName)
	errorLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= zapcore.WarnLevel
	})
	errorZapCore := zapcore.NewCore(zapcore.NewJSONEncoder(encoder), zapcore.AddSync(errorSyncer), errorLevel)
	coreArr = append(coreArr, errorZapCore)

	core := zapcore.NewTee(coreArr...)
	zapLogger := zap.New(core, opts...)
	l.log = zapLogger
	l.Sync = zapLogger.Sync
	return l
}

// Log Implementation of logger interface.
func (l *ZapLogger) Log(level log.Level, keyvals ...interface{}) error {
	if len(keyvals) == 0 || len(keyvals)%2 != 0 {
		l.log.Warn(fmt.Sprint("Keyvalues must appear in pairs: ", keyvals))
		return nil
	}
	// Zap.Field is used when keyvals pairs appear
	var data []zap.Field
	for i := 0; i < len(keyvals); i += 2 {
		data = append(data, zap.Any(fmt.Sprint(keyvals[i]), fmt.Sprint(keyvals[i+1])))
	}
	switch level {
	case log.LevelDebug:
		l.log.Debug("", data...)
	case log.LevelInfo:
		l.log.Info("", data...)
	case log.LevelWarn:
		l.log.Warn("", data...)
	case log.LevelError:
		l.log.Error("", data...)
	}
	return nil
}
