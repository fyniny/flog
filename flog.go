package flog

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var (
	log   *zap.Logger
	level zap.AtomicLevel
)

type Logger interface {
	Debug(msg string, fields ...Field)
	Info(msg string, fields ...Field)
	Warn(msg string, fields ...Field)
	Error(msg string, fields ...Field)
	Fatal(msg string, fields ...Field)
}

// Config 日志配置
type Config struct {
	Level      string `mapstructure:"level"`
	Filename   string `mapstructure:"filename"`
	MaxSize    int    `mapstructure:"max_size"`
	MaxBackups int    `mapstructure:"max_backups"`
	MaxAge     int    `mapstructure:"max_age"`
	Compress   bool   `mapstructure:"compress"`
}

type (
	Level = zapcore.Level
	Field = zap.Field
)

var DebugLevel = zapcore.DebugLevel
var InfoLevel = zapcore.InfoLevel
var WarnLevel = zapcore.WarnLevel
var ErrorLevel = zapcore.ErrorLevel
var FatalLevel = zapcore.FatalLevel

func Init(cnf Config) error {
	// 创建日志目录
	if err := os.MkdirAll(filepath.Dir(cnf.Filename), 0755); err != nil {
		return fmt.Errorf("failed to create log directory: %w", err)
	}

	// 设置日志级别
	level, err := zapcore.ParseLevel(cnf.Level)
	if err != nil {
		return fmt.Errorf("failed to parse log level: %w", err)
	}

	// 创建日志轮转配置
	hook := &lumberjack.Logger{
		Filename:   cnf.Filename,
		MaxSize:    cnf.MaxSize,    // 每个日志文件的最大尺寸（MB）
		MaxBackups: cnf.MaxBackups, // 保留的旧日志文件数量
		MaxAge:     cnf.MaxAge,     // 日志文件保留天数
		Compress:   cnf.Compress,   // 是否压缩旧日志
	}

	// 创建编码器配置
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder

	// 创建核心
	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderConfig),
		zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(hook)),
		level,
	)

	// 创建日志记录器
	log = zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))
	return nil
}

func init() {
	var err error
	c := zap.NewProductionConfig()
	level = zap.NewAtomicLevelAt(DebugLevel)
	c.Encoding = "console"
	c.Level = level
	log, err = c.Build(
		zap.AddCaller(),
		zap.AddCallerSkip(1),
	)
	if err != nil {
		panic(err)
	}
}

func SetLevel(lvl Level) {
	level.SetLevel(lvl)
}

func DebugWithContext(ctx context.Context, msg string, fields ...Field) {
	log.Debug(msg, fields...)
}

func InfoWithContext(ctx context.Context, msg string, fields ...Field) {
	log.Info(msg, fields...)
}

func WarnWithContext(ctx context.Context, msg string, fields ...Field) {
	log.Warn(msg, fields...)
}

func ErrorWithContext(ctx context.Context, msg string, fields ...Field) {
	log.Error(msg, fields...)
}

func FatalWithContext(ctx context.Context, msg string, fields ...Field) {
	log.Fatal(msg, fields...)
}

func Debug(msg string, fields ...Field) {
	log.Debug(msg, fields...)
}

func Info(msg string, fields ...Field) {
	log.Info(msg, fields...)
}

func Warn(msg string, fields ...Field) {
	log.Warn(msg, fields...)
}

func Error(msg string, fields ...Field) {
	log.Error(msg, fields...)
}

func Fatal(msg string, fields ...Field) {
	log.Fatal(msg, fields...)
}
