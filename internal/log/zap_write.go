package log

import (
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"time"
)
var logger *zap.Logger

type ZapWriteLogger struct {
	log  *zap.Logger
	Sync func() error
}

// logpath 日志文件路径
// loglevel 日志级别
func NewZapWriteLogger(logpath string, env string, level zap.AtomicLevel, opts ...zap.Option) *ZapWriteLogger {
	// 日志分割
	hook := lumberjack.Logger{
		Filename:   logpath, // 日志文件路径，默认 os.TempDir()
		MaxSize:    10,      // 每个日志文件保存10M，默认 100M
		MaxBackups: 30,      // 保留30个备份，默认不限
		MaxAge:     7,       // 保留7天，默认不限
		Compress:   true,    // 是否压缩，默认不压缩
	}
	var write zapcore.WriteSyncer
	write = zapcore.AddSync(&hook)
	if env == "test" {
		write = zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), write)
	}
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		// CallerKey:      "linenum",
		// MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,  // 小写编码器
		EncodeTime:     zapcore.ISO8601TimeEncoder,     // ISO8601 UTC 时间格式
		EncodeDuration: zapcore.SecondsDurationEncoder, //
		EncodeCaller:   zapcore.FullCallerEncoder,      // 全路径编码器
		EncodeName:     zapcore.FullNameEncoder,
	}
	// 设置日志级别
	//atomicLevel := zap.NewAtomicLevel()
	//atomicLevel.SetLevel(level)
	core := zapcore.NewCore(
		// zapcore.NewConsoleEncoder(encoderConfig),
		zapcore.NewJSONEncoder(encoderConfig),
		// zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(&write)), // 打印到控制台和文件
		write,
		level,
	)
	// 设置初始化字段,如：添加一个服务器名称
	// filed := zap.Fields(zap.String("serviceName", "serviceName"))
	// 构造日志
	logger = zap.New(core,  opts...)
	return &ZapWriteLogger{log: logger, Sync: logger.Sync}
}

func (l *ZapWriteLogger) Log(level log.Level, keyvals ...interface{}) error {
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

func main() {
	// 历史记录日志名字为：all.log，服务重新启动，日志会追加，不会删除
	NewZapWriteLogger("./all.log", "pre", zap.NewAtomicLevelAt(zapcore.DebugLevel), zap.AddStacktrace(
		zap.NewAtomicLevelAt(zapcore.ErrorLevel)),zap.AddCaller(), zap.Development())
	// 强结构形式
	logger.Info("test",
		zap.String("string", "string"),
		zap.Int("int", 3),
		zap.Duration("time", time.Second),
	)
	// 必须 key-value 结构形式 性能下降一点
	logger.Sugar().Infow("test-",
		"string", "string",
		"int", 1,
		"time", time.Second,
	)
}
